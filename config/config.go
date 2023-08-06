package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type JWTConfig struct {
	SecretKey       string `yaml:"secret_key"`
	ExpirationHours int    `yaml:"expiration_hours"`
}

type AuthConfig struct {
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
	SSLMode    string `yaml:"ssl_mode"`
	Port       string `yaml:"port"`
}

type TodoConfig struct {
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
	SSLMode    string `yaml:"ssl_mode"`
	Port       string `yaml:"port"`
}

type Config struct {
	JWT  JWTConfig  `yaml:"jwt"`
	Auth AuthConfig `yaml:"auth"`
	Todo TodoConfig `yaml:"todo"`
}

func LoadConfig() (Config, error) {
	var config Config
	file, err := os.Open("config/config.yaml")
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
