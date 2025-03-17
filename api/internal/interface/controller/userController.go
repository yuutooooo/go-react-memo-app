package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-react-memo-app/internal/interface/dto"
	"github.com/yourusername/go-react-memo-app/internal/usecase"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) GetAllUser(e echo.Context) error {
	users, err := c.userUsecase.GetAllUser()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.CreateUserResponse(users)
	return e.JSON(http.StatusOK, response)
}
