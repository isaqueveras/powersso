// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/middleware"
)

var logg *zap.SugaredLogger

type Logger struct {
	cfg *config.Config
}

// NewLogger logger constructor
func NewLogger(cfg *config.Config) *Logger {
	return &Logger{cfg: cfg}
}

// InitLogger initialize zap logger
func (l *Logger) InitLogger() {
	var cfg zap.Config

	if l.cfg.Server.Mode != config.ModeDevelopment {
		cfg = zap.NewProductionConfig()
		cfg.DisableStacktrace = true
		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		cfg.OutputPaths = []string{l.cfg.Server.AccessLogDirectory, "stdout"}
		cfg.ErrorOutputPaths = []string{l.cfg.Server.ErrorLogDirectory, "stderr"}
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

	logs, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	logger := zap.New(logs.Core(), zap.AddCaller(), zap.AddCallerSkip(1))
	logg = logger.Sugar()
	defer func() { _ = logg.Sync() }()
}

func (l *Logger) ZapLogger() *zap.Logger {
	return logg.Desugar()
}

func (l *Logger) Debug(args ...interface{}) {
	logg.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	logg.Debugf(template, args...)
}

func (l *Logger) Info(args ...interface{}) {
	logg.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	logg.Infof(template, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	logg.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	logg.Warnf(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	logg.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	logg.Errorf(template, args...)
}

func (l *Logger) DPanic(args ...interface{}) {
	logg.DPanic(args...)
}

func (l *Logger) DPanicf(template string, args ...interface{}) {
	logg.DPanicf(template, args...)
}

func (l *Logger) Panic(args ...interface{}) {
	logg.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	logg.Panicf(template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	logg.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	logg.Fatalf(template, args...)
}

// PanicRecovery handles recovered panics
func PanicRecovery(msg interface{}) (err error) {
	logg.Error("PANIC detected: ", msg)
	return
}
