package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

// LoadConfig config file from given path
func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName("./config/config-local")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig parse config file the of aplication
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	var err error
	if err = v.Unmarshal(&c); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
