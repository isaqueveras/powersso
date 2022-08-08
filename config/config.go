// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *Config

// LoadConfig config file from given path
func LoadConfig() {
	v := viper.New()

	v.SetConfigName("./config/config-local")
	v.AddConfigPath(".")
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
