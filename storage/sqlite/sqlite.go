package sqlite

import (
	"MyBot/Bot"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SqliteStorage struct {
	db *sql.DB
}

func (s *SqliteStorage) Save(update *Bot.Update) error {
	q := `INSERT INTO pages (url, user_name) VALUES (?, ?)`
	if _, err := s.db.Exec(q, update.Message.Text, update.Message.From.UserName); err != nil {
		return err
	}

	return nil

}
func (s *SqliteStorage) Remove(update *Bot.Update) error {
	//TODO implement me
	panic("implement me")
}
func (s *SqliteStorage) Random(update *Bot.Update) (string, error) {
	q := `SELECT url FROM pages WHERE user_name = ? ORDER BY RANDOM() LIMIT 1`

	var link string
	err := s.db.QueryRow(q, update.Message.From.UserName).Scan(&link)
	if err != nil {
		return "", err
	}

	return link, nil
}
func (s *SqliteStorage) IsExist(update *Bot.Update) bool {
	q := `SELECT COUNT(*) FROM pages WHERE url = ? AND user_name = ?`

	var link string
	if err := s.db.QueryRow(q, update.Message.Text, update.Message.From.UserName).Scan(&link); err != nil {
		return false
	}

	return true
}
func (s *SqliteStorage) Delete(update *Bot.Update) error {
	q := `DELETE FROM pages WHERE url = ? AND user_name = ?`

	_, err := s.db.Exec(q, update.Message.Text, update.Message.From.UserName)
	if err != nil {
		return err
	}

	return nil
}
func (s *SqliteStorage) Init() error {
	q := `CREATE TABLE IF NOT EXISTS pages (url TEXT, user_name TEXT)`

	if _, err := s.db.Exec(q); err != nil {
		return fmt.Errorf("can't init create table: %w", err)
	}

	return nil
}

func New(databasePath string) SqliteStorage {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatalf("can't open to database: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("can't connect to databse: %s", err.Error())
	}

	return SqliteStorage{db: db}
}
