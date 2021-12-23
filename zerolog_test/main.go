package main

import (
	"example/zerolog_test/logger"
	//"io"
	//"os"
)

/**
1. 输出级别
2. 输出重定向
3. 子 log
如果是 debug, 就输出容易肉眼看的
否则输出格式化日志 json 格式
*/
var log = logger.NewLog("./log/info.log")

func main() {
	log.Debug().Msg("This message appears only when log level set to debug")
	log.Info().Msg("This message appears when log level set to debug or info")

	if e := log.Debug(); e.Enabled() {
		e.Str("foo", "bar").Msg("some debug message")
	}

	//io.MultiWriter(os.Stderr, io.Writer())
	//mainLogger := zerolog.New(os.Stderr).With().Logger()
	//mainLogger.Info().Msg("This is the output from the main logger")
	//subLogger := mainLogger.
	//subLogger := mainLogger.With().Str("component","componentA").Logger()
	//subLogger.Info().Msg("This is the the extended output from the sublogger")
}
