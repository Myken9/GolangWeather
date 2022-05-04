package storage

import (
	"GolangWeather/pkg/telegram"
	"context"
	"fmt"
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

func (s *Storage) SaveUserMessage(msg telegram.Message, answer string) (e error) {
	tx, e := s.db.Begin(context.Background())
	defer func() {
		if e != nil {
			e = tx.Rollback(context.Background())
		} else {
			e = tx.Commit(context.Background())
		}
	}()

	_, e = s.db.Exec(context.Background(),
		"INSERT INTO users"+
			" (chat_id, first_name, last_name, user_name, language_code)"+
			"VALUES ($1, $2, $3, $4, $5)"+
			"ON CONFLICT (chat_id) DO NOTHING;",
		msg.ChatId, msg.From.FirstName, msg.From.LastName, msg.From.UserName, msg.From.LanguageCode)

	var longitude, latitude string
	if msg.Location != nil {
		longitude = fmt.Sprintf("%v", msg.Location.Longitude)
		latitude = fmt.Sprintf("%v", msg.Location.Latitude)
	}
	_, e = s.db.Exec(context.Background(),
		"INSERT INTO message"+
			" (chat_id, msg_text, longitude, latitude, receive_at, response_text, response_at)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7);",
		msg.ChatId, msg.Text, longitude, latitude, msg.ReceiveAt, answer, msg.ResponseAt)
	return e
}
