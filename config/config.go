package config

import (
	"log"

	"github.com/VooDooStack/FitStackAPI/models"
)

type Config struct {
	Port string
	Env  string
}

type Application struct {
	Config Config
	Logger *log.Logger
	Models models.Models
}
