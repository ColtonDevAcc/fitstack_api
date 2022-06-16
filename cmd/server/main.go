package main

import (
	"fmt"
	"os"

	"github.com/VooDooStack/FitStackAPI/internal/comment"
	"github.com/VooDooStack/FitStackAPI/internal/database"
	transportHTTP "github.com/VooDooStack/FitStackAPI/internal/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// contains app information
type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"AppName": app.Name,
		"Version": app.Version,
	}).Info("Starting FitStackAPI")

	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	h := handler.SetupRoutes()

	// set db to gin context with a middleware to all incoming request
	h.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(port)
	return nil
}

func main() {
	app := App{
		Name:    "FitStackAPIDev",
		Version: "0.0.1",
	}

	if err := app.Run(); err != nil {
		log.Error("error starting server error:", err)
		log.Fatal(err)
	}

}
