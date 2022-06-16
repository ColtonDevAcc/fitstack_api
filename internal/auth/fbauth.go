package auth

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

var (
	firebaseConfigFile = os.Getenv("FIREBASE_CONFIG_FILE")
)

func InitAuth() (*auth.Client, error) {
	opt := option.WithCredentialsFile(firebaseConfigFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing firebase auth (create firebase app)")
	}

	client, errAuth := app.Auth(context.Background())
	if errAuth != nil {
		return nil, errors.Wrap(errAuth, "error initializing firebase auth (creating client)")
	}

	return client, nil
}
