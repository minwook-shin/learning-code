package main_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func dbInsert(db *sql.DB, id, count int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE entity SET views = views"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO entity (id, count) VALUES (?, ?)", id, count); err != nil {
		return
	}
	return
}

func TestDBUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE entity").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO entity").WithArgs(1, 10).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = dbInsert(db, 1, 10); err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
