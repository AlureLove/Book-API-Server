package config

import (
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"os"
)

var conf = &Config{
	App: &App{
		Host: "127.0.0.1",
		Port: 8080,
	},
	MySQL: &MySQL{
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
		Username: "root",
		Password: "123456",
		Debug:    true,
	},
	Log: &Log{
		CallerDeep: 3,
		Level:      zerolog.DebugLevel,
		Console: Console{
			Enable:  true,
			NoColor: true,
		},
		File: File{
			Enable:     true,
			MaxSize:    100,
			MaxBackups: 6,
		},
	},
}

func Get() *Config {
	if conf == nil {
		panic("configuration not initialized")
	}

	return conf
}

func LoadConfigFromYaml(configFilePath string) error {
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, &conf)
}
