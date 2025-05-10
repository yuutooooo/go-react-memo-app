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
	Token     string    `json:"token"`
}

func CreateUserResponseMany(users []model.User) UserListResponse {
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

func CreateUserResponseSingle(user model.User, token string) User {
	return User{
		ID:        user.ID(),
		Name:      user.Name(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
		Token:     token,
	}
}

type UserIndexResponse struct {
	User            User            `json:"user"`
	FolderAndNoteTree []*FolderNoteTree `json:"folderAndNoteTree"`
}

func CreateUserIndexResponse(user model.User, folderAndNoteTree []*FolderNoteTree) UserIndexResponse {
	return UserIndexResponse{
		User: User{
			ID:        user.ID(),
			Name:      user.Name(),
			Email:     user.Email(),
			CreatedAt: user.CreatedAt(),
			UpdatedAt: user.UpdatedAt(),
		},
		FolderAndNoteTree: folderAndNoteTree,
	}
}
