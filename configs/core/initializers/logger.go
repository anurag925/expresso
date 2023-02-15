package initializers

import (
	"go.uber.org/zap"
)

type Logger interface {
	With(...any) Logger
	Info(...any)
	Infof(template string, args ...any)
	Error(fields ...any)
	Errorf(template string, args ...any)
	Debug(...any)
	Warn(...any)
	LogAndShutdown(...any)
	Fatal(...any)
}

type FileLogger struct {
	zap     *zap.Logger
	sugared *zap.SugaredLogger
	with    []any
}

var _ Logger = (*FileLogger)(nil)

func NewFileLogger(zap *zap.Logger, sugared *zap.SugaredLogger, with ...any) *FileLogger {
	return &FileLogger{zap, sugared, with}
}

func (f *FileLogger) ZapLogger() *zap.Logger {
	return f.zap
}

func (f *FileLogger) Info(fields ...any) {
	f.sugared.With(f.with...).Info(fields...)
}

func (f *FileLogger) Infof(template string, args ...any) {
	f.sugared.With(f.with...).Infof(template, args...)
}

func (f *FileLogger) Error(fields ...any) {
	f.sugared.With(f.with...).Error(fields...)
}

func (f *FileLogger) Errorf(template string, args ...any) {
	f.sugared.With(f.with...).Infof(template, args...)
}

func (f *FileLogger) Debug(fields ...any) {
	f.sugared.With(f.with...).Debug(fields...)
}

func (f *FileLogger) Debugf(template string, args ...any) {
	f.sugared.With(f.with...).Debugf(template, args...)
}

func (f *FileLogger) LogAndShutdown(fields ...any) {
	f.sugared.With(f.with...).Fatal(fields...)
}

// same as log and shutdown but do not use
func (f *FileLogger) Fatal(fields ...any) {
	f.sugared.With(f.with...).Fatal(fields...)
}

func (f *FileLogger) Warn(fields ...any) {
	f.sugared.With(f.with...).Warn(fields...)
}

func (f *FileLogger) With(args ...any) Logger {
	f.with = args
	return f
}
