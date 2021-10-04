package loggers

import "go.uber.org/zap"

type Logger interface {
	info(message string, platform string, boundary, name string, args ...interface{})
	warning(message string, platform string, boundary, name string, args ...interface{})
	error(message string, platform string, boundary, name string, args ...interface{})
	debug(message string, platform string, boundary, name string, args ...interface{})
}

type ZapLogger struct {
	slog *zap.SugaredLogger
}

func NewZapLogger() (*ZapLogger, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return &ZapLogger{
		slog: logger.Sugar(),
	}, nil
}

func (zl *ZapLogger) info(message string, platform string, boundary, name string, args ...interface{}) {
	zl.slog.Infow(message,
		zap.String("platform", platform),
		zap.String("boundary", boundary),
		zap.String("name", name),
		args,
	)
}
func (zl *ZapLogger) warning(message string, platform string, boundary, name string, args ...interface{}) {
	zl.slog.Warnw(message,
		zap.String("platform", platform),
		zap.String("boundary", boundary),
		zap.String("name", name),
		args,
	)
}
func (zl *ZapLogger) error(message string, platform string, boundary, name string, args ...interface{}) {
	zl.slog.Errorw(message,
		zap.String("platform", platform),
		zap.String("boundary", boundary),
		zap.String("name", name),
		args,
	)
}
func (zl *ZapLogger) debug(message string, platform string, boundary, name string, args ...interface{}) {
	zl.slog.Debugw(message,
		zap.String("platform", platform),
		zap.String("boundary", boundary),
		zap.String("name", name),
		args,
	)
}
