package storage

import "MyBot/Bot"

type Storage interface {
	Save(update *Bot.Update) error
	Remove(update *Bot.Update) error
	Random(update *Bot.Update) string
	IsExist(update *Bot.Update) bool
}
