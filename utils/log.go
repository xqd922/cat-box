package utils

import (
	"log"
	"os"

	"github.com/daifiyum/cat-box/config"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
	logFile  *os.File
)

func init() {
	var err error
	logFile, err = os.OpenFile(config.Config("LOG_PATH"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	// 创建信息日志记录器
	infoLog = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	// 创建错误日志记录器
	errorLog = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string) {
	infoLog.Println(message)
}

func LogError(message string) {
	errorLog.Println(message)
}

func GetLogFile() *os.File {
	return logFile
}

func CloseLog() {
	logFile.Close()
}
