package bitconfig

import (
	"time"
)

// Config holds all application configuration values
type Config struct {
	Env            string
	IP             string
	Port           uint
	App            string
	DBHost         string
	DBPort         uint
	DBReadTimeout  time.Duration
	DBWriteTimeout time.Duration
	Static         string
	Logger         string
}

// ConfigMap holds the configuration data for each given environment
var ConfigMap = map[string]Config{
	"dev": Config{
		Env:            "dev",
		IP:             "0.0.0.0",
		Port:           8080,
		App:            "bitChat",
		DBHost:         "0.0.0.0",
		DBPort:         27017,
		DBReadTimeout:  5 * time.Second,
		DBWriteTimeout: 5 * time.Second,
		Static:         "/public/",
		Logger:         "local",
	},
}
