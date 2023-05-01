// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var config *Config

// LoadConfig config file from given path
func LoadConfig() {
	path := "../app.json"
	if val, set := os.LookupEnv("CONFIG_POWER_SSO"); set && val != "" {
		path = val
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
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
