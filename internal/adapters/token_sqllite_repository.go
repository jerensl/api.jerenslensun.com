package adapters

import (
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

func (s SQLiteTokenRepository) UpdatedToken(token string) error {
	return s.updatedToken(token)
}

func (s SQLiteTokenRepository) updatedToken(token string) error {
	insert := `INSERT INTO token (token) VALUES (?)`

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
	
	row := s.db.QueryRow("SELECT 1 FROM token WHERE token = (?)",value)
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
	err = s.db.Select(&subscriber,"SELECT token FROM token")
	if err != nil {
		return nil,  errors.Wrap(err, "unable to get all token from db")
	}	

	return subscriber, nil
}

func (s SQLiteTokenRepository) DeleteToken(token string) error {
	return s.deleteToken(token)
}

func (s SQLiteTokenRepository) deleteToken(token string) error {
	insert := `DELETE FROM token WHERE token = (?)`

	_, err := s.db.Exec(insert, token)
	if err != nil {
		return errors.Wrap(err, "Unable to insert token to database")
	}
	return nil
}