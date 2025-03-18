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

func (c *UserController) GetAllUser(ctx echo.Context) error {
	users, err := c.userUsecase.GetAllUser()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.CreateUserResponseMany(users)
	return ctx.JSON(http.StatusOK, response)
}

func (c *UserController) Signup(ctx echo.Context) error {
	req, err := dto.CreateUserRequest(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.userUsecase.CheckEmail(req.Email); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	// ここまでがバリデーションのチェックなので、usecaseのメソッドを呼び出す（しょうがない）
	// todo どこまで分離するかを再検討

	user, token, err := c.userUsecase.CreateUser(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.CreateUserResponseSingle(*user, token)
	return ctx.JSON(http.StatusOK, response)
}

func (c *UserController) Signin(ctx echo.Context) error {
	req, err := dto.SigninUserRequest(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	user, token, err := c.userUsecase.Login(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.CreateUserResponseSingle(*user, token)
	return ctx.JSON(http.StatusOK, response)
}

func (c *UserController) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "index")
}
