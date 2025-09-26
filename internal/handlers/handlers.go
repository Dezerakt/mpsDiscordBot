package handlers

import (
	"context"
	"fmt"
	"log"
	"mpsDiscordBot/internal/actions"
	v1 "mpsDiscordBot/internal/handlers/v1"
	"mpsDiscordBot/internal/repository"
	"mpsDiscordBot/vo"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	channelRepo repository.IChannel
}

func NewHandler(channelRepo repository.IChannel) Handler {
	return Handler{channelRepo: channelRepo}
}

func (obj *Handler) RegisterNewActions(s *discordgo.Session, m *discordgo.MessageCreate) {
	ctx := context.Background()

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content[0] == '~' {
		return
	}

	isUserAdmin, err := actions.IsAuthorAdmin(s, m)
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("Error checking user permissions"))
		return
	}

	if !isUserAdmin {
		s.ChannelMessageSend(m.ChannelID, ")")
		return
	}

	statement := strings.Split(m.Content, " ")
	if len(statement) == 0 {
		return
	}

	switch statement[0] {
	case "~channel":
		err = v1.HandleChannel(ctx, obj.channelRepo, s, m)
	}

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
	}
}

func (obj *Handler) HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	ctx := context.Background()

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content[0] == '~' {
		return
	}

	channelData, err := obj.channelRepo.GetChannelInfo(ctx, m.ChannelID)
	if err != nil {
		log.Println(err)
		return
	}

	if channelData == nil {
		return
	}

	for _, chActionFlag := range channelData.ActionFlags {
		var err error

		switch chActionFlag {
		case vo.ThreadCreate:
			err = actions.CreateThread(s, m)

		}
		if err != nil {
			log.Println(err)
			return
		}
	}
}
