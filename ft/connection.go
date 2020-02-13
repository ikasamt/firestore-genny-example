package ft

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

const ProjectID = `talcjp-tokyo`

// GetConnection is get firestore connection
func GetConnection(ctx context.Context) (*firestore.Client, error) {
	conf := &firebase.Config{ProjectID: ProjectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Println(err)
		return &firestore.Client{}, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println(err)
		return &firestore.Client{}, err
	}
	return client, nil
}

type Dict map[string]interface{}
