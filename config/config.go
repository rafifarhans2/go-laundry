package config

import (
	"fmt"
	"os"
)

type Dbconfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type Config struct {
	Dbconfig
}

func (c *Config) ReadConfig() error {
	c.Dbconfig = Dbconfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	if c.Dbconfig.Host == "" || c.Dbconfig.Port == "" || c.Dbconfig.Name == "" || c.Dbconfig.User == "" || c.Dbconfig.Password == "" || c.Dbconfig.Driver == "" {
		return fmt.Errorf("missing required enviorment variable")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
