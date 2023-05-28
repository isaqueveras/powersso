// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package config

import (
	"time"
)

const (
	// modeDevelopment represents the development environment mode
	modeDevelopment string = "dev"
	// modeDevelopment represents the production environment mode
	modeProduction string = "prod"
)

// LoggerEncodingConsole represents the encoding form that the log represents
const LoggerEncodingConsole string = "console"

// type represents the application settings
type
// App config struct
Config struct {
	Meta          MetaConfig     `json:"meta"`
	Server        ServerConfig   `json:"server"`
	Logger        Logger         `json:"logger"`
	Mailer        MailerConfig   `json:"mailer"`
	Database      DatabaseConfig `json:"database"`
	UserAuthToken TokenConfig    `json:"user_auth_token"`
}

// MetaConfig models the meta configuration
type MetaConfig struct {
	ProjectName string `json:"project_name"`
	ProjectURL  string `json:"project_url"`
}

// ServerConfig models a server's configuration data
type ServerConfig struct {
	Version            string        `json:"version"`
	Port               string        `json:"port"`
	PprofPort          string        `json:"pprof_port"`
	Mode               string        `json:"mode"`
	JwtSecretKey       string        `json:"jwt_secret_key"`
	CookieName         string        `json:"cookie_name"`
	AccessLogDirectory string        `json:"access_log_directory"`
	ErrorLogDirectory  string        `json:"error_log_directory"`
	PermissionBase     string        `json:"permission_base"`
	SSL                bool          `json:"ssl"`
	CSRF               bool          `json:"srf"`
	Debug              bool          `json:"debug"`
	StartHTTP          bool          `json:"start_http"`
	StartGRPC          bool          `json:"start_grpc"`
	CtxDefaultTimeout  time.Duration `json:"ctx_default_timeout"`
	ReadTimeout        time.Duration `json:"read_timeout"`
	WriteTimeout       time.Duration `json:"write_timeout"`
}

// DatabaseConfig models postgres database configuration data
type DatabaseConfig struct {
	Host            string        `json:"host"`
	Port            int           `json:"port"`
	User            string        `json:"user"`
	Password        string        `json:"password"`
	Dbname          string        `json:"dbname"`
	Driver          string        `json:"driver"`
	SSLMode         bool          `json:"sslmode"`
	MaxOpenConns    int           `json:"max_open_conns"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	Timeout         int64         `json:"timeout"`
	ConnMaxLifetime time.Duration `json:"conn_max_life_time"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time"`
}

// Logger models the data for the logs configuration
type Logger struct {
	Development       bool   `json:"development"`
	DisableCaller     bool   `json:"disable_caller"`
	DisableStacktrace bool   `json:"disable_stacktrace"`
	Encoding          string `json:"encoding"`
	Level             string `json:"level"`
}

// MailerConfig models the data for the mailer configuration
type MailerConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	TLS      bool   `json:"tls"`
}

// TokenConfig models the data for the token configuration
type TokenConfig struct {
	SecretKey string `json:"secret_key"`
	Duration  int64  `json:"duration"`
}

// IsModeDevelopment returns if in development mode
func (sc *ServerConfig) IsModeDevelopment() bool {
	return sc.Mode == modeDevelopment
}

// IsModeProduction returns if in production mode
func (sc *ServerConfig) IsModeProduction() bool {
	return sc.Mode == modeProduction
}
