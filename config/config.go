// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var config *Config

// LoadConfig config file from given path
func LoadConfig(path ...string) {
	v := viper.New()

	if path == nil {
		path[0] = "."
	}

	v.SetConfigName(getConfigPath(os.Getenv("CONFIG_POWER_SSO")))
	v.AddConfigPath(path[0])
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found")
		}
		log.Fatal(err)
	}

	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

// Get returns a pointer to a
// Config struct which holds a valid config
func Get() *Config {
	if config == nil {
		log.Fatal("config was not successfully loaded")
	}
	return config
}

// Get config path for dev or production
func getConfigPath(configPath string) string {
	if configPath == modeProduction {
		return "./config/config-prod"
	}
	return "./config/config-dev"
}
