package logger

import (
	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	cfg           *config.Config
	sugaredLogger *zap.SugaredLogger
}

// NewLogger logger constructor
func NewLogger(cfg *config.Config) *logger {
	return &logger{cfg: cfg}
}

// InitLogger initialize zap logger
func (l *logger) InitLogger() (*zap.Logger, error) {
	var cfg zap.Config

	if l.cfg.Server.Mode == "production" {
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

	return cfg.Build()
}
