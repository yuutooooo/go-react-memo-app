package dto

import (
	"errors"
	"regexp"

	"github.com/labstack/echo/v4"
)

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserRequest(c echo.Context) (CreateUserReq, error) {
	var req CreateUserReq
	if err := c.Bind(&req); err != nil {
		return CreateUserReq{}, err
	}
	if err := req.validate(); err != nil {
		return CreateUserReq{}, err
	}
	return req, nil
}

func (req *CreateUserReq) validate() error {
	// 空文字の確認
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return errors.New("invalid request")
	}
	// メールアドレスの形式確認
	matched, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, req.Email)
	if err != nil {
		return errors.New("email validation error")
	}
	if !matched {
		return errors.New("invalid email format")
	}
	// パスワードの長さ確認
	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginUserReq) validate() error {
	if req.Email == "" || req.Password == "" {
		return errors.New("invalid request")
	}
	// メールアドレスの形式確認
	matched, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, req.Email)
	if err != nil {
		return errors.New("email validation error")
	}
	if !matched {
		return errors.New("invalid email format")
	}
	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func SigninUserRequest(c echo.Context) (LoginUserReq, error) {
	var req LoginUserReq
	if err := c.Bind(&req); err != nil {
		return LoginUserReq{}, err
	}
	if err := req.validate(); err != nil {
		return LoginUserReq{}, err
	}
	return req, nil
}