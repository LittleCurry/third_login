package third_login

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
)

var client = &auth.Client{}

func InitFirebase(path string) error {
	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(path))
	if err != nil {
		log.Printf("init firebase: %v\n\n", err)
		return err
	}
	tempClient, err := app.Auth(context.Background())
	if err != nil {
		log.Printf("init firebase: %v\n\n", err)
		return err
	}
	client = tempClient
	return nil
}

func CheckIdToken(idToken string) (*auth.Token, error) {
	return client.VerifyIDToken(context.Background(), idToken)
}
