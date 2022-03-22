package adapters

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteConnection(file string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", file)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to sqlite")
	}

	schema := `CREATE TABLE IF NOT EXISTS token (
		token text NOT NULL PRIMARY KEY
	);`

	_, err = db.Exec(schema)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to sqlite")
	}

	return	db, nil
}

type SQLiteTokenRepository struct {
	db *sqlx.DB
}

func NewSQLiteTokenRepository(db *sqlx.DB) *SQLiteTokenRepository  {

	return &SQLiteTokenRepository{
		db: db,
	}
}

func (s SQLiteTokenRepository) UpdatedToken(ctx context.Context, token string) error {
	return s.updatedToken(ctx, token)
}

func (s SQLiteTokenRepository) updatedToken(ctx context.Context, token string) error {
	insert := `INSERT INTO token (token) VALUES (?)`

	_, err := s.db.Exec(insert, token)
	if err != nil {
		return errors.Wrap(err, "Unable to insert token to database")
	}
	
	return nil
}

func (s SQLiteTokenRepository) GetToken(ctx context.Context, value string) (hasValue bool, err error) {
	var values int
	
	row := s.db.QueryRow("SELECT 1 FROM token WHERE token = (?)",value)
	err = row.Scan(&values)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, errors.Wrap(err, "unable to get token from db")
	}
	return true, nil
}

func (s SQLiteTokenRepository) GetAllToken(ctx context.Context) (subscriber []string, err error) {
	err = s.db.Select(&subscriber,"SELECT token FROM token")
	if err != nil {
		return nil, err
	}	

	return subscriber, nil
}