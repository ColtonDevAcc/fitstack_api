// config/firebase.go
package config

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"google.golang.org/api/option"
)

func SetupFirebase() (*auth.Client, *storage.Client, error) {
	var opt option.ClientOption
	var app *firebase.App
	var err error
	var storage *storage.Client

	config := &firebase.Config{
		StorageBucket: os.Getenv("BUCKET"),
	}

	file := os.Getenv("FIREBASE_CREDENTIALS_FILE")
	if len(file) == 0 {
		fmt.Println("init firebase using default credentials file")

		//Firebase admin SDK initialization
		app, err = firebase.NewApp(context.Background(), nil)
		if err != nil {
			panic(fmt.Sprintf("error initializing app: %v", err))
		}
	} else {
		fmt.Println("init firebase using specified credentials file")

		opt = option.WithCredentialsFile(file)

		//Firebase admin SDK initialization
		app, err = firebase.NewApp(context.Background(), config, opt)
		if err != nil {
			panic(fmt.Sprintf("error initializing app: %v", err))
		}

		//setup storage bucket
		storage, err = app.Storage(context.Background())
		if err != nil {
			panic(fmt.Sprintf("error initializing storage bucket: %v", err))
		}

		if storage == nil {
			fmt.Printf("bucket is null ==============")
		}
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(fmt.Sprintf("error initializing app: %v", err))
	}

	return auth, storage, nil
}
