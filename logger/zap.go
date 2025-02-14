package logger

import (
	core "github.com/hifat/goroger-core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	log *zap.Logger
}

func New(log *zap.Logger) core.Logger {
	return &zapLog{log}
}

func (z *zapLog) Skip(skip int) core.Logger {
	return &zapLog{
		log: z.log.WithOptions(zap.AddCallerSkip(skip)),
	}
}

func (z *zapLog) handleField(fields []core.LoggerField) []zap.Field {
	zfields := []zap.Field{}

	for _, f := range fields {
		zfields = append(zfields, zap.Field{
			Key:       f.Key,
			Type:      zapcore.FieldType(f.Type),
			Integer:   f.Integer,
			String:    f.String,
			Interface: f.Interface,
		})
	}

	return zfields
}

func (z *zapLog) Info(message string, fields ...core.LoggerField) {
	z.log.Info(message, z.handleField(fields)...)
}

func (z *zapLog) Debug(message string, fields ...core.LoggerField) {
	z.log.Debug(message, z.handleField(fields)...)
}

func (z *zapLog) Warn(message string, fields ...core.LoggerField) {
	z.log.Warn(message, z.handleField(fields)...)
}

func (z *zapLog) Fatal(message string, fields ...core.LoggerField) {
	z.log.Fatal(message, z.handleField(fields)...)
}

func (z *zapLog) Error(message any, fields ...core.LoggerField) {
	f := z.handleField(fields)

	switch v := message.(type) {
	case error:
		z.log.Error(v.Error(), f...)
	case string:
		z.log.Error(v, f...)
	}
}
