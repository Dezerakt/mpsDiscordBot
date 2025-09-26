package main

import (
	"fmt"
	"log"
	"mpsDiscordBot/config"
	"mpsDiscordBot/internal/handlers"
	mongoRepo "mpsDiscordBot/internal/repository/mongo"
	mongoPkg "mpsDiscordBot/pkg/mongo"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// config
	configObj, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	mongoConnection, err := mongoPkg.NewConnection(configObj)
	if err != nil {
		log.Fatal(err)
	}

	disSession, err := discordgo.New("Bot " + configObj.App.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	err = disSession.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer disSession.Close()

	// repos
	channelRepo := mongoRepo.NewChannelRepository(mongoConnection)

	// handlers
	handler := handlers.NewHandler(channelRepo)
	disSession.AddHandler(handler.RegisterNewActions)
	disSession.AddHandler(handler.HandleMessage)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Bot disabled")
}
