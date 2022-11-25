// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package config

import (
	"time"
)

const (
	// ModeDevelopment represents the development environment mode
	ModeDevelopment string = "development"
	// ModeDevelopment represents the production environment mode
	ModeProduction string = "production"
	// ModeDevelopment represents the testing environment mode
	ModeTesting string = "testing"
)

const (
	// LoggerEncodingConsole represents the encoding form that the log represents
	LoggerEncodingConsole string = "console"
)

// type represents the application settings
type (
	// App config struct
	Config struct {
		Meta MetaConfig

		Server ServerConfig
		Logger Logger
		Mailer MailerConfig

		Postgres PGConfig
		Redis    RedisConfig

		UserAuthToken TokenConfig
	}

	// MetaConfig models the meta configuration
	MetaConfig struct {
		ProjectName string
		ProjectURL  string
	}

	// ServerConfig models a server's configuration data
	ServerConfig struct {
		Version            string
		Port               string
		PprofPort          string
		Mode               string
		JwtSecretKey       string
		CookieName         string
		AccessLogDirectory string
		ErrorLogDirectory  string
		PermissionBase     string
		SSL                bool
		CSRF               bool
		Debug              bool
		CtxDefaultTimeout  time.Duration
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
	}

	// PGConfig models postgres database configuration data
	PGConfig struct {
		Host            string
		Port            string
		User            string
		Password        string
		Dbname          string
		SSLMode         bool
		Driver          string
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifetime time.Duration
		ConnMaxIdleTime time.Duration
		Timeout         int64
	}

	// RedisConfig models redis database configuration data
	RedisConfig struct {
		RedisAddr      string
		RedisPassword  string
		RedisDB        string
		RedisDefaultdb string
		MinIdleConns   int
		PoolSize       int
		PoolTimeout    int
		Password       string
		DB             int
	}

	// Logger models the data for the logs configuration
	Logger struct {
		Development       bool
		DisableCaller     bool
		DisableStacktrace bool
		Encoding          string
		Level             string
	}

	// MailerConfig models the data for the mailer configuration
	MailerConfig struct {
		Host     string
		Port     int
		Email    string
		Username string
		Password string
		TLS      bool
	}

	// TokenConfig models the data for the token configuration
	TokenConfig struct {
		SecretKey string
		Duration  int64
	}
)
