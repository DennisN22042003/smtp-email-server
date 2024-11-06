package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Server struct {
		PlainPort          string `yaml:"plain_port"`
		TLSPort            string `yaml:"tls_port"`
		TLSEnabled         bool   `yaml:"tls_enabled"`
		CertFile           string `yaml:"cert_file"`
		KeyFile            string `yaml:"key_file"`
		SMTPAuthentication struct {
			Enabled bool `yaml:"enabled"`
			Users   []struct {
				Email    string `yaml:"email"`
				Password string `yaml:"password"`
			} `yaml:"users"`
		} `yaml:"smtp_authentication"`
	} `yaml:"Server"`
	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
}

var config Config

func loadConfig() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}
