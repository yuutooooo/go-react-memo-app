package service

import (
	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	"github.com/yourusername/go-react-memo-app/internal/domain/repository"
)

type UserService interface {
	GetAllUser() ([]model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetAllUser() ([]model.User, error) {
	users, err := s.userRepository.GetAllUser()
	if err != nil {
		return nil, err
	}
	println("service debug")
	for _, user := range users {
		println(user.ID(), user.Name(), user.Email(), user.Password())
	}
	return users, nil
}
