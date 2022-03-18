package adapters

import (
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteTokenRepository struct {
	db *sqlx.DB
}

func NewSQLiteTokenRepository(db *sqlx.DB) *SQLiteTokenRepository  {
	return &SQLiteTokenRepository{
		db: db,
	}
}

func NewSQLiteConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", ":memory")
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to sqlite")
	}

	return	db, nil
}