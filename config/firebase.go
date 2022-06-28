// config/firebase.go
package config

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func SetupFirebase() (*auth.Client, error) {
	var opt option.ClientOption

	file := os.Getenv("FIREBASE_CREDENTIALS_FILE")
	if len(file) == 0 {
		opt = nil
	} else {
		opt = option.WithCredentialsFile(file)
	}

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("failed to initialize firebase admin sdk")
		return nil, err
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		fmt.Println("failed to generate firebase auth")
		return nil, err
	}

	return auth, nil
}