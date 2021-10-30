package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	vipr *viper.Viper
}

func InitConfig(env string) (IConfig, error) {
	conf := viper.New()
	conf.SetConfigType("json")
	conf.SetConfigName(env)
	conf.AddConfigPath("../config/")
	conf.AddConfigPath("config/")
	conf.AddConfigPath(".")
	err := conf.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		fmt.Println("Config File not found.")
		return nil, errors.New(env + " not found in directory")
	} else if err != nil {
		fmt.Println("Error:", err.Error())
	}
	return &Config{
		vipr: conf,
	}, nil
}

// GetString return value as string
func (conf *Config) GetString(key string) string {
	return conf.vipr.GetString(key)
}

// GetInt64 return value as Int64
func (conf *Config) GetInt64(key string) int64 {
	return conf.vipr.GetInt64(key)
}

// GetBool return value as Boolean
func (conf *Config) GetBool(key string) bool {
	return conf.vipr.GetBool(key)
}
