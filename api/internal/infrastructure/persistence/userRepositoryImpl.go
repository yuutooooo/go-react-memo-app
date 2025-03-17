package persistence

import (
	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	gormmodel "github.com/yourusername/go-react-memo-app/internal/domain/model/gormModel"
	"github.com/yourusername/go-react-memo-app/internal/domain/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetAllUser() ([]model.User, error) {
	var users []gormmodel.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return convertToUserModel(users), nil
}

// gormmodelからusermodelに変換
func convertToUserModel(gormUsers []gormmodel.User) []model.User {
	userModels := make([]model.User, len(gormUsers))
	for i, user := range gormUsers {
		userModels[i] = model.User{}
		userModels[i].SetID(user.ID)
		userModels[i].SetName(user.Name)
		userModels[i].SetEmail(user.Email)
		userModels[i].SetPassword(user.Password)
		userModels[i].SetCreatedAt(user.CreatedAt)
		userModels[i].SetUpdatedAt(user.UpdatedAt)
	}
	return userModels
}
