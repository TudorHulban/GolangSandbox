package loggerZap

import (
	"go.uber.org/zap"
)

type Config struct {
	// LevDefault Is default log level.
	LevDefault int8
	LevCurrent int8
	pod        int8
}

// should satisy the defined interface
type TLogger struct {
	l zap.SugaredLogger
}

func (el TLogger) Debugf(format string, args ...interface{}) {
	el.l.Debugf(format, args...)
}

func (el TLogger) Infof(format string, args ...interface{}) {
	el.l.Infof(format, args...)
}

func (el TLogger) Warnf(format string, args ...interface{}) {
	el.l.Warnf(format, args...)
}

func (el TLogger) Errorf(format string, args ...interface{}) {
	el.l.Errorf(format, args...)
}

func (el TLogger) Fatalf(format string, args ...interface{}) {
	el.l.Fatalf(format, args...)
}

func (el TLogger) Printf(format string, args ...interface{}) {
	el.l.Infof(format, args...)
}

func NewExtLogger(cfg Config) (*TLogger, error) {
	sugar, err := newZapLogger(cfg)
	if err != nil {
		return &TLogger{}, err
	}

	return &TLogger{
		l: *sugar,
	}, nil
}
