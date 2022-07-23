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
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	SSLMode  bool
	PgDriver string
}

// Logger models the data for the logs configuration
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}
