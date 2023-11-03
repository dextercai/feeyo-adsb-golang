package log

import (
	"fmt"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000 MST",
	})
}

func InitLog(logLevel, logPath, logFile string, logRotationTime, logMaxAge, logRotationSize int64, logRotationCount uint) {
	if _, ok := LevelMap[logLevel]; !ok {
		Logger.WithField("scope", "log").Warnf("输入了不存在的日志等级: %s，将使用: %s", logLevel, "info")
		logLevel = "info"
	}
	Logger.SetLevel(LevelMap[logLevel])

	absPath, _ := os.Getwd()
	if filepath.IsAbs(logPath) {
		absPath = ""
	}

	lpth := fmt.Sprintf("%s/%s", absPath, logPath)
	if !FileExists(lpth) {
		err := os.Mkdir(lpth, os.ModePerm)
		if err != nil {
			Logger.WithField("scope", "log").WithError(err).Fatalf("创建日志文件夹失败")
		}
	}

	logFileName := fmt.Sprintf("%s/%s/%s.%s", absPath, logPath, logFile, "%Y-%m-%d")
	logFileName, _ = filepath.Abs(logFileName)
	logier, err := rotateLogs.New(
		logFileName,
		//rotateLogs.WithLinkName(fmt.Sprintf("%s/%s", logPath, logFile)), # TODO Win兼容性存疑
		rotateLogs.WithRotationTime(time.Duration(logRotationTime)*time.Second),
		rotateLogs.WithMaxAge(time.Duration(logMaxAge)*time.Second),
		rotateLogs.WithRotationSize(logRotationSize<<20),
		rotateLogs.WithRotationCount(logRotationCount),
	)

	if err != nil {
		Logger.WithField("scope", "log").WithError(err).Fatalf("日志轮转配置错误")
	}
	var lfHook *lfshook.LfsHook

	lfHook = lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: logier,
		logrus.InfoLevel:  logier,
		logrus.WarnLevel:  logier,
		logrus.ErrorLevel: logier,
		logrus.FatalLevel: logier,
		logrus.PanicLevel: logier,
	}, &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05.000 MST"})

	Logger.AddHook(lfHook)

}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	Logger.WithField("scope", "log").WithError(err).Errorf("FileExists err")
	return false
}
