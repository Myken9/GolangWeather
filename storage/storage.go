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

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{conn}
}

func (s *Storage) SaveUserMessage(msg tgbotapi.Message, answer string) error {
	tx, err := s.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}

	_, err = s.Exec(context.Background(),
		"INSERT INTO users"+
			" (chat_id, first_name, last_name, user_name, language_code)"+
			"VALUES ($1, $2, $3, $4, $5)"+
			"ON CONFLICT (chat_id) DO NOTHING;",
		msg.Chat.ID, msg.From.FirstName, msg.From.LastName, msg.From.UserName, msg.From.LanguageCode)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	_, err = s.Exec(context.Background(),
		"INSERT INTO message"+
			" (chat_id, msg_text, receive_at, response_text, response_at)"+
			"VALUES ($1, $2, $3, $4, $5);",
		msg.Chat.ID, msg.Text, msg.Date, answer, time.Now().Unix())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	err = tx.Commit(context.Background())
	return err
}

func (s *Storage) SaveUserLocation(msg tgbotapi.Message, answer string) error {
	tx, err := s.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}

	_, err = s.Exec(context.Background(),
		"INSERT INTO users"+
			" (chat_id, first_name, last_name, user_name, language_code)"+
			"VALUES ($1, $2, $3, $4, $5)"+
			"ON CONFLICT (chat_id) DO NOTHING;",
		msg.Chat.ID, msg.From.FirstName, msg.From.LastName, msg.From.UserName, msg.From.LanguageCode)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	_, err = s.Exec(context.Background(),
		"INSERT INTO message"+
			" (chat_id, msg_text, longitude, latitude, receive_at, response_text, response_at)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7);",
		msg.Chat.ID, msg.Text, msg.Location.Longitude, msg.Location.Latitude, msg.Date, answer, time.Now().Unix())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	err = tx.Commit(context.Background())
	return err
}
