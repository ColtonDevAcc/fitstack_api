package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/VooDooStack/FitStackAPI/config"
	"github.com/VooDooStack/FitStackAPI/infrastructure/database"
	"github.com/gin-gonic/gin"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func run() error {
	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development|production")
	flag.StringVar(&cfg.DbConnection, "dsn", "postgres://tcs@localhost/go_movies?sslmode=disable", "Postgres connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//! configure firebase
	//! Request must have the following
	//! Authorization: Bearer {{user_firebase_token}}
	firebaseAuth := config.SetupFirebase()

	//create router
	http := NewRouter()

	db, err := database.NewDatabase(cfg)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	defer db.Close()

	//! set db & firebase auth to gin context with a middleware to all incoming request
	http.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", firebaseAuth)
	})

	http.Run(fmt.Sprintf(":%d", cfg.Port))

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
