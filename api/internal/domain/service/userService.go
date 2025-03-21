package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	"github.com/yourusername/go-react-memo-app/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUser() ([]model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	CreateToken(user *model.User) (string, error)
	HashPassword(user *model.User) error
	CheckPassword(user *model.User, password string) error
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
	return users, nil
}

func (s *userService) CreateUser(user *model.User) (*model.User, error) {
	// ドメインサービス内で新たなユーザーを作成する。
	newUser := model.NewUser(user.Name(), user.Email(), user.Password())
	user, err := s.userRepository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) CreateToken(user *model.User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("secret key is not set")
	}
	println("service debugです")
	println(user.ID())

	// クレームを設定
	claims := jwt.MapClaims{
		"sub": user.ID(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// トークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *userService) HashPassword(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password()), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.SetPassword(string(hashedPassword))
	return nil
}

func (s *userService) CheckPassword(user *model.User, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(password)); err != nil {
		return err
	}
	return nil
}
