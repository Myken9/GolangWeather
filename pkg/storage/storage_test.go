package storage

import (
	"GolangWeather/pkg/telegram"
	"context"

	"github.com/pashagolub/pgxmock"
	"testing"
)

func TestStorage_SaveUserLocation(t *testing.T) {

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close(context.Background())

	sdf := NewStorage(mock)

	mock.ExpectExec("INSERT INTO users").WithArgs(2, "FirstName", "FirstName", "FirstName", "FirstName").WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mock.ExpectExec("INSERT INTO message").WithArgs(2, "Text", 0.2, 0.2, 10, "df", 11).WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mock.ExpectCommit()

	if err = sdf.SaveUserMessage(telegram.Message{}, "df"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
