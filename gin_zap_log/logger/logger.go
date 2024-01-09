package logger

import (
	"gin_zap_log/config"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var lg *zap.Logger

// core 三个参数之 日志编码器
func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	encoderConfig.MessageKey = "msg"
	encoderConfig.CallerKey = "caller"
	return zapcore.NewJSONEncoder(encoderConfig)
	//return zapcore.NewConsoleEncoder(encoderConfig)
}

// core 三个参数之 日志输出器
func GetLogWriter() zapcore.WriteSyncer {
	lumberjackLoger := &lumberjack.Logger{
		Filename:   "D:\\code\\golang\\code\\gin_zap_log\\gin.log",
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberjackLoger), zapcore.AddSync(os.Stdout))
	//return zapcore.AddSync(zapcore.AddSync(lumberjackLoger))
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.Request.URL.Path
		queryRaw := context.Request.URL.RawQuery
		context.Next()

		cost := time.Since(start)
		lg.Info(
			path,
			zap.Int("StatusCode", context.Writer.Status()),
			zap.String("method", context.Request.Method),
			zap.String("path", path),
			zap.String("query", queryRaw),
			zap.String("ip", context.ClientIP()),
			zap.String("user-agent", context.Request.UserAgent()),
			zap.String("errors", context.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关的日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				var brokenPipe bool
				ne, ok := err.(*net.OpError)
				if ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				dumpRequest, _ := httputil.DumpRequest(context.Request, false)
				if brokenPipe {
					lg.Error(
						context.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(dumpRequest)),
					)

					context.Error(err.(error))
					context.Abort()
					return
				}

				if stack {
					lg.Error(
						"[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(dumpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error(
						"[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(dumpRequest)),
					)
				}

				context.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		context.Next()
	}
}

// InitLogger 初始化Logger
func InitLogger(cfg config.LogConfig) (err error) {
	encoder := GetEncoder()
	writeSyncer := GetLogWriter()
	l := new(zapcore.Level)
	err = l.UnmarshalText([]byte((cfg.Level)))
	if err != nil {
		return err
	}

	core := zapcore.NewCore(encoder, writeSyncer, l)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger) //替换zap包中全局的Logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}
