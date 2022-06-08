package main

import (
	"net/http"

	"github.com/VooDooStack/FitStackAPI/internal/comment"
	"github.com/VooDooStack/FitStackAPI/internal/database"
	transportHTTP "github.com/VooDooStack/FitStackAPI/internal/transport/http"
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
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

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
