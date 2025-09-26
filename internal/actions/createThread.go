package actions

import "github.com/bwmarrin/discordgo"

func CreateThread(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.MessageThreadStartComplex(m.ChannelID, m.ID, &discordgo.ThreadStart{
		Name:                m.Author.Username,
		AutoArchiveDuration: 10,
		Invitable:           false,
		RateLimitPerUser:    0,
	})
	if err != nil {
		return err
	}

	return nil
}
