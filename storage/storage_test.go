package storage

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pashagolub/pgxmock"
	"testing"
)

func TestStorage_SaveUserLocation(t *testing.T) {

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close(context.Background())

	mock.ExpectExec("UPDATE products").WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mock.ExpectCommit()

	if err = SaveUserMessage(tgbotapi.Message, "df"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}
