package zerolog

import (
	"fmt"
	"os"

	log "github.com/InVisionApp/go-logger"
	"github.com/rs/zerolog"
)

type shim struct {
	logger *zerolog.Logger
}

// New can be used to override the default logger.
// Optionally pass in an existing zerolog logger or
// pass in `nil` to use the default logger
func New(logger *zerolog.Logger) log.Logger {
	if logger == nil {
		lg := zerolog.New(os.Stdout).With().Timestamp().Logger()
		logger = &lg
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
	s.logger.Debug().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Info(msg ...interface{}) {
	s.logger.Info().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Warn(msg ...interface{}) {
	s.logger.Warn().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Error(msg ...interface{}) {
	s.logger.Error().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Fatal(msg ...interface{}) {
	s.logger.Fatal().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Panic(msg ...interface{}) {
	s.logger.Panic().Msg(fmt.Sprint(spaceSep(msg)...))
}

/*******************************************************************
*ln funcs
zerolog is a json-only structured logger.
To implement human-readable console logging,
you must initialize the logger using zerolog.ConsoleWriter
as your io.Writer.  Calling a *ln func when zerolog
is in structured logging mode is a no-op
*******************************************************************/

func (s *shim) Debugln(msg ...interface{}) {
	msg = append(msg, "\n")
	s.logger.Debug().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Infoln(msg ...interface{}) {
	msg = append(msg, "\n")
	s.logger.Info().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Warnln(msg ...interface{}) {
	msg = append(msg, "\n")
	s.logger.Warn().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Errorln(msg ...interface{}) {
	msg = append(msg, "\n")
	s.logger.Error().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Fatalln(msg ...interface{}) {
	msg = append(msg, "\n")
	s.logger.Fatal().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Panicln(msg ...interface{}) {
	msg = append(msg, "\n")
	s.logger.Panic().Msg(fmt.Sprint(spaceSep(msg)...))
}

func (s *shim) Debugf(format string, args ...interface{}) {
	s.logger.Debug().Msgf(format, args...)
}

func (s *shim) Infof(format string, args ...interface{}) {
	s.logger.Info().Msgf(format, args...)
}

func (s *shim) Warnf(format string, args ...interface{}) {
	s.logger.Warn().Msgf(format, args...)
}

func (s *shim) Errorf(format string, args ...interface{}) {
	s.logger.Error().Msgf(format, args...)
}

func (s *shim) Fatalf(format string, args ...interface{}) {
	s.logger.Fatal().Msgf(format, args...)
}

func (s *shim) Panicf(format string, args ...interface{}) {
	s.logger.Panic().Msgf(format, args...)
}

// WithFields will return a new logger derived from the original
// zerolog logger, with the provided fields added to the log string,
// as a key-value pair
func (s *shim) WithFields(fields log.Fields) log.Logger {
	lg := s.logger.With().Fields(fields).Logger()
	s.logger = &lg

	return s
}
