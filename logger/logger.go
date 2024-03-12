package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog"
)

// Interface -.
type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

var (
	instance *Logger
	once     sync.Once
)

// Logger -.
type Logger struct {
	logger *zerolog.Logger
	once   sync.Once
}

var _ Interface = (*Logger)(nil)
// func New() *Logger {
// 	if instance == nil {
// 		instance = &Logger{}
// 	}
// 	return instance
// }

// func (l *Logger) Initialize() {
// 	l.once.Do(func() {
// 		zerolog.SetGlobalLevel(zerolog.InfoLevel)

// 		skipFrameCount := 3
// 		logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()
// 		l.logger = &logger
// 	})
// }

// New -.
func New() *Logger {
	
	
	once.Do(func() {
		// switch strings.ToLower(level) {
		// case "error":
		// 	l = zerolog.ErrorLevel
		// case "warn":
		// 	l = zerolog.WarnLevel
		// case "info":
		// 	l = zerolog.InfoLevel
		// case "debug":
		// 	l = zerolog.DebugLevel
		// default:
		// 	l = zerolog.InfoLevel
		// }

		zerolog.SetGlobalLevel(zerolog.InfoLevel)

		skipFrameCount := 3
		logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

		instance = &Logger{
			logger: &logger,
		}
	})
	return instance
}

func (log *Logger) SetLevel(level string) {
	
	var l zerolog.Level
	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel

	}

	zerolog.SetGlobalLevel(l)
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.msg(zerolog.DebugLevel, message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {

	//l.msg("info", message, args...)
	l.log(zerolog.InfoLevel, message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.log(zerolog.WarnLevel, message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	/*if l.logger.GetLevel() <= zerolog.DebugLevel {
		l.Debug(message, args...)
	}*/

	l.msg(zerolog.ErrorLevel, message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg(zerolog.FatalLevel, message, args...)

	os.Exit(1)
}

func (l *Logger) log(level zerolog.Level, message string, args ...interface{}) {

	//fmt.Println("log level:", l.logger.GetLevel())
	loggers := l.logger.WithLevel(level)
	if len(args) == 0 {
		loggers.Msg(message)
		//l.logger.Info().Msg(message)
	} else {
		loggers.Msgf(message, args...)
		//l.logger.Info().Msgf(message, args...)
	}

}

func (l *Logger) msg(level zerolog.Level, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(level, msg.Error(), args...)
	case string:
		l.log(level, msg, args...)
	default:
		l.log(zerolog.InfoLevel, fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
