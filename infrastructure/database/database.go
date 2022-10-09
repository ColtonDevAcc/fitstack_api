package database

import (
	"context"
	"fmt"
	"os"

	"github.com/VooDooStack/FitStackAPI/config"
	"github.com/jackc/pgx/v5"
)

func NewDatabase(config config.Config) (*pgx.Conn, error) {
	fmt.Println("Setting up database...")

	db, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println("failed to setup database error:", err)
		return db, err
	}

	return db, nil
}

//go.sum,Dockerfile,*.yaml,*.json,.gitignore,.gcloudignore,.env,*.md,Makefile
