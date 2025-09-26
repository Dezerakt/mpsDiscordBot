package mongoPkg

import (
	"context"
	"mpsDiscordBot/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Wrap struct {
	*mongo.Client
}

func NewConnection(config *config.Config) (*Wrap, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(config.Mongo.ConnectionString))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &Wrap{
		Client: client,
	}, nil
}

func (obj *Wrap) EventCollection(collection Collection) *mongo.Collection {
	return obj.Database(Events.Str()).Collection(collection.Str())
}
