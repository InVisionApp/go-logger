package testlog

import (
	"bytes"
	"fmt"

	"github.com/InVisionApp/go-logger"
)

type TestLogger struct {
	buf    *bytes.Buffer
	fields map[string]interface{}
}

func NewTestLog() *TestLogger {
	b := &bytes.Buffer{}

	return &TestLogger{
		buf: b,
	}
}

// Bytes returns all bytes of the log buffer
func (t *TestLogger) Bytes() []byte {
	return t.buf.Bytes()
}

// Reset the log buffer
func (t *TestLogger) Reset() {
	t.buf.Reset()
}

// Debug log message
func (t *TestLogger) Debug(msg ...interface{}) {
	t.buf.WriteString(fmt.Sprintf("[DEBUG] %s %s", fmt.Sprint(msg...), pretty(t.fields)+"\n"))
}

// Info log message
func (t *TestLogger) Info(msg ...interface{}) {
	t.buf.WriteString(fmt.Sprintf("[INFO] %s %s", fmt.Sprint(msg...), pretty(t.fields)+"\n"))
}

// Warn log message
func (t *TestLogger) Warn(msg ...interface{}) {
	t.buf.WriteString(fmt.Sprintf("[WARN] %s %s", fmt.Sprint(msg...), pretty(t.fields)+"\n"))
}

// Error log message
func (t *TestLogger) Error(msg ...interface{}) {
	t.buf.WriteString(fmt.Sprintf("[ERROR] %s %s", fmt.Sprint(msg...), pretty(t.fields)+"\n"))
}

// Debugf log message with formatting
func (t *TestLogger) Debugf(format string, args ...interface{}) {
	t.buf.WriteString(fmt.Sprintf("[DEBUG] "+format, args...) + " " + pretty(t.fields) + "\n")
}

// Infof log message with formatting
func (t *TestLogger) Infof(format string, args ...interface{}) {
	t.buf.WriteString(fmt.Sprintf("[INFO] "+format, args...) + " " + pretty(t.fields) + "\n")
}

// Warnf log message with formatting
func (t *TestLogger) Warnf(format string, args ...interface{}) {
	t.buf.WriteString(fmt.Sprintf("[WARN] "+format, args...) + " " + pretty(t.fields) + "\n")
}

// Errorf log message with formatting
func (t *TestLogger) Errorf(format string, args ...interface{}) {
	t.buf.WriteString(fmt.Sprintf("[ERROR] "+format, args...) + " " + pretty(t.fields) + "\n")
}

// WithFields will return a new logger based on the original logger
// with the additional supplied fields
func (t *TestLogger) WithFields(fields log.Fields) log.Logger {
	cp := &TestLogger{
		buf: t.buf,
	}

	if t.fields == nil {
		cp.fields = fields
		return cp
	}

	cp.fields = make(map[string]interface{}, len(t.fields)+len(fields))
	for k, v := range t.fields {
		cp.fields[k] = v
	}

	for k, v := range fields {
		cp.fields[k] = v
	}

	return cp
}

// helper for pretty printing of fields
func pretty(m map[string]interface{}) string {
	if len(m) < 1 {
		return ""
	}

	s := ""
	for k, v := range m {
		s += fmt.Sprintf("%s=%v ", k, v)
	}

	return s[:len(s)-1]
}
