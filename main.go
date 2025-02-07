package main

import (
	"MyBot/Bot"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	tgBot := Bot.New(os.Getenv("TOKEN"))
	if err := tgBot.GetMe(); err != nil {
		log.Fatalln(err)
	}
	offset := 0
	for {
		updates, err := tgBot.GetUpdates(offset)
		if err != nil {
			log.Printf("No Updates")
			time.Sleep(5 * time.Second)
			continue
		}
		offset = Bot.PrintUpdates(updates)
		for _, update := range updates {
			tgBot.SendMessage(update.Message.Chat.ID, "Hello, Nigga!")
		}
		time.Sleep(5 * time.Second)
	}
}
