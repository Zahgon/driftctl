package logger

import (
	"bytes"
	"time"

	"github.com/sirupsen/logrus"
)

var baseTimestamp time.Time

func init() {
	baseTimestamp = time.Now()
}

// TextFormatter formats logs into text
type TextFormatter struct {
	// The max length of the level text, generated dynamically on init if == 0
	levelTextMaxLength int
}

func NewTextFormatter(levelTextMaxLength int) *TextFormatter { _ = "STUB: not implemented"; return nil }

func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *TextFormatter) writeCaller(entry *logrus.Entry, b *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *TextFormatter) writeContext(entry *logrus.Entry, b *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *TextFormatter) writeMessage(entry *logrus.Entry, b *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *TextFormatter) writeElapsedTime(entry *logrus.Entry, b *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *TextFormatter) writeLevel(entry *logrus.Entry, b *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

// TRUNCATE if needed

// and then pad to f.levelTextMaxLength
