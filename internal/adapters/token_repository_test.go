package adapters_test

import (
	"os"
	"testing"
	"time"

	"github.com/jerensl/api.jerenslensun.com/internal/adapters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	testCases := m.Run()
	os.Remove(os.Getenv("SQLITE_DB_TEST"))
	os.Exit(testCases)
}


func newSqlLiteRepository(t *testing.T) *adapters.SQLiteTokenRepository {
	db, err := adapters.NewSQLiteConnection(os.Getenv("SQLITE_DB_TEST"))
	require.NoError(t, err)

	return adapters.NewSQLiteTokenRepository(db)
}

func TestRepository(t *testing.T) {
	r := newSqlLiteRepository(t)

	t.Run("Test Insert token", func(t *testing.T) {
		testInsertToken(t, r)
	})

	t.Run("Test Update token", func(t *testing.T) {
		testUpdateToken(t, r)
	})

	t.Run("Test Update token not exist", func(t *testing.T) {
		testUpdateTokenNotExist(t, r)
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

	t.Run("Test Count Statistic token", func(t *testing.T) {
		testCountStatisticToken(t, r)
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

func testUpdateTokenNotExist(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	token, err := repository.UpdatedToken("abc456", time.Now().Unix())
	require.NoError(t, err)
	require.True(t, token.IsActive())

	err = repository.DeleteToken("abc456")
	require.NoError(t, err)
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

func testCountStatisticToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	totalToken := 2
	totalActiveToken := 1
	totalInactiveToken := 1

	stat, err := repository.GetStatisticToken()
	require.NoError(t, err)

	assert.Equal(t, totalToken, stat.TotalSubs())
	assert.Equal(t, totalActiveToken, stat.TotalActiveSubs())
	assert.Equal(t, totalInactiveToken, stat.TotalInactiveSubs())
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