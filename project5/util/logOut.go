package util

import (
	"fmt"
	"log"
	"os"
)

func PrintLog(user, content, logName string) {
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create ./content.log err : %v\n", err)
	}
	if logFile != nil {
		defer func(file *os.File) { file.Close() }(logFile)
	}

	debugLog := log.New(logFile, "[Broadcast]", log.Ldate)

	debugLog.SetPrefix("[Broadcast]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Printf("user:%s,content:%s", user, content)
}
