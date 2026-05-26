package logger

import (
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
	"github.com/sirupsen/logrus"
)

type terraformPluginFormatter struct {
	logrus.Formatter
}

func (f *terraformPluginFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TerraformPluginLogger struct {
	logger *logrus.Logger
}

func NewTerraformPluginLogger() TerraformPluginLogger {
	_ = "STUB: not implemented"
	return *new(TerraformPluginLogger)
}

// Disable terraform provider log if we are not in trace level

func (t TerraformPluginLogger) Trace(msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (t TerraformPluginLogger) Debug(msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (t TerraformPluginLogger) Info(msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (t TerraformPluginLogger) Warn(msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (t TerraformPluginLogger) Error(msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (t TerraformPluginLogger) IsTrace() bool { _ = "STUB: not implemented"; return false }

func (t TerraformPluginLogger) IsDebug() bool { _ = "STUB: not implemented"; return false }

func (t TerraformPluginLogger) IsInfo() bool { _ = "STUB: not implemented"; return false }

func (t TerraformPluginLogger) IsWarn() bool { _ = "STUB: not implemented"; return false }

func (t TerraformPluginLogger) IsError() bool { _ = "STUB: not implemented"; return false }

func (t TerraformPluginLogger) With(args ...interface{}) hclog.Logger {
	_ = "STUB: not implemented"
	return *new(hclog.Logger)
}

func (t TerraformPluginLogger) Named(name string) hclog.Logger {
	_ = "STUB: not implemented"
	return *new(hclog.Logger)
}

func (t TerraformPluginLogger) ResetNamed(name string) hclog.Logger {
	_ = "STUB: not implemented"
	return *new(hclog.Logger)
}

func (t TerraformPluginLogger) SetLevel(level hclog.Level) { _ = "STUB: not implemented"; return }

func (t TerraformPluginLogger) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	_ = "STUB: not implemented"
	return nil
}

func (t TerraformPluginLogger) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (t TerraformPluginLogger) Log(level hclog.Level, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (t TerraformPluginLogger) ImpliedArgs() []interface{} { _ = "STUB: not implemented"; return nil }

func (t TerraformPluginLogger) Name() string { _ = "STUB: not implemented"; return "" }
