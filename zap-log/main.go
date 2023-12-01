package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"sync"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//return zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewJSONEncoder(encoderConfig)
}

func GetLogWriter() zapcore.WriteSyncer {
	lumLogConfig := &lumberjack.Logger{
		Filename:   "D:\\code\\golang\\code\\zap-log\\test.log", //日志文件存储路径及文件名，不存在则自动创建
		MaxSize:    1,
		MaxAge:     3,
		MaxBackups: 3,
		LocalTime:  true,  //  本地时间
		Compress:   false, // 是否对旧日志文件压缩
	}
	return zapcore.AddSync(lumLogConfig)
}

func InitLogConfig() {
	encoder := GetEncoder()
	writeSyncer := GetLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zap.InfoLevel)
	logger = zap.New(core, zap.AddCaller())
	//sugarLogger = logger.Sugar()
}

func SimpleHttpGet(url string) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		logger.Info(
			"Success fetching url...",
			zap.String("url", url),
			zap.Int("StatusCode", response.StatusCode),
			zap.String("Header", response.Header.Get("Content-Type")),
		)
		response.Body.Close()
	}

}

var wg sync.WaitGroup

func main() {
	InitLogConfig()
	defer logger.Sync()

	//模拟并发，进行1万次并发查看日志分割和输出情况
	wg.Add(10000)
	//创建一组并发，请求baidu,通过simplehttpGet方法
	for i := 0; i < 10000; i++ {
		go SimpleHttpGet("http://www.baidu.com")
	}
	wg.Wait()

}
