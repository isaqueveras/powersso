// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/middleware"
)

type logger struct {
	cfg  *config.Config
	logg *zap.SugaredLogger
}

// NewLogger logger constructor
func NewLogger(cfg *config.Config) *logger {
	return &logger{cfg: cfg}
}

// InitLogger initialize zap logger
func (l *logger) InitLogger() {
	var cfg zap.Config

	if l.cfg.Server.Mode == config.ModeDevelopment {
		cfg = zap.NewProductionConfig()
		cfg.DisableStacktrace = true
		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	cfg.Encoding = "json"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.NameKey = "name"
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.StacktraceKey = "stack_trace"
	cfg.InitialFields = map[string]interface{}{
		"application": "power-sso",
		"version":     middleware.Version,
	}

	cfg.OutputPaths = []string{l.cfg.Server.AccessLogDirectory, "stdout"}
	cfg.ErrorOutputPaths = []string{l.cfg.Server.ErrorLogDirectory, "stderr"}

	logs, err := cfg.Build()
	if err != nil {
		l.logg.Error(err)
	}

	logger := zap.New(logs.Core(), zap.AddCaller(), zap.AddCallerSkip(1))
	l.logg = logger.Sugar()
	if err := l.logg.Sync(); err != nil {
		l.logg.Error(err)
	}
}

func (l *logger) ZapLogger() *zap.Logger {
	return l.logg.Desugar()
}

func (l *logger) Debug(args ...interface{}) {
	l.logg.Debug(args...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.logg.Debugf(template, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.logg.Info(args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.logg.Infof(template, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.logg.Warn(args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.logg.Warnf(template, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.logg.Error(args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.logg.Errorf(template, args...)
}

func (l *logger) DPanic(args ...interface{}) {
	l.logg.DPanic(args...)
}

func (l *logger) DPanicf(template string, args ...interface{}) {
	l.logg.DPanicf(template, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.logg.Panic(args...)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.logg.Panicf(template, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.logg.Fatal(args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.logg.Fatalf(template, args...)
}
