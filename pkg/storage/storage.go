package storage

import (
	"GolangWeather/pkg/telegram"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
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

func (s *Storage) SaveUserMessage(msg telegram.Message, answer string) error {
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return err
	}

	_, err = s.db.Exec(context.Background(),
		"INSERT INTO users"+
			" (chat_id, first_name, last_name, user_name, language_code)"+
			"VALUES ($1, $2, $3, $4, $5)"+
			"ON CONFLICT (chat_id) DO NOTHING;",
		msg.ChatId, msg.From.FirstName, msg.From.LastName, msg.From.UserName, msg.From.LanguageCode)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	_, err = s.db.Exec(context.Background(),
		"INSERT INTO message"+
			" (chat_id, msg_text, receive_at, response_text, response_at)"+
			"VALUES ($1, $2, $3, $4, $5);",
		msg.ChatId, msg.Text, msg.ReceiveAt, answer, msg.ResponseAt)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	err = tx.Commit(context.Background())
	return err
}

func (s *Storage) SaveUserLocation(msg telegram.Message, answer string) error {
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return err
	}

	_, err = s.db.Exec(context.Background(),
		"INSERT INTO users"+
			" (chat_id, first_name, last_name, user_name, language_code)"+
			"VALUES ($1, $2, $3, $4, $5)"+
			"ON CONFLICT (chat_id) DO NOTHING;",
		msg.ChatId, msg.From.FirstName, msg.From.LastName, msg.From.UserName, msg.From.LanguageCode)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	_, err = s.db.Exec(context.Background(),
		"INSERT INTO message"+
			" (chat_id, msg_text, longitude, latitude, receive_at, response_text, response_at)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7);",
		msg.ChatId, msg.Text, msg.Location.Longitude, msg.Location.Latitude, msg.ReceiveAt, answer, msg.ResponseAt)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	err = tx.Commit(context.Background())
	return err
}
