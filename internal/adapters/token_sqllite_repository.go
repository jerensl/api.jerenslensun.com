package adapters

import (
	"database/sql"
	"os"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", os.Getenv("SQLITE_DB"))
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to sqlite")
	}

	schema := `CREATE TABLE IF NOT EXISTS users (
		token text NOT NULL
	);
			CREATE UNIQUE INDEX IF NOT EXISTS idx_token_user
		ON users (token);
	`

	_, err = db.Exec(schema)
	if err != nil {
		return nil, errors.Wrap(err, "cannot execute schema to sqlite")
	}

	return	db, nil
}

type SQLiteTokenRepository struct {
	db *sqlx.DB
}

func NewSQLiteTokenRepository(db *sqlx.DB) *SQLiteTokenRepository  {
	if db == nil {
		panic("missing database")
	} 
	
	return &SQLiteTokenRepository{
		db: db,
	}
}

func (s SQLiteTokenRepository) UpdatedToken(token string) error {
	return s.updatedToken(token)
}

func (s SQLiteTokenRepository) updatedToken(token string) error {
	insert := `INSERT INTO users (token) VALUES (?)`

	_, err := s.db.Exec(insert, token)
	if err != nil {
		return errors.Wrap(err, "Unable to insert token to database")
	}
	
	return nil
}
func (s SQLiteTokenRepository) GetToken(value string) (hasValue bool, err error) {
	return s.getToken(value)
}

func (s SQLiteTokenRepository) getToken(value string) (hasValue bool, err error) {
	var values int
	
	row := s.db.QueryRow("SELECT 1 FROM users WHERE token = (?)",value)
	err = row.Scan(&values)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, errors.Wrap(err, "unable to get token from db")
	}
	return true, nil
}

func (s SQLiteTokenRepository) GetAllToken() (subscriber []string, err error) {
	return s.getAllToken()
}

func (s SQLiteTokenRepository) getAllToken() (subscriber []string, err error) {
	err = s.db.Select(&subscriber,"SELECT token FROM users")
	if err != nil {
		return nil,  errors.Wrap(err, "unable to get all token from db")
	}	

	return subscriber, nil
}

func (s SQLiteTokenRepository) DeleteToken(token string) error {
	return s.deleteToken(token)
}

func (s SQLiteTokenRepository) deleteToken(token string) error {
	insert := "DELETE FROM users WHERE token = (?) RETURNING token"

	err := s.db.QueryRow(insert, token).Scan(&token)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.Wrap(err, "Cannot find token from database")
	} else if err != nil {
		return errors.Wrap(err, "Unable to delete token from database")
	}

	return nil
}