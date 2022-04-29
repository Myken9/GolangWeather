package storage

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4"
	"time"
)

type Storage struct {
	*pgx.Conn
}

func NewStorage(db *pgx.Conn) *Storage {
	return &Storage{db}
}

func (s *Storage) SaveUser(msg tgbotapi.Message) {
	s.Exec(context.Background(),
		"INSERT INTO users"+
			" (chat_id, FirstName, LastName, UserName, LanguageCode)"+
			"VALUES ($1, $2, $3, $4, $5);",
		msg.Chat.ID, msg.From.FirstName, msg.From.LastName, msg.From.UserName, msg.From.LanguageCode)

}

func (s *Storage) SaveMessage(msg tgbotapi.Message, answer string) {
	s.Exec(context.Background(),
		"INSERT INTO message"+
			" (chat_id, msg_text, longitude, latitude, receive_at, response_text, response_at)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7);",
		msg.Chat.ID, msg.Text, msg.Location.Longitude, msg.Location.Latitude, msg.Date, answer, time.Now().Unix())

}
