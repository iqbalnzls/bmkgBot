package config

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/bmkgBot/internal/shared/rest_client"
)

type Config struct {
	Telebot *TelebotConfig `yaml:"telebot" mapstructure:"telebot"`
	BMKG    *BMKGConfig    `yaml:"bmkg" mapstructure:"bmkg"`
}

type TelebotConfig struct {
	APIKey string `yaml:"api_key" mapstructure:"api_key"`
}

type BMKGConfig struct {
	Options rest_client.Options `yaml:"options" mapstructure:"options"`
	Path    map[string]string   `yaml:"path" mapstructure:"path"`
}

func NewConfig(path string) *Config {
	fmt.Println("Try NewConfig ... ")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var conf Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
