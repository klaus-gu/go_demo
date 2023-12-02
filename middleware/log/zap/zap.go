package main

import (
	"go_demo/middleware/log/zap/logger"
	"os"
)

func main() {
	// 控制台输出日志
	log := logger.NewConsoleLogger()
	log.Info("这是一条普通日志，将会输出到控制台", logger.String("key", "value"), logger.Int("ikey", 3))

	// 所有日志输出到指定日志文件
	logger.ResetDefault(log)
	log2 := logger.NewFileLoggerWithOption("./logs/demo.log", logger.InfoLevel)
	defer log2.Sync()
	log2.Info("这条日志将被写入到文件中", logger.String("key", "value"), logger.Int("ikey", 3))

	// 日志分级别输出不同文件
	infoFile, err := os.OpenFile("./logs/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	errFile, err := os.OpenFile("./logs/error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	options := []logger.TeeOption{
		{
			W: infoFile,
			Lef: func(lvl logger.Level) bool {
				return lvl <= logger.InfoLevel
			},
		},
		{
			W: errFile,
			Lef: func(lvl logger.Level) bool {
				return lvl > logger.ErrorLevel
			},
		},
	}
	multiWriterLogger := logger.NewMultiWriterLogger(options)
	defer multiWriterLogger.Sync()
	logger.ResetDefault(multiWriterLogger)
	multiWriterLogger.Error("错误日志", logger.String("日志类型", "ERROR"))
	multiWriterLogger.Info("普通日志", logger.String("日志类型", "INFO"))

	var tops = []logger.TeeOption{
		{
			FileName: "./logs/rotate_access.log",
			Ropt: logger.RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3,
				Compress:   true,
			},
			Lef: func(lvl logger.Level) bool {
				return lvl <= logger.InfoLevel
			},
		},
		{
			FileName: "./logs/rotate_error.log",
			Ropt: logger.RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3,
				Compress:   true,
			},
			Lef: func(lvl logger.Level) bool {
				return lvl > logger.InfoLevel
			},
		},
	}
	rotateLogger := logger.NewMultiWriterRotateLogger(tops)
	logger.ResetDefault(rotateLogger)

	// 为了演示自动rotate效果，这里多次调用log输出
	for i := 0; i < 200000; i++ {
		rotateLogger.Info("demo3:", logger.String("app", "start ok"),
			logger.Int("major version", 3))
		rotateLogger.Error("demo3:", logger.String("app", "crash"),
			logger.Int("reason", -1))
	}
}
