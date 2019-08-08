package testlog

import (
	"bytes"
	"fmt"
	"os"

	log "github.com/InVisionApp/go-logger"
)

// TestLogger is used to capture logs during the execution of a test.
// It writes the logs to a byte buffer which can be dumped and
// inspected. It also tracks a call count of the total number of
// times the logger has been called. Note that this logger is not
// meant to be used in production. It is meant only for tests.
type TestLogger struct {
	buf    *bytes.Buffer
	count  *counter
	fields map[string]interface{}
}

// NewTestLog generates a new TestLogger
func New() *TestLogger {
	b := &bytes.Buffer{}

	return &TestLogger{
		buf:   b,
		count: newCounter(),
	}
}

// Bytes returns all bytes of the log buffer
func (t *TestLogger) Bytes() []byte {
	return t.buf.Bytes()
}

// CallCount returns the number of times this logger was called
func (t *TestLogger) CallCount() int {
	return t.count.val()
}

// Reset the log buffer and call count
func (t *TestLogger) Reset() {
	t.buf.Reset()
	t.count.reset()
}

func (t *TestLogger) write(level, msg string) {
	t.buf.WriteString(fmt.Sprintf("[%s] %s %s", level, msg, pretty(t.fields)+"\n"))
	t.count.inc()
}

//Debugln log line message
func (t *TestLogger) Debugln(msg ...interface{}) {
	a := fmt.Sprintln(msg...)
	t.write("DEBUG", a[:len(a)-1])
}

//Infoln log line message
func (t *TestLogger) Infoln(msg ...interface{}) {
	a := fmt.Sprintln(msg...)
	t.write("INFO", a[:len(a)-1])
}

//Warnln log line message
func (t *TestLogger) Warnln(msg ...interface{}) {
	a := fmt.Sprintln(msg...)
	t.write("WARN", a[:len(a)-1])
}

//Errorln log line message
func (t *TestLogger) Errorln(msg ...interface{}) {
	a := fmt.Sprintln(msg...)
	t.write("ERROR", a[:len(a)-1])
}

//Fatalln log line message
func (t *TestLogger) Fatalln(msg ...interface{}) {
	a := fmt.Sprintln(msg...)
	t.write("FATAL", a[:len(a)-1])
	os.Exit(1)
}

// Debug log message
func (t *TestLogger) Debug(msg ...interface{}) {
	t.write("DEBUG", fmt.Sprint(msg...))
}

// Info log message
func (t *TestLogger) Info(msg ...interface{}) {
	t.write("INFO", fmt.Sprint(msg...))
}

// Warn log message
func (t *TestLogger) Warn(msg ...interface{}) {
	t.write("WARN", fmt.Sprint(msg...))
}

// Error log message
func (t *TestLogger) Error(msg ...interface{}) {
	t.write("ERROR", fmt.Sprint(msg...))
}

// Fatal log message (and exit)
func (t *TestLogger) Fatal(msg ...interface{}) {
	t.write("FATAL", fmt.Sprint(msg...))
	os.Exit(1)
}

// Panic log message (and exit)
func (t *TestLogger) Panic(msg ...interface{}) {
	t.write("PANIC", fmt.Sprint(msg...))
	panic(fmt.Sprint(msg...))
}

// Debugf log message with formatting
func (t *TestLogger) Debugf(format string, args ...interface{}) {
	t.write("DEBUG", fmt.Sprintf(format, args...))
}

// Infof log message with formatting
func (t *TestLogger) Infof(format string, args ...interface{}) {
	t.write("INFO", fmt.Sprintf(format, args...))
}

// Warnf log message with formatting
func (t *TestLogger) Warnf(format string, args ...interface{}) {
	t.write("WARN", fmt.Sprintf(format, args...))
}

// Errorf log message with formatting
func (t *TestLogger) Errorf(format string, args ...interface{}) {
	t.write("ERROR", fmt.Sprintf(format, args...))
}

// Fatalf log message with formatting
func (t *TestLogger) Fatalf(format string, args ...interface{}) {
	t.write("FATAL", fmt.Sprintf(format, args...))
}

// WithFields will return a new logger based on the original logger
// with the additional supplied fields
func (t *TestLogger) WithFields(fields log.Fields) log.Logger {
	cp := &TestLogger{
		buf:   t.buf,
		count: t.count,
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
