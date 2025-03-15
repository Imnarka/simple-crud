package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func InitLogger() {
	logFile, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Ошибка при открытии файла лога: ", err)
	}

	log.SetOutput(logFile)

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}

func Info(msg string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	log.WithFields(fields).Info(msg)
}

func Error(msg string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	log.WithFields(fields).Error(msg)
}
