package repository

import "github.com/yourusername/go-react-memo-app/internal/domain/model"

type UserRepository interface {
	GetAllUser() ([]model.User, error)
}
