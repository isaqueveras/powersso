// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
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
)

const (
	// LoggerEncodingConsole represents the encoding form that the log represents
	LoggerEncodingConsole string = "console"
)

// App config struct
type Config struct {
	Server   ServerConfig
	Postgres PGConfig
	Redis    RedisConfig
	Logger   Logger
}

// ServerConfig models a server's configuration data
type ServerConfig struct {
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
}

// PGConfig models postgres database configuration data
type PGConfig struct {
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
type RedisConfig struct {
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
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}
