package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
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
		Filename:   "D:\\\\code\\\\golang\\\\code\\\\zap_log\\\\info.log",
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
		Filename:   "D:\\\\code\\\\golang\\\\code\\\\zap_log\\\\error.log",
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

}
