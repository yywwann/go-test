package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

//只能输出结构化日志，但是性能要高于 SugaredLogger
var Logger *zap.Logger
//可以输出 结构化日志、非结构化日志。性能茶语 zap.Logger，具体可见上面的的单元测试
var Log *zap.SugaredLogger

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,  //最大M数，超过则切割
		MaxBackups: 5,   //最大文件保留数，超过就删除最老的日志文件
		MaxAge:     30,  //保存30天
		Compress:   false,//是否压缩
	}
}

func InitLog(logPath, errPath string, logLevel zapcore.Level) {
	config := zapcore.EncoderConfig{
		MessageKey:   "msg",  //结构化（json）输出：msg的key
		LevelKey:     "level",//结构化（json）输出：日志级别的key（INFO，WARN，ERROR等）
		TimeKey:      "ts",   //结构化（json）输出：时间的key（INFO，WARN，ERROR等）
		CallerKey:    "file", //结构化（json）输出：打印日志的文件对应的Key
		EncodeLevel:  zapcore.CapitalLevelEncoder, //将日志级别转换成大写（INFO，WARN，ERROR等）
		EncodeCaller: zapcore.ShortCallerEncoder, //采用短文件路径编码输出（test/main.go:14 ）
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			//enc.AppendString(t.Format("2006-01-02 15:04:05"))
			enc.AppendString(t.Format(time.RFC3339))
		},//输出的时间格式
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},//
	}
	//自定义日志级别：自定义Info级别
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})

	//自定义日志级别：自定义Warn级别
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})

	// 获取io.Writer的实现
	infoWriter := getWriter(logPath)
	warnWriter := getWriter(errPath)

	// 实现多个输出
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(infoWriter), infoLevel), //将info及以下写入logPath，NewConsoleEncoder 是非结构化输出
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(warnWriter), warnLevel),//warn及以上写入errPath
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel),//同时将日志输出到控制台，NewJSONEncoder 是结构化输出
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	Log = Logger.Sugar()
}

func init() {
	InitLog("./log/info.log", "./log/error.log", zap.InfoLevel)
}