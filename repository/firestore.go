package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitFirestore(ctx context.Context) (*firestore.Client, error) {
	sa := option.WithCredentialsFile("go-rest-a2f67-firebase-adminsdk-xc5t4-7cc2619ba7.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client, nil
}
