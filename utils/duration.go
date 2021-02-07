package utils

import (
	"time"
	"github.com/ccding/go-logging/logging"
)

var logger *logging.Logger

func init() {
	logger, _ = logging.FileLogger("RuncTiming", logging.INFO, logging.BasicFormat, logging.DefaultTimeFormat, "/tmp/myrunc/log.txt", false)
}


func Track(msg string) (string, time.Time) {
	logger.Infof("%v begin:", msg)
	// logrus.Info(fmt.Sprintf("%v begin", msg))
    return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	logger.Infof("%v end, %v", msg, time.Since(start))
    // logrus.Info(fmt.Sprintf("%v end: %v", msg, time.Since(start)))
}

func Tik(msg string) time.Time {
	logger.Infof("%v begin:", msg)
	// logrus.Info(fmt.Sprintf("%v begin", msg))
	return time.Now()
}

func LogFlush() {
	logger.Destroy()
}