package logger

import (
	"flag"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"io"
	"os"
	"time"
)

var debug *bool

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,    //最大M数，超过则切割
		MaxBackups: 5,     //最大文件保留数，超过就删除最老的日志文件
		MaxAge:     30,    //保存30天
		Compress:   false, //是否压缩
	}
}

func NewLog(filename string) zerolog.Logger {
	var logger zerolog.Logger
	var out io.Writer

	if *debug {
		out = io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}, getWriter(filename))
		logger = zerolog.New(out).With().Timestamp().Caller().Logger()
	} else {
		out = io.MultiWriter(os.Stderr, getWriter(filename))
		logger = zerolog.New(out).With().Timestamp().Logger()
	}

	return logger
}

func init() {
	debug = flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	zerolog.TimeFieldFormat = time.RFC3339
}
