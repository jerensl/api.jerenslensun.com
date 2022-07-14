package notification_test

import (
	"testing"

	"github.com/jerensl/api.jerenslensun.com/internal/domain/notification"
	"github.com/stretchr/testify/require"
)

func TestNewStats(t *testing.T) {
	totalSubs := 2
	totalActiveSubs := 1
	totalInactiveSubs := 1
	
	stats, err := notification.NewStats(totalSubs, totalActiveSubs, totalInactiveSubs)
	require.NoError(t, err)

	require.Equal(t, totalSubs, stats.TotalSubs())
	require.Equal(t, totalActiveSubs, stats.TotalActiveSubs())
	require.Equal(t, totalInactiveSubs, stats.TotalInactiveSubs())
}

func TestUnmarshalStatsFromDB(t *testing.T) {
	totalSubs := 2
	totalActiveSubs := 1
	totalInactiveSubs := 1
	
	stats, err := notification.UnmarshalStatsFromDatabase(totalSubs, totalActiveSubs, totalInactiveSubs)
	require.NoError(t, err)

	require.Equal(t, totalSubs, stats.TotalSubs())
	require.Equal(t, totalActiveSubs, stats.TotalActiveSubs())
	require.Equal(t, totalInactiveSubs, stats.TotalInactiveSubs())
}


func TestInvalidNewStats(t *testing.T) {
	totalSubs := 2
	totalActiveSubs := 1
	totalInactiveSubs := 0
	
	_, err := notification.NewStats(-1, totalActiveSubs, totalInactiveSubs)
	require.Error(t, err)

	_, err = notification.NewStats(totalSubs, -1, totalInactiveSubs)
	require.Error(t, err)

	_, err = notification.NewStats(totalSubs, totalActiveSubs, -1)
	require.Error(t, err)

	_, err = notification.NewStats(totalSubs, totalActiveSubs, totalInactiveSubs)
	require.Error(t, err)
}