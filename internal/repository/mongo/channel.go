package mongoRepo

import (
	"context"
	"errors"
	"log"
	"mpsDiscordBot/internal/domain"
	"mpsDiscordBot/internal/repository"
	mongoPkg "mpsDiscordBot/pkg/mongo"
	"mpsDiscordBot/vo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type channelRepository struct {
	*mongoPkg.Wrap
	channelCollection *mongo.Collection
}

func NewChannelRepository(mongo *mongoPkg.Wrap) repository.IChannel {
	return &channelRepository{
		Wrap:              mongo,
		channelCollection: mongo.EventCollection(mongoPkg.Channels),
	}
}

func (obj *channelRepository) GetChannelInfo(ctx context.Context, channelId string) (*domain.Channel, error) {
	filter := bson.M{"channelId": channelId}

	var channel domain.Channel
	findResult := obj.channelCollection.FindOne(ctx, filter)
	if err := findResult.Decode(&channel); err != nil {
		return nil, err
	}

	if errors.Is(findResult.Err(), mongo.ErrNoDocuments) {
		return nil, nil
	}

	return &channel, nil
}

func (obj *channelRepository) StoreNewAction(ctx context.Context, channel *domain.Channel, action vo.ActionFlag, switchType bool) (err error) {
	filter := bson.M{"channelId": channel.ChannelId}

	var upsert bson.M
	if switchType {
		upsert = bson.M{
			"$addToSet": bson.M{
				"serverId":    channel.ServerId,
				"actionFlags": action,
			},
		}
	} else {
		upsert = bson.M{
			"$pull": bson.M{
				"actionFlags": action,
			},
		}
	}

	option := options.UpdateOne().SetUpsert(true)

	updateResult, err := obj.channelCollection.UpdateOne(ctx, filter, upsert, option)
	if err != nil {
		return err
	}

	log.Printf("updateResult: %+v\n", updateResult)
	return nil
}
