package repository

import (
	"github.com/yourusername/go-react-memo-app/internal/domain/model"
)

type UserRepository interface {
	GetAllUser() ([]model.User, error)
	FindByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
}
