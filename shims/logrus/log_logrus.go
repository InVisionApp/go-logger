package logrus

import (
	"github.com/InVisionApp/go-logger"
	"github.com/sirupsen/logrus"
)

type shim struct {
	l *logrus.Entry
}

// NewLogrus can be used to override the default logger.
// Optionally pass in an existing logrus logger or pass in
// `nil` to use the default logger.
func NewLogrus(logger *logrus.Logger) log.Logger {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	return &shim{l: logrus.NewEntry(logger)}
}

// Debug log message
func (s *shim) Debug(msg ...interface{}) {
	s.l.Debug(msg...)
}

// Info log message
func (s *shim) Info(msg ...interface{}) {
	s.l.Info(msg...)
}

// Warn log message
func (s *shim) Warn(msg ...interface{}) {
	s.l.Warn(msg...)
}

// Error log message
func (s *shim) Error(msg ...interface{}) {
	s.l.Error(msg...)
}

// Debugf log message with formatting
func (s *shim) Debugf(format string, args ...interface{}) {
	s.l.Debugf(format, args...)
}

// Infof log message with formatting
func (s *shim) Infof(format string, args ...interface{}) {
	s.l.Infof(format, args...)
}

// Warnf log message with formatting
func (s *shim) Warnf(format string, args ...interface{}) {
	s.l.Warnf(format, args...)
}

// Errorf log message with formatting
func (s *shim) Errorf(format string, args ...interface{}) {
	s.l.Errorf(format, args...)
}

// WithFields will return a new logger based on the original logger with
// the additional supplied fields. Wrapper for logrus Entry.WithFields()
func (s *shim) WithFields(fields log.Fields) log.Logger {
	cp := &shim{
		l: s.l.WithFields(logrus.Fields(fields)),
	}
	return cp
}
