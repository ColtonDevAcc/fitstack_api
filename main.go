package main

import (
	"fmt"
	"log"
	"os"

	"github.com/VooDooStack/FitStackAPI/api"
	"github.com/VooDooStack/FitStackAPI/config"
	"github.com/VooDooStack/FitStackAPI/infrastructure/database"
	"github.com/joho/godotenv"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func run() error {
	godotenv.Load(".env")
	cfg := config.Config{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("ENV"),
	}


	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := database.NewDatabase(cfg)
	if err != nil {
		logger.Fatal(err)
		return err
	}

	//create router
	http := api.NewRouter(db)

	// flag.IntVar(&cfg.Port, "port", 4000, "Server port to listen on")
	err = http.Run(fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		logger.Fatal(err)
		return err
	}
	//print where http server is listening. Address and port
	logger.Println("Listening on port", cfg.Port)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("error starting server error:", err)
		panic(err)
	}
}

//! this is to generate the firebase auth token
//! token, err := firebaseAuth.CustomToken(context.Background(), "firebase_UID")
