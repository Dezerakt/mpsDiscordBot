package domain

import "mpsDiscordBot/vo"

type Channel struct {
	ServerId    string          `json:"serverId" bson:"serverId"`
	ChannelId   string          `json:"channelId" bson:"channelId"`
	ActionFlags []vo.ActionFlag `json:"actionFlags" bson:"actionFlags"`
}
