package log

import (
	"github.com/ishansd94/sample-app/pkg/env"
	"github.com/sirupsen/logrus"
)

func init() {
	if env.Get("LOG_LEVEL" , "debug") == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
		// logrus.SetReportCaller(true)
		Debug("gateway", "logging level is set to debug")
	}
}

func Debug(process string, msg string) {
	logrus.Debugf("[%s] %s", process, msg)
}

func Info(process string, msg string) {
	logrus.Infof("[%s] %s \n", process, msg)
}

func Warn(process string, msg string) {
	logrus.Warnf("[%s] %s \n", process, msg)
}

func Error(process string, msg string, err error) {
	logrus.Errorf("[%s] %s \n %s", process, msg, err.Error())
}

func Fatal(process string, msg string, err error) {
	logrus.Fatalf("[%s] %s \n %s", process, msg, err.Error())
}