package adapters

import (
	"database/sql"
	"time"

	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteToken struct {
	TokenID		string		`db:"token_id"`
	IsActive	int			`db:"is_active"`
	UpdatedAt	int64		`db:"updated_at"`
}

type sqliteTokenStat struct {
	TotalTokens			int			`db:"total_tokens"`
	TotalActiveTokens	int			`db:"total_active_tokens"`
	TotalInactiveTokens	int			`db:"total_inactive_tokens"`
}


func NewSQLiteConnection(path string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to sqlite")
	}

	schema := `CREATE TABLE IF NOT EXISTS tokens (
		token_id VARCHAR(250) NOT NULL,
		is_active INTEGER,
		updated_at INTEGER
	);
			CREATE UNIQUE INDEX IF NOT EXISTS idx_token_user
		ON tokens (token_id);
	`

	db.MustExec(schema)

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

func (s SQLiteTokenRepository) UpdatedToken(tokenID string, updatedAt int64) (*notification.Token, error) {
	return s.updatedToken(tokenID, updatedAt)
}

func (s SQLiteTokenRepository) updatedToken(tokenID string, updatedAt int64) (*notification.Token, error) {
	token, isExist, err := s.getToken(tokenID)
	if err != nil {
		return nil, err
	}
	

	if !isExist {
		token, err := s.insertedToken(tokenID, updatedAt)
		if err != nil {
			return nil, err
		}		
		return token, nil
	}

	reverseIsActive := 1

	if token.IsActive() {
		reverseIsActive = 0
	}

	query := `UPDATE tokens SET is_active=? , updated_at=? WHERE token_id=? RETURNING *;`
	var tokenFromDB sqliteToken
	err = s.db.QueryRowx(query,reverseIsActive, updatedAt, tokenID).StructScan(&tokenFromDB)
	if err != nil {
		return nil,errors.Wrap(err, "Unable to insert token to database")
	}

	newToken, err := notification.UnmarshalTokenFromDatabase(tokenFromDB.TokenID, tokenFromDB.IsActive == 1, tokenFromDB.UpdatedAt)
	if err != nil {
		return nil,errors.Wrap(err, "Unable to parse token from database")
	}
	return newToken,nil
}

func (s SQLiteTokenRepository) InsertedToken(tokenID string, updatedAt int64) (*notification.Token, error) {
	return s.insertedToken(tokenID, updatedAt)
}

func (s SQLiteTokenRepository) insertedToken(tokenID string, updatedAt int64) (*notification.Token, error) {
	query := `INSERT INTO tokens (token_id, is_active, updated_at) VALUES (?, ?, ?) RETURNING *;`
	var tokenFromDB sqliteToken
	err := s.db.QueryRowx(query, tokenID, 1, updatedAt).StructScan(&tokenFromDB)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to insert token to database")
	}

	token, err := notification.UnmarshalTokenFromDatabase(tokenFromDB.TokenID, tokenFromDB.IsActive == 1, tokenFromDB.UpdatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to insert token to database")
	}

	return token, nil
}

func (s SQLiteTokenRepository) GetToken(value string) (*notification.Token, bool, error) {
	return s.getToken(value)
}

func (s SQLiteTokenRepository) getToken(tokenID string) (*notification.Token, bool, error)  {
	var tokenFromDB sqliteToken
	
	err := s.db.QueryRowx("SELECT * FROM tokens WHERE token_id = (?)", tokenID).StructScan(&tokenFromDB)
	if errors.Is(err, sql.ErrNoRows) {
		token, err := notification.NewToken(tokenID, false, time.Now().Unix())
		if err != nil {
			return nil, true, errors.Wrap(err, "Unable to parse token")
		}
		return token, false, nil
	} else if err != nil {
		return nil, false,errors.Wrap(err, "unable to get token from db")
	}

	token, err := notification.UnmarshalTokenFromDatabase(tokenFromDB.TokenID, tokenFromDB.IsActive == 1, tokenFromDB.UpdatedAt)
	if err != nil {
		return nil, true, errors.Wrap(err, "Unable to get token from database")
	}

	return token, true, nil
}

func (s SQLiteTokenRepository) GetAllToken() (subscriber []string, err error) {
	return s.getAllToken()
}

func (s SQLiteTokenRepository) getAllToken() (subscriber []string, err error) {
	err = s.db.Select(&subscriber,"SELECT token_id FROM tokens WHERE is_active = true")
	if err != nil {
		return nil,  errors.Wrap(err, "unable to get all token from db")
	}	

	return subscriber, nil
}

func (s SQLiteTokenRepository) GetStatisticToken() (*notification.Stats, error) {
	return s.getStatisticToken()
}

func (s SQLiteTokenRepository) getStatisticToken() (*notification.Stats, error) {
	var TokenStat sqliteTokenStat

	err := s.db.QueryRowx("SELECT count(*) AS total_tokens, SUM(is_active = 1) AS total_active_tokens, SUM(is_active = 0) AS total_inactive_tokens FROM tokens").StructScan(&TokenStat)

	if errors.Is(err, sql.ErrNoRows) {
		return &notification.Stats{}, errors.Wrap(err, "Unable to parse token")
	} else if err != nil {
		return &notification.Stats{}, errors.Wrap(err, "unable to get token from db")
	}

	stats, err := notification.UnmarshalStatsFromDatabase(TokenStat.TotalTokens, TokenStat.TotalActiveTokens, TokenStat.TotalInactiveTokens)
	if err != nil {
		return nil,  errors.Wrap(err, "unable to parse stats from db")
	}	

	return stats, nil
}

func (s SQLiteTokenRepository) DeleteToken(token string) error {
	return s.deleteToken(token)
}

func (s SQLiteTokenRepository) deleteToken(token string) error {
	insert := "DELETE FROM tokens WHERE token_id = (?) RETURNING token_id"

	err := s.db.QueryRow(insert, token).Scan(&token)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.Wrap(err, "Cannot find token from database")
	} else if err != nil {
		return errors.Wrap(err, "Unable to delete token from database")
	}

	return nil
}