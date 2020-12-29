package loggerZap

import (
	"errors"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLevels = map[int8]zapcore.Level{-1: zapcore.DebugLevel, 0: zapcore.InfoLevel, 1: zapcore.WarnLevel, 2: zapcore.ErrorLevel, 5: zapcore.FatalLevel}

func toZapLevel(logLevel int8) (zapcore.LevelEnabler, error) {
	if _, exists := zapLevels[logLevel]; !exists {
		return nil, errors.New("log level not found")
	}
	return zapLevels[logLevel], nil
}

func newZapLogger(cfg Config) (*zap.SugaredLogger, error) {
	cores := []zapcore.Core{}
	writer := zapcore.Lock(os.Stdout)

	zapConfig := zap.NewProductionEncoderConfig()
	zapConfig.EncodeTime = zapcore.EpochNanosTimeEncoder

	zaplev, err := toZapLevel(cfg.LevCurrent)
	if err != nil {
		return nil, err
	}
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(zapConfig), writer, zaplev)
	cores = append(cores, core)
	combinedCore := zapcore.NewTee(cores...)

	result := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()
	return result.With("POD", cfg.pod), nil
}
