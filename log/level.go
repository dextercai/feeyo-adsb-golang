package log

import (
	"github.com/sirupsen/logrus"
)

const (
	LogLevelError = "error"
	LogLevelWarn  = "warn"
	LogLevelInfo  = "info"
	LogLevelDebug = "debug"
	LogLevelTrace = "trace"
)

var LevelMap = map[string]logrus.Level{
	LogLevelError: logrus.ErrorLevel,
	LogLevelWarn:  logrus.WarnLevel,
	LogLevelInfo:  logrus.InfoLevel,
	LogLevelDebug: logrus.DebugLevel,
	LogLevelTrace: logrus.TraceLevel,
}
