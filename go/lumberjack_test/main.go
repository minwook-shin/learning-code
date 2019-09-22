package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	logStruct := &lumberjack.Logger{
		Filename:   "test_file.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     10,
		Compress:   true,
	}
	log.SetOutput(logStruct)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for {
			<-c
			logStruct.Rotate()
		}
	}()

	logStruct.Write([]byte("TEST logs"))
}
