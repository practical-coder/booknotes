package db

import (
	"github.com/practical-coder/booknotes/zlg"
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var once sync.Once
var ctx = context.Background()

func init() {
	once.Do(func() {
		uri := os.Getenv("MONGODB_URI")
		zlg.Logger.Info().Str("MONGODB_URI", uri).Send()
		clientOpts := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(ctx, clientOpts)
		if err != nil {
			zlg.Logger.Error().Err(err).Send()
			os.Exit(1)
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			zlg.Logger.Error().Err(err).Send()
			os.Exit(2)
		}

		Client = client
	})
}
