package notification

import (
	"errors"
)

type Token struct {
	tokenID string
	isActive bool
	updatedAt int64
}

func NewToken(tokenID string, isActive bool, updateAt int64) (*Token, error) {
	if tokenID == "" {
		return nil, errors.New("empty token id")
	}

	if updateAt < 1 {
		return nil, errors.New("invalid date")
	}

	return &Token{
		tokenID: tokenID,
		isActive: isActive,
		updatedAt: updateAt,
	}, nil
}

func UnmarshalTokenFromDatabase(tokenID string, isActive bool, updateAt int64) (*Token, error) {
	token, err := NewToken(tokenID, isActive, updateAt)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (t Token) TokenID() string {
	return t.tokenID
}

func (t Token) IsActive() bool {
	return t.isActive
}

func (t Token) UpdatedAt() int64 {
	return t.updatedAt
}