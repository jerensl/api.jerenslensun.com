package adapters_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jerensl/api.jerenslensun.com/internal/adapters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	testCases := m.Run()
	os.Remove(os.Getenv("../../database/unit_test.sqlite"))
	os.Exit(testCases)
}


func newSqlLiteRepository(t *testing.T) *adapters.SQLiteTokenRepository {
	db, err := adapters.NewSQLiteConnection()
	require.NoError(t, err)

	return adapters.NewSQLiteTokenRepository(db)
}

func TestRepository(t *testing.T) {
	r := newSqlLiteRepository(t)

	t.Run("Test Insert token", func(t *testing.T) {
		testInsertToken(t, r)
	})

	t.Run("Test Update token token", func(t *testing.T) {
		testUpdateToken(t, r)
	})

	t.Run("Test Insert Existing token", func(t *testing.T) {
		testInsertExistingToken(t, r)
	})

	t.Run("Test Get token", func(t *testing.T) {
		testGetToken(t, r)
	})

	t.Run("Test Get token Not Exist", func(t *testing.T) {
		testGetTokenNotExist(t, r)
	})

	t.Run("Test Get All token", func(t *testing.T) {
		testGetAllToken(t, r)
	})

	t.Run("Test Delete token", func(t *testing.T) {
		testDeleteToken(t, r)
	})

	t.Run("Test Delete Token Not Exist", func(t *testing.T) {
		testDeleteTokenNotExist(t, r)
	})

	t.Run("Test Get All token After Delete One", func(t *testing.T) {
		testGetAllTokenAfterDeleteOne(t, r)
	})

	err := os.Remove("../../database/unit_test.sqlite")
	if err != nil {
		fmt.Println("cannot remove database")
	}
}

func testInsertToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	tokenOne, err := repository.InsertedToken("abc123", time.Now().Unix())
	require.Equal(t, "abc123", tokenOne.TokenID())
	require.NoError(t, err)

	tokenTwo, err := repository.InsertedToken("abc321", time.Now().Unix())
	require.Equal(t, "abc321", tokenTwo.TokenID())
	require.NoError(t, err)
}

func testInsertExistingToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	_, err := repository.InsertedToken("abc123", time.Now().Unix())
	require.ErrorContains(t, err, "Unable to insert token to database")
}

func testUpdateToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	token, err := repository.UpdatedToken("abc321", time.Now().Unix())
	require.NoError(t, err)
	require.False(t, token.IsActive())
}

func testGetToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	token, isExist, err := repository.GetToken("abc123")
	require.NoError(t, err)

	assert.True(t, token.IsActive())
	assert.True(t, isExist)
}

func testGetTokenNotExist(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	_, isExist, _ := repository.GetToken("abc1233")
	require.False(t, isExist)
}

func testGetAllToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	expected := []string{"abc123"}

	subscriber, err := repository.GetAllToken()
	require.NoError(t, err)

	assert.Equal(t, expected,subscriber)
}

func testDeleteToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	err := repository.DeleteToken("abc321")
	require.NoError(t, err)
}

func testGetAllTokenAfterDeleteOne(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	expected := []string{"abc123"}

	subscriber, err := repository.GetAllToken()
	require.NoError(t, err)

	assert.Equal(t, expected,subscriber)
}

func testDeleteTokenNotExist(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	err := repository.DeleteToken("abc1232")
	require.ErrorContains(t, err, "Cannot find token from database")
}