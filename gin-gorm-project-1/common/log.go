package common

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var Logger *zap.Logger
var SugaredLogger *zap.SugaredLogger

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:    "M",
		LevelKey:      "L",
		TimeKey:       "T",
		NameKey:       "N",
		CallerKey:     "C",
		StacktraceKey: "Stack",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime:    timeEncoder,
		EncodeLevel:   levelEncoder,
		//EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   shortCallerEncoder,
	}

}

func shortCallerEncoder(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(fmt.Sprintf("[%s]", caller.TrimmedPath()))
}

func levelEncoder(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
	var levelStr string
	switch level {
	case zapcore.DebugLevel:
		levelStr = "[DEBUG]"
	case zapcore.InfoLevel:
		levelStr = "[INFO]"
	case zapcore.WarnLevel:
		levelStr = "[WARN]"
	case zapcore.ErrorLevel:
		levelStr = "[ERROR]"
	case zapcore.DPanicLevel:
		levelStr = "[DPANIC]"
	case zapcore.PanicLevel:
		levelStr = "[PANIC]"
	case zapcore.FatalLevel:
		levelStr = "[FATAL]"
	default:
		levelStr = fmt.Sprintf("[LEVEL(%d)]", level)
	}
	encoder.AppendString(levelStr)
}

func timeEncoder(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
}

// core 三个参数之 Encoder获取编码器
func GetEncoder() zapcore.Encoder {
	//自定义日志编码配置，下方NewJSONEncoder输出如下的日志格式
	//{"L":"[INFO]","T":"2022-09-16 14:24:59.552","C":"[prototest/main.go:113]","M":"name = xiaoli, age = 18"}
	//return zapcore.NewJSONEncoder(NewEncoderConfig())

	//下方NewConsoleEncoder输出如下的日志格式
	//2022-09-16 14:26:02.933 [INFO]  [prototest/main.go:113] name = xiaoli, age = 18
	return zapcore.NewConsoleEncoder(NewEncoderConfig())
}

// core 三个参数之 日志记录器，日志输出路径
func GetInforWriterSyncer() zapcore.WriteSyncer {
	//info级别日志写入文件路径
	//file, _ := os.Create("D:\\code\\golang\\code\\zap_log\\info.log")
	//return zapcore.AddSync(file)

	//引入第三方库 Lumberjack 加入日志切割功能
	infoLumberIO := &lumberjack.Logger{
		Filename:   "D:\\code\\i-go\\gin-gorm-project-1\\info.log",
		MaxSize:    1,
		MaxAge:     3,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	}

	return zapcore.AddSync(infoLumberIO)
}

func GetErrorWriter() zapcore.WriteSyncer {
	errorLumberIO := &lumberjack.Logger{
		Filename:   "D:\\code\\i-go\\gin-gorm-project-1\\error.log",
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 1,
		LocalTime:  true,
		Compress:   false,
	}
	return zapcore.AddSync(errorLumberIO)
}

// logs.Debug(...)再封装
func Debugf(format string, v ...interface{}) {
	Logger.Sugar().Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	Logger.Sugar().Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	Logger.Sugar().Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	Logger.Sugar().Errorf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	Logger.Sugar().Fatalf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	Logger.Sugar().Panicf(format, v...)
}

// logs.Debug(...)再封装
func Debug(format string, fields ...zapcore.Field) {
	Logger.Debug(format, fields...)
}

func Info(format string, fields ...zapcore.Field) {
	Logger.Info(format, fields...)
}

func Warn(format string, fields ...zapcore.Field) {
	Logger.Warn(format, fields...)
}

func Error(format string, field ...zapcore.Field) {
	Logger.Error(format, field...)
}

func Panic(format string, field ...zapcore.Field) {
	Logger.Error(format, field...)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.Request.URL.Path
		queryRaw := context.Request.URL.RawQuery
		// 读取请求体
		bodyBytes, _ := io.ReadAll(context.Request.Body)
		// 恢复请求体，因为读取后 Body 就被清空了
		context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		context.Next()

		cost := time.Since(start)
		Logger.Info(
			path,
			zap.Int("StatusCode", context.Writer.Status()),        //  状态码
			zap.String("method", context.Request.Method),          //  请求方法
			zap.String("path", path),                              //  请求路径
			zap.String("query-param", queryRaw),                   //  请求中的查询参数
			zap.String("request-body", string(bodyBytes)),         //  请求中的请求体
			zap.String("ip", context.ClientIP()),                  //  客户端IP地址
			zap.String("user-agent", context.Request.UserAgent()), //  客户端UA
			zap.String("errors", context.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost-time", cost), //  执行时间
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
					Logger.Error(
						context.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(dumpRequest)),
					)

					context.Error(err.(error))
					context.Abort()
					return
				}

				if stack {
					Logger.Error(
						"[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(dumpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					Logger.Error(
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

func InitLogger() {
	//获取编码器
	encoder := GetEncoder()

	//日志级别
	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool { //  error级别
		return level >= zap.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel && level < zapcore.ErrorLevel
	})

	//info文件写入器
	inforWriterSyncer := GetInforWriterSyncer()
	//error文件写入器
	errorWriterSyncer := GetErrorWriter()

	//生成core日志记录器
	//同时输出日志到控制台和指定的日志文件中
	infoFileCore := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(inforWriterSyncer, zapcore.AddSync(os.Stdout)),
		lowPriority,
	)

	errorFileCore := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(errorWriterSyncer, zapcore.AddSync(os.Stdout)),
		highPriority,
	)

	//将infoFileCore和errorFileCore加入core切片
	var coreArr []zapcore.Core
	coreArr = append(coreArr, infoFileCore, errorFileCore)

	//生成日志记录器实例
	Logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
	SugaredLogger = Logger.Sugar()
	zap.ReplaceGlobals(Logger) //替换zap包中全局的Logger实例，后续在其他包中只需使用zap.L()调用即可
}
