package ft

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/cheekybits/genny/generic"
)

type Something generic.Type

func CreateSomething(ctx context.Context, client *firestore.Client, data Dict) (*firestore.DocumentRef, error) {
	df, _, err := client.Collection("somethings").Add(ctx, data)
	log.Println(df.ID)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return &firestore.DocumentRef{}, err
	}
	return df, nil
}

func GetSomething(ctx context.Context, client *firestore.Client, docID string) (*Something, error) {
	dsnap, err := client.Collection("somethings").Doc(docID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var something Something
	dsnap.DataTo(&something)
	log.Println("Document data: %#v\n", something)
	return &something, nil
}

func GetSomethingAsDict(ctx context.Context, client *firestore.Client, docID string) (Dict, error) {
	dsnap, err := client.Collection("somethings").Doc(docID).Get(ctx)
	if err != nil {
		return nil, err
	}

	m := dsnap.Data()
	log.Println("Document data: %#v\n", m)
	return m, nil
}
