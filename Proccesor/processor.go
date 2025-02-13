package Proccesor

import (
	"MyBot/Bot"
	"MyBot/consts"
	"MyBot/storage/sqlite"
	"database/sql"
	"net/url"
)

type Processor struct {
	TgBot   *Bot.TGBot
	Storage *sqlite.SqliteStorage
}

func (p *Processor) Process(update *Bot.Update) error {
	if p.isURL(update.Message.Text) {
		return p.Add(update)
	}
	switch update.Message.Text {
	case consts.MsgStart:
		return p.SendMessage(update, consts.RepHello)
	case consts.MsgHelp:
		return p.SendMessage(update, consts.RepHelp)
	case consts.MsgRnd:
		return p.Rnd(update)
	case consts.MsgRm:
		return p.SendMessage(update, consts.RepRm)
	default:
		return p.SendMessage(update, consts.RepUnknown)
	}
}
func (p *Processor) Add(update *Bot.Update) error {
	if err := p.Storage.Save(update); err != nil {
		return err
	}

	return p.SendMessage(update, "Добавлено!")
}
func (p *Processor) Rnd(update *Bot.Update) error {
	link, err := p.Storage.Random(update)
	if err == sql.ErrNoRows {
		link = "Нет ссылок"
	}
	if err != nil {
		return err
	}
	return p.SendMessage(update, link)
}
func (p *Processor) Rm(update *Bot.Update) error {
	return p.SendMessage(update, "Удалено")
}
func (p *Processor) SendMessage(update *Bot.Update, msg string) error {
	_, err := p.TgBot.SendMessage(update.Message.Chat.ID, msg)
	if err != nil {
		return err
	}
	return nil
}

func (p *Processor) isURL(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}

func New(TgBot *Bot.TGBot, db *sqlite.SqliteStorage) Processor {
	return Processor{
		TgBot:   TgBot,
		Storage: db,
	}
}
