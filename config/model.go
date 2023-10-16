// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package config

import "time"

const (
	// modeDevelopment represents the development environment mode
	modeDevelopment string = "dev"
	// modeDevelopment represents the production environment mode
	modeProduction string = "prod"
)

// Config type represents the application settings
type Config struct {
	ProjectName string         `json:"project_name"`
	Server      ServerConfig   `json:"server"`
	Database    DatabaseConfig `json:"database"`

	SecretsDuration int64   `json:"secrets_duration"`
	SecretsTokens   Secrets `json:"secrets_tokens"`
}

// GetSecrets returns a list of tokens
func (c *Config) GetSecrets() (keys []string) {
	return []string{c.SecretsTokens.User, c.SecretsTokens.Admin, c.SecretsTokens.Integration}
}

// ServerConfig models a server's configuration data
type ServerConfig struct {
	Version                  string        `json:"version"`
	Port                     string        `json:"port"`
	PprofPort                string        `json:"pprof_port"`
	Mode                     string        `json:"mode"`
	CookieName               string        `json:"cookie_name"`
	AccessLogDirectory       string        `json:"access_log_directory"`
	ErrorLogDirectory        string        `json:"error_log_directory"`
	PermissionBase           string        `json:"permission_base"`
	AccessControlAllowOrigin string        `json:"access_control_allow_origin"`
	SSL                      bool          `json:"ssl"`
	CSRF                     bool          `json:"srf"`
	Debug                    bool          `json:"debug"`
	StartHTTP                bool          `json:"start_http"`
	StartGRPC                bool          `json:"start_grpc"`
	CtxDefaultTimeout        time.Duration `json:"ctx_default_timeout"`
	ReadTimeout              time.Duration `json:"read_timeout"`
	WriteTimeout             time.Duration `json:"write_timeout"`
}

// DatabaseConfig models postgres database configuration data
type DatabaseConfig struct {
	Host            string        `json:"host"`
	Port            int           `json:"port"`
	User            string        `json:"user"`
	Password        string        `json:"password"`
	Name            string        `json:"dbname"`
	SSLMode         bool          `json:"sslmode"`
	MaxOpenConns    int           `json:"max_open_conns"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	Timeout         int64         `json:"timeout"`
	ConnMaxLifetime time.Duration `json:"conn_max_life_time"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time"`
}

// Secrets models the data for the token configuration
type Secrets struct {
	User        string `json:"user"`
	Admin       string `json:"admin"`
	Integration string `json:"integration"`
}

// IsModeDevelopment returns if in development mode
func (sc *ServerConfig) IsModeDevelopment() bool {
	return sc.Mode == modeDevelopment
}

// IsModeProduction returns if in production mode
func (sc *ServerConfig) IsModeProduction() bool {
	return sc.Mode == modeProduction
}
