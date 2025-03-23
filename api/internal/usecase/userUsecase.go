package usecase

import (
	"errors"

	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	"github.com/yourusername/go-react-memo-app/internal/domain/repository"
	"github.com/yourusername/go-react-memo-app/internal/domain/service"
	"github.com/yourusername/go-react-memo-app/internal/interface/dto"
)

type UserUsecase struct {
	userService    service.UserService
	userRepository repository.UserRepository
	folderService  service.FolderService
	noteService    service.NoteService
}

func NewUserUsecase(userService service.UserService, userRepository repository.UserRepository, folderService service.FolderService, noteService service.NoteService) UserUsecase {
	return UserUsecase{
		userService:    userService,
		userRepository: userRepository,
		folderService:  folderService,
		noteService:    noteService,
	}
}

func (u *UserUsecase) GetAllUser() ([]model.User, error) {
	users, err := u.userService.GetAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserUsecase) CreateUser(req dto.CreateUserReq) (*model.User, string, error) {
	user := model.NewUser(req.Name, req.Email, req.Password)
	if err := u.userService.HashPassword(user); err != nil {
		return nil, "", err
	}
	user, err := u.userService.CreateUser(user)
	if err != nil {
		return nil, "", err
	}
	token, err := u.userService.CreateToken(user)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (u *UserUsecase) CheckEmail(email string) error {
	user, _ := u.userRepository.FindByEmail(email)
	if user != nil {
		return errors.New("email already exists")
	}
	return nil
}


func (u *UserUsecase) Login(req dto.LoginUserReq) (*model.User, string, error) {
	user, err := u.userRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, "", err
	}
	if err := u.userService.CheckPassword(user, req.Password); err != nil {
		return nil, "", err
	}
	token, err := u.userService.CreateToken(user)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (u *UserUsecase) GetUserById(id string) (*model.User, error) {
	user, err := u.userService.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) GetAllFolderAndNoteByUserID(userID string) ([]*model.Folder, []*model.Note, error) {
	folders, err := u.folderService.GetFolderByUserID(userID)
	if err != nil {
		return nil, nil, err
	}
	notes, err := u.noteService.GetNoteByUserID(userID)
	if err != nil {
		return nil, nil, err
	}
	return folders, notes, nil
}