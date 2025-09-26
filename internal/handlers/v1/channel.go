package v1

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mpsDiscordBot/internal/repository"
	"mpsDiscordBot/vo"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleChannel(ctx context.Context, channelRepo repository.IChannel, s *discordgo.Session, m *discordgo.MessageCreate) error {
	statement := strings.Split(m.Content, " ")

	if len(statement) < 1 {
		return errors.New("you should mention a channel")
	}

	rawChannelId := statement[1]
	if rawChannelId[0] != '<' && rawChannelId[1] != '#' && rawChannelId[len(rawChannelId)-1] != '>' {
		return errors.New("invalid channel id")
	}
	formattedChannelId := strings.ReplaceAll(rawChannelId, "#", "")
	formattedChannelId = strings.ReplaceAll(formattedChannelId, ">", "")
	formattedChannelId = strings.ReplaceAll(formattedChannelId, "<", "")

	if len(statement) < 2 {
		return errors.New("you should put an action")
	}
	action := statement[2]

	if len(statement) < 3 {
		return errors.New("you should specify switch type on/off")
	}
	rawSwitchType := statement[3]

	var switchType bool
	switch rawSwitchType {
	case "on":
		switchType = true
	case "off":
		switchType = false
	default:
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("unknown switch type: %s", rawSwitchType))
		return errors.New("you should specify switch type on/off")
	}

	var err error
	switch action {
	case "threaded":
		log.Println("threaded")
		err = channelRepo.StoreNewAction(ctx, formattedChannelId, vo.ThreadCreate, switchType)
	}
	if err != nil {
		return err
	}

	return nil
}
