package logrus

import (
	"fmt"
	"os"
	// "io"

	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (log *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{log.WithField(k, v)}
}

func init() {

	log := logrus.New()

	file, _ := os.OpenFile("logrus.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		EnvironmentOverrideColors: true,
		DisableColors:             false,
		FullTimestamp:             true,
	}
	// log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(log)
}
