package domain

import "mpsDiscordBot/vo"

type Channel struct {
	ChannelId   string          `json:"channelId" bson:"channelId"`
	ActionFlags []vo.ActionFlag `json:"actionFlags" bson:"actionFlags"`
}
