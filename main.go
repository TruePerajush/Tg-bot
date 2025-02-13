package main

import (
	"MyBot/Bot"
	"MyBot/Proccesor"
	"MyBot/storage/sqlite"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	db := sqlite.New(os.Getenv("DATABASE_PATH"))
	if err := db.Init(); err != nil {
		log.Fatalf("Не подключилась база данных: %s", err.Error())
	}
	tgBot := Bot.New(os.Getenv("TOKEN"))
	if err := tgBot.GetMe(); err != nil {
		log.Fatalln(err)
	}
	processor := Proccesor.New(&tgBot, &db)
	offset := 0
	for {
		updates, err := tgBot.GetUpdates(offset)
		if err != nil || len(updates) == 0 {
			//log.Printf("Нет обновлений")
			time.Sleep(1 * time.Second)
			continue
		}
		Bot.PrintUpdates(updates)
		for _, update := range updates {
			if err := processor.Process(&update); err != nil {
				log.Printf("Ошибка: %v", err.Error())
			}
		}
		offset = updates[len(updates)-1].UpdateID + 1
		time.Sleep(1 * time.Second)
	}
}
