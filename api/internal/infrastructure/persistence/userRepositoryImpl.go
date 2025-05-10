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
	return convertToUserModelMany(users), nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user gormmodel.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return convertToUserModel(user), nil
}

func (r *UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	gormUser := convertToGormUser(user)
	if err := r.db.Create(gormUser).Error; err != nil {
		return nil, err
	}
	userModel := convertToUserModel(*gormUser)
	return userModel, nil
}

func (r *UserRepositoryImpl) GetUserById(id string) (*model.User, error) {
	var user gormmodel.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return convertToUserModel(user), nil
}

// gormmodelからusermodelに変換
func convertToUserModelMany(gormUsers []gormmodel.User) []model.User {
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

func convertToUserModel(gormUser gormmodel.User) *model.User {
	userModel := model.User{}
	userModel.SetID(gormUser.ID)
	userModel.SetName(gormUser.Name)
	userModel.SetEmail(gormUser.Email)
	userModel.SetPassword(gormUser.Password)
	userModel.SetCreatedAt(gormUser.CreatedAt)
	userModel.SetUpdatedAt(gormUser.UpdatedAt)
	return &userModel
}

func convertToGormUser(user *model.User) *gormmodel.User {
	gormUser := gormmodel.User{}
	gormUser.ID = user.ID()
	gormUser.Name = user.Name()
	gormUser.Email = user.Email()
	gormUser.Password = user.Password()
	gormUser.CreatedAt = user.CreatedAt()
	gormUser.UpdatedAt = user.UpdatedAt()
	return &gormUser
}
