package mongodb

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db   *mongo.Database
	once sync.Once
)

func Connect() *mongo.Database {

	once.Do(func() {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
		if err != nil {
			panic(err)
		}
		if err = client.Ping(context.TODO(), nil); err != nil {
			panic(err)
		}
		db = client.Database("MONGODB_DATABASE")
	})
	return db
}
