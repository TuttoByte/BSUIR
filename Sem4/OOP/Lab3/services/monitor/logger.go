package monitor

import "go.uber.org/zap"

type CustomLogger struct {
	lg *zap.SugaredLogger
}

func NewCustomLogger() *CustomLogger {
	logger, _ := zap.NewDevelopment()
	return &CustomLogger{lg: logger.Sugar()}
}

func (l *CustomLogger) Info(name string, args ...interface{}) {
	l.lg.Named(name).Info(args...)
}

func (l *CustomLogger) Error(name string, args ...interface{}) {
	l.lg.Named(name).Error(args...)
}

func (l *CustomLogger) Debug(name string, args ...interface{}) {
	l.lg.Named(name).Debug(args...)
}

func (l *CustomLogger) Warn(name string, args ...interface{}) {
	l.lg.Named(name).Warn(args...)
}
