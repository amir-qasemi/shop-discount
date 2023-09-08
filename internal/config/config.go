package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var config *Config

type Config struct {
	ServerConfig ServerConfig `yaml:"server"`
	DbConfig     DbConfig     `yaml:"db"`

	initialized bool
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
}

func GetConfig() (*Config, error) {
	if !config.initialized {
		//log.P("Config is not initialized!")
	}

	return config, nil
}

func New(configFilePath string) (*Config, error) {
	f, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	config.initialized = true
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Config is initialized from path: %s", configFilePath)

	return config, nil
}
