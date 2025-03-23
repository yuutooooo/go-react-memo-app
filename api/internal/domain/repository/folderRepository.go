package repository

import "github.com/yourusername/go-react-memo-app/internal/domain/model"

type FolderRepository interface {
	CreateFolder(folder *model.Folder) (*model.Folder, error)
	GetFolderByParentFolderID(parentFolderID *string, userID string, path string) ([]*model.Folder, error)
	GetRootFolder(userID string, path string) ([]*model.Folder, error)
	UpdateFolder(folder *model.Folder, id string) (*model.Folder, error)
	DeleteFolder(id string) error
	GetFolderByID(id string) (*model.Folder, error)
	GetFolderByUserID(userID string) ([]*model.Folder, error)
}

