package loggers

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	slog *zap.SugaredLogger
}

func NewZapLogger() (*ZapLogger, error) {
	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
		},
	}
	cfg.EncoderConfig = ecszap.ECSCompatibleEncoderConfig(cfg.EncoderConfig)
	logger, _ := cfg.Build(ecszap.WrapCoreOption(), zap.AddCaller())

	return &ZapLogger{
		slog: logger.Sugar(),
	}, nil
}

func (zl *ZapLogger) GetName() string {
	return "zap-logger"
}

func (zl *ZapLogger) PlainInfo(message string) {
	zl.slog.Infow(message, zap.String("eventSeverity", "Info"))
}

func (zl *ZapLogger) Info(message string, args ...interface{}) {
	zl.slog.Infow(message,
		zap.String("eventSeverity", "Info"),
		zap.Any("metadata", args),
	)
}

func (zl *ZapLogger) Warning(message string, args ...interface{}) {
	zl.slog.Warnw(message,
		zap.String("eventSeverity", "Warning"),
		zap.Any("metadata", args),
	)
}

func (zl *ZapLogger) Error(message string, args ...interface{}) {
	zl.slog.Errorw(message,
		zap.String("eventSeverity", "Error"),
		zap.Any("metadata", args),
	)
}

func (zl *ZapLogger) PlainError(message string) {
	zl.slog.Error(message, zap.String("eventSeverity", "Error"))
}

func (zl *ZapLogger) Debug(message string, args ...interface{}) {
	zl.slog.Debugw(message,
		zap.String("eventSeverity", "Info"),
		zap.Any("metadata", args),
	)
}

func (zl *ZapLogger) Close() error {
	return zl.slog.Sync()
}
