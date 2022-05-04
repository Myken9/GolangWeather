package storage

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

type Queryer interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
	Prepare(context.Context, string, string) (*pgconn.StatementDescription, error)
}

type Storage struct {
	db Queryer
}

func NewStorage(conn Queryer) *Storage {
	return &Storage{db: conn}
}

func (s *Storage) SaveUserMessage(msg tgbotapi.Message, answer string) error {
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return err
	}

	_, err = s.db.Exec(context.Background(),
		"INSERT INTO users"+
			" (chat_id, first_name, last_name, user_name, language_code)"+
			"VALUES ($1, $2, $3, $4, $5)"+
			"ON CONFLICT (chat_id) DO NOTHING;",
		msg.Chat.ID, msg.From.FirstName, msg.From.LastName, msg.From.UserName, msg.From.LanguageCode)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	_, err = s.db.Exec(context.Background(),
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
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return err
	}

	_, err = s.db.Exec(context.Background(),
		"INSERT INTO users"+
			" (chat_id, first_name, last_name, user_name, language_code)"+
			"VALUES ($1, $2, $3, $4, $5)"+
			"ON CONFLICT (chat_id) DO NOTHING;",
		msg.Chat.ID, msg.From.FirstName, msg.From.LastName, msg.From.UserName, msg.From.LanguageCode)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	_, err = s.db.Exec(context.Background(),
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
