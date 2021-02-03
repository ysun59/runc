package utils

import (
	"fmt"
	"time"
	"github.com/sirupsen/logrus"
)

func Track(msg string) (string, time.Time) {
	logrus.Info(fmt.Sprintf("%v begin", msg))
    return msg, time.Now()
}

func Duration(msg string, start time.Time) {
    logrus.Info(fmt.Sprintf("%v end: %v", msg, time.Since(start)))
}

func Tik(msg string) time.Time {
	tik := time.Now()
	logrus.Info(fmt.Sprintf("%v begin", msg))
	return tik
}