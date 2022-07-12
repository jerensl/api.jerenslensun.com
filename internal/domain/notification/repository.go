package notification

type Repository interface {
	GetAllToken() ([]string, error)
	GetToken(tokenID string) (*Token, bool, error)
	UpdatedToken(tokenID string, updatedAt int64) (*Token, error)
	InsertedToken(tokenID string, updatedAt int64) (*Token, error)
	DeleteToken(tokenID string) error
} 