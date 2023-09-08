package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var config *Config

// Config wrapper for all config related fields
type Config struct {
	ServerConfig ServerConfig `yaml:"server"`
	DbConfig     DbConfig     `yaml:"db"`

	initialized bool
}

// ServerConfig all the configs related to server
type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// DbConfig all the configs related to db
type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
}

// GetConfig returns the global config(a singleton).
// If config is not already initilized with New, an error will be returnd
func GetConfig() (*Config, error) {
	if !config.initialized {
		return nil, errors.New("Config not inited. Call New first")
	}

	return config, nil
}

// New Builds a config object from the given path.
// The given file should be yaml file with the structure defined in Config
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
