package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type Adapter struct {
	App *firebase.App
	DB  *firestore.Client
}

func NewFireBaseAdapter(c Config) *Adapter {
	var opt option.ClientOption
	if c.CertPath != "" {
		opt = option.WithCredentialsFile(c.CertPath)
	} else {
		opt = option.WithCredentialsJSON([]byte(c.CertJSON))
	}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	db, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	return &Adapter{
		App: app,
		DB:  db,
	}
}
