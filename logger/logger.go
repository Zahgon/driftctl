package logger

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Level        logrus.Level
	Formatter    logrus.Formatter
	ReportCaller bool
}

func Init() { _ = "STUB: not implemented"; return }

// Libs that use logger (like grpc provider) will log at TRACE level
