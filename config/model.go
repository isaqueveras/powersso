package config

import (
	"time"
)

// App config struct
type Config struct {
	Server   ServerConfig
	Postgres PGConfig
}

// ServerConfig models a server's configuration data
type ServerConfig struct {
	Version           string
	Port              string
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	SSL               bool
	CSRF              bool
	Debug             bool
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
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
