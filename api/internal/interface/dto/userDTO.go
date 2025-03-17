package dto

import (
	"time"

	"github.com/yourusername/go-react-memo-app/internal/domain/model"
)

type UserListResponse struct {
	Users []User `json:"users"`
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUserResponse(users []model.User) UserListResponse {
	userResponses := make([]User, len(users))
	for i, user := range users {
		userResponses[i] = User{
			ID:        user.ID(),
			Name:      user.Name(),
			Email:     user.Email(),
			CreatedAt: user.CreatedAt(),
			UpdatedAt: user.UpdatedAt(),
		}
	}
	return UserListResponse{
		Users: userResponses,
	}
}
