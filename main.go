package main

import (
	"flag"
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
	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development|production")
	flag.StringVar(&cfg.DbConnection, "dsn", "postgres://tcs@localhost/go_movies?sslmode=disable", "Postgres connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//! configure firebase
	//! Request must have the following
	//! Authorization: Bearer {{user_firebase_token}}
	firebaseAuth, err := config.SetupFirebase()
	if err != nil {
		logger.Fatal(err)

		return err
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		logger.Fatal(err)
		return err
	}

	//create router
	http := api.NewRouter(db, firebaseAuth)

	err = http.Run(fmt.Sprintf(":%d", cfg.Port))
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
