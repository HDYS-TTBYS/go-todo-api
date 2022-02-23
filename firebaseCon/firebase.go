package firebasecon

import (
	"context"
	"encoding/base64"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/HDYS-TTBYS/go-todo-api/config"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App

func Init() {
	decFirebaseCredential, err := base64.StdEncoding.DecodeString(config.GetConfig().FIREBASE_CREDENTIAL)
	if err != nil {
		log.Fatal(err)
	}
	opt := option.WithCredentialsJSON(decFirebaseCredential)
	firebaseApp, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing furebase app: %v", err)
	}
}

func GetFirebaseApp() *firebase.App {
	return firebaseApp
}
