package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"item/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewMongoDB)

// Data .
type Data struct {
	db  *mongo.Database
	log *log.Helper
}

func NewMongoDB(conf conf.Data) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongodb.GetUri()))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client.Database(conf.Mongodb.GetDatabase())
}

// NewData .
func NewData(database *mongo.Database, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "item-service/data"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	d := &Data{
		db:  database,
		log: log,
	}
	return d, func() {
		if err := d.db.Client().Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}, nil
}
