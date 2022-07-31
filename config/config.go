package config

import (
	"log"
)

type Config struct {
	Port string
	Env  string
}

type Application struct {
	Config Config
	Logger *log.Logger
}
