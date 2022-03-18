package adapters_test

import (
	"context"
	"testing"

	"github.com/jerensl/jerens-web-api/internal/adapters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newSqlLiteRepository(t *testing.T) *adapters.SQLiteTokenRepository {
	db, err := adapters.NewSQLiteConnection()
	require.NoError(t, err)

	return adapters.NewSQLiteTokenRepository(db)
}

func TestRepository(t *testing.T) {
	r := newSqlLiteRepository(t)

	t.Run("Test update token", func(t *testing.T) {
		t.Skip()
		testUpdatedToken(t, r)
	})

	t.Run("Test update token", func(t *testing.T) {
		testGetToken(t, r)
	})
}

func testUpdatedToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	ctx := context.Background()
	err := repository.UpdatedToken(ctx, "abc123")
	require.NoError(t, err)
}

func testGetToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	ctx := context.Background()

	hasValue, err := repository.GetToken(ctx, "abc123")
	require.NoError(t, err)

	assert.True(t, hasValue)
}