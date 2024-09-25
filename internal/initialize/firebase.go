package initialize

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/hainguyen27798/open-notification/global"
	"google.golang.org/api/option"
)

func SetupFirestore() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	global.FirestoreClient = client
}
