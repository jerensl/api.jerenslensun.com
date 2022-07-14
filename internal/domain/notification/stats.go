package notification

import "errors"

type Stats struct {
	totalSubs int
	totalActiveSubs int
	totalInactiveSubs int
}

func NewStats(totalSubs, totalActiveSubs, totalInactiveSubs int) (*Stats, error) {
	if totalSubs < 0 {
		return nil, errors.New("total subs cannot be negative")
	}

	if totalActiveSubs < 0 {
		return nil, errors.New("total active subs cannot be negative")
	}

	if totalInactiveSubs < 0 {
		return nil, errors.New("total inactive subs cannot be negative")
	}

	if totalSubs != (totalActiveSubs + totalInactiveSubs) {
		return nil, errors.New("total is not balance")
	}

	return &Stats{
		totalSubs: totalSubs,
		totalActiveSubs: totalActiveSubs,
		totalInactiveSubs: totalInactiveSubs,
	}, nil
}

func (s Stats) TotalSubs() int {
	return s.totalSubs
}


func (s Stats) TotalActiveSubs() int {
	return s.totalActiveSubs
}


func (s Stats) TotalInactiveSubs() int {
	return s.totalInactiveSubs
}

func UnmarshalStatsFromDatabase(totalSubs, totalActiveSubs, totalInactiveSubs int) (*Stats, error) {
	stats, err := NewStats(totalSubs, totalActiveSubs, totalInactiveSubs)
	if err != nil {
		return nil, err
	}

	return stats, nil
}