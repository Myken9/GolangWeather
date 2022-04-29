package storage

import (
	"context"
	"fmt"
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

func (s *Storage) SaveUserMessage(msg tgbotapi.Message, answer string) {
	s.saveUser(msg)
	s.saveMessage(msg, answer)
}

func (s *Storage) SaveUserLocation(msg tgbotapi.Message, answer string) {
	s.saveUser(msg)
	s.saveLocation(msg, answer)
}

func (s *Storage) saveUser(msg tgbotapi.Message) {
	_, err := s.Exec(context.Background(),
		"INSERT INTO users"+
			" (chat_id, first_name, last_name, user_name, language_code)"+
			"VALUES ($1, $2, $3, $4, $5)"+
			"ON CONFLICT (chat_id) DO NOTHING;",
		msg.Chat.ID, msg.From.FirstName, msg.From.LastName, msg.From.UserName, msg.From.LanguageCode)
	if err != nil {
		fmt.Println(err)
	}

}

func (s *Storage) saveMessage(msg tgbotapi.Message, answer string) {
	_, err := s.Exec(context.Background(),
		"INSERT INTO message"+
			" (chat_id, msg_text, receive_at, response_text, response_at)"+
			"VALUES ($1, $2, $3, $4, $5);",
		msg.Chat.ID, msg.Text, msg.Date, answer, time.Now().Unix())
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Storage) saveLocation(msg tgbotapi.Message, answer string) {
	_, err := s.Exec(context.Background(),
		"INSERT INTO message"+
			" (chat_id, msg_text, longitude, latitude, receive_at, response_text, response_at)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7);",
		msg.Chat.ID, msg.Text, msg.Location.Longitude, msg.Location.Latitude, msg.Date, answer, time.Now().Unix())
	if err != nil {
		fmt.Println(err)
	}

}
