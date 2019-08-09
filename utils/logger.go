package utils

import (
	"time"

	"path"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

// SetLogger set logger
func SetLogger(logFileName string) *log.Logger {
	logLevelString := APIConfig.MainConfig.LogLevel
	logLevel, err := log.ParseLevel(logLevelString)
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
	logClient := log.New()
	logClient.AddHook(setLFShook(APIConfig.MainConfig.LogDir,
		logFileName, 7*24*time.Hour, 24*time.Hour))
	return logClient
}

func setLFShook(logPath, logFileName string, maxAge, rotationTime time.Duration) *lfshook.LfsHook {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	return lfshook.NewHook(
		lfshook.WriterMap{
			log.DebugLevel: writer,
			log.InfoLevel:  writer,
			log.WarnLevel:  writer,
			log.ErrorLevel: writer,
			log.FatalLevel: writer,
			log.PanicLevel: writer,
		}, &log.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05.000"})
}
