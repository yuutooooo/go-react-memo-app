package usecase

import (
	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	"github.com/yourusername/go-react-memo-app/internal/domain/service"
)

type UserUsecase struct {
	userService service.UserService
}

func NewUserUsecase(userService service.UserService) UserUsecase {
	return UserUsecase{
		userService: userService,
	}
}

func (u *UserUsecase) GetAllUser() ([]model.User, error) {
	users, err := u.userService.GetAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}
