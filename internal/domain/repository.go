package domain

type Repository interface {
	GetAllToken() ([]string, error)
	GetToken(token string) (bool, error)
	UpdatedToken(token string) error
	DeleteToken(token string) error
} 