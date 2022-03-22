package adapters_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jerensl/jerens-web-api/internal/adapters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newSqlLiteRepository(t *testing.T, dbPath string) *adapters.SQLiteTokenRepository {
	db, err := adapters.NewSQLiteConnection(dbPath)
	require.NoError(t, err)

	return adapters.NewSQLiteTokenRepository(db)
}

func TestRepository(t *testing.T) {
	dbPath := "../../database/db_test.sqlite"
	r := newSqlLiteRepository(t, dbPath)

	t.Run("Test update token", func(t *testing.T) {
		testUpdatedToken(t, r)
	})

	t.Run("Test update token", func(t *testing.T) {
		testGetToken(t, r)
	})

	t.Run("Test update token", func(t *testing.T) {
		testGetAllToken(t, r)
	})

	err := os.Remove(dbPath)
	if err != nil {
		fmt.Println("cannot remove database")
	}
}

func testUpdatedToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	ctx := context.Background()
	err := repository.UpdatedToken(ctx, "abc123")
	require.NoError(t, err)
	err = repository.UpdatedToken(ctx, "abc321")
	require.NoError(t, err)
}

func testGetToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	ctx := context.Background()

	hasValue, err := repository.GetToken(ctx, "abc123")
	require.NoError(t, err)

	assert.True(t, hasValue)
}

func testGetAllToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	ctx := context.Background()
	expected := []string{"abc123", "abc321"}

	subscriber, err := repository.GetAllToken(ctx)
	require.NoError(t, err)

	assert.Equal(t, expected,subscriber)
}