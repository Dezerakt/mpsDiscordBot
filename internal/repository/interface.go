package repository

import (
	"context"
	"mpsDiscordBot/internal/domain"
	"mpsDiscordBot/vo"
)

type (
	IChannel interface {
		GetChannelInfo(ctx context.Context, channelId string) (*domain.Channel, error)
		StoreNewAction(ctx context.Context, channelId string, action vo.ActionFlag, switchType bool) error
	}
)
