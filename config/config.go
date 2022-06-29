package config

import (
	"log"

	"github.com/VooDooStack/FitStackAPI/models"
)

type Config struct {
	Port         int
	Env          string
	DbConnection string
}

type Application struct {
	Config Config
	Logger *log.Logger
	Models models.Models
}
