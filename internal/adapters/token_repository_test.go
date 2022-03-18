package adapters_test

import (
	"testing"

	"github.com/jerensl/jerens-web-api/internal/adapters"
	"github.com/stretchr/testify/require"
)

func newSqlLiteRepository(t *testing.T) *adapters.SQLiteTokenRepository {
	db, err := adapters.NewSQLiteConnection()
	require.NoError(t, err)

	return adapters.NewSQLiteTokenRepository(db)
}

func TestRepository(t *testing.T) {
	newSqlLiteRepository(t)
}

