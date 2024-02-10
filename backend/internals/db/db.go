package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbTimeout = 5 * time.Second

type Database struct {
	Client *mongo.Client
}

func NewDb(uri string) (*Database, error) {
	client, err := newClient(uri)
	if err != nil {
		return nil, err
	}
	db := new(Database)
	db.Client = client

	for i := 1; i <= 5; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
		defer cancel()
		if err = db.Client.Ping(ctx, nil); err == nil {
			return db, nil
		} else {
			log.Printf("database ping check %v failed, waiting 2 seconds before trying again\n", i)
		}
		time.Sleep(2 * time.Second)
	}

	return nil, err
}

func newClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}