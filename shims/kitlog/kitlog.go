package kitlog

import (
	"fmt"
	"os"

	"github.com/InVisionApp/go-logger"
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

// WithFields will return a new logger derived from the original
// kitlog logger, with the provided fields added to the log string,
// as a key-value pair
func (s *shim) WithFields(fields log.Fields) log.Logger {
	keyvals := make([]interface{}, len(fields)*2)
	i := 0

	for key, value := range fields {
		keyvals[i] = key
		keyvals[i+1] = value

		i += 2
	}

	return &shim{
		logger: kitlog.With(s.logger, keyvals...),
	}
}
