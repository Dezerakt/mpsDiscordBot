package actions

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func IsAuthorAdmin(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	perms, err := s.State.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		perms, err = s.UserChannelPermissions(m.Author.ID, m.ChannelID)
		if err != nil {
			log.Println("ошибка получения прав:", err)
			return false, nil
		}
	}

	if perms&discordgo.PermissionAdministrator == 0 {
		return false, nil
	}

	return true, nil
}
