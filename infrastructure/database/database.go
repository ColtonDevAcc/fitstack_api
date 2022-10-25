package database

import (
	"context"
	"fmt"
	"os"

	"github.com/VooDooStack/FitStackAPI/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase(config config.Config) (*pgxpool.Pool, error) {
	fmt.Println("Setting up database...")

	db, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println("failed to setup database error:", err)
		return db, err
	}

	err = MigrateDB(db)
	if err != nil {
		fmt.Println(err)
	}

	return db, nil
}

//go.sum,Dockerfile,*.yaml,*.json,.gitignore,.gcloudignore,.env,*.md,Makefile
