package kitlog

import (
	"fmt"
	"os"

	log "github.com/InVisionApp/go-logger"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type shim struct {
	logger kitlog.Logger
}

// New can be used to override the default logger.
// Optionally pass in an existing kitlog logger or
// pass in `nil` to use the default logger.
func New(logger kitlog.Logger) log.Logger {
	if logger == nil {
		logger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stdout))
	}

	return &shim{logger: logger}
}

// this will add a space between all elements in the slice
// this func is needed because fmt.Sprint will not separate
// inputs by a space in all cases, which makes the resulting
// output very hard to read
func spaceSep(a []interface{}) []interface{} {
	aLen := len(a)
	if aLen <= 1 {
		return a
	}

	// we only allocate enough room to add a single space between
	// all elements, so len(a) - 1
	spaceSlice := make([]interface{}, aLen-1)
	// add the empty space to the end of the original slice
	a = append(a, spaceSlice...)

	// stagger the values.  this will leave an empty slot between all
	// values to be filled with a space
	for i := aLen - 1; i > 0; i-- {
		a[i+i] = a[i]
		a[i+i-1] = " "
	}

	return a
}

func (s *shim) Debug(msg ...interface{}) {
	level.Debug(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Info(msg ...interface{}) {
	level.Info(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Warn(msg ...interface{}) {
	level.Warn(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Error(msg ...interface{}) {
	level.Error(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Fatal(msg ...interface{}) {
	// Since kitlog does not support fatal, emulate it the best we can
	msg = append([]interface{}{"[FATAL]"}, msg...)
	level.Error(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
	os.Exit(1)
}

func (s *shim) Panic(msg ...interface{}) {
	// Since kitlog does not support panic, emulate it the best we can
	msg = append([]interface{}{"[PANIC]"}, msg...)
	level.Error(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
	panic(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Debugln(msg ...interface{}) {
	level.Debug(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Infoln(msg ...interface{}) {
	level.Info(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Warnln(msg ...interface{}) {
	level.Warn(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Errorln(msg ...interface{}) {
	level.Error(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Fatalln(msg ...interface{}) {
	// Since kitlog does not support fatal, emulate it the best we can
	msg = append([]interface{}{"[FATAL]"}, msg...)
	msg = append(msg, "\n")
	level.Error(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
	os.Exit(1)
}

func (s *shim) Panicln(msg ...interface{}) {
	// Since kitlog does not support panic, emulate it the best we can
	msg = append([]interface{}{"[PANIC]"}, msg...)
	msg = append(msg, "\n")
	level.Error(s.logger).Log("msg", fmt.Sprint(spaceSep(msg)...))
	panic(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Debugf(format string, args ...interface{}) {
	level.Debug(s.logger).Log("msg", fmt.Sprintf(format, args...))
}

func (s *shim) Infof(format string, args ...interface{}) {
	level.Info(s.logger).Log("msg", fmt.Sprintf(format, args...))
}

func (s *shim) Warnf(format string, args ...interface{}) {
	level.Warn(s.logger).Log("msg", fmt.Sprintf(format, args...))
}

func (s *shim) Errorf(format string, args ...interface{}) {
	level.Error(s.logger).Log("msg", fmt.Sprintf(format, args...))
}

func (s *shim) Fatalf(format string, args ...interface{}) {
	// Since kitlog does not support fatal, emulate it the best we can
	format = "[FATAL] " + format
	level.Error(s.logger).Log("msg", fmt.Sprintf(format, args...))
	os.Exit(1)
}

func (s *shim) Panicf(format string, args ...interface{}) {
	// Since kitlog does not support panic, emulate it the best we can
	format = "[PANIC] " + format
	level.Error(s.logger).Log("msg", fmt.Sprintf(format, args...))
	panic(fmt.Sprintf(format, args...))
}

// WithFields will return a new logger derived from the original
// kitlog logger, with the provided fields added to the log string,
// as a key-value pair
func (s *shim) WithFields(fields log.Fields) log.Logger {
	var keyvals []interface{}

	for key, value := range fields {
		keyvals = append(keyvals, key, value)
	}

	return &shim{
		logger: kitlog.With(s.logger, keyvals...),
	}
}
