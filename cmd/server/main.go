package main

import (
	"fmt"
	"net/http"

	"github.com/VooDooStack/FitStackAPI/internal/comment"
	"github.com/VooDooStack/FitStackAPI/internal/database"
	transportHTTP "github.com/VooDooStack/FitStackAPI/internal/transport/http"
)

// contains things like db connection, routes, etc.
type App struct {
}

func (a *App) Run() error {
	fmt.Println("Setting up app...")

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
		fmt.Println("failed to setup server error:", err)
		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("error starting server error:", err)
		panic(err)
	}

}
