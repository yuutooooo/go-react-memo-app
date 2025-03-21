package service

import (
	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	"github.com/yourusername/go-react-memo-app/internal/domain/repository"
)

type FolderService interface {
	CreateFolder(folder *model.Folder) (*model.Folder, error)
	GetFolderByParentFolderID(parentFolderID *string, userID string, path string) ([]*model.Folder, error)
	GetRootFolder(userID string, path string) ([]*model.Folder, error)
	GetFolderByID(id string) (*model.Folder, error)
	UpdateFolder(folder *model.Folder, id string) (*model.Folder, error)
	DeleteFolder(id string) error
}

type folderService struct {
	folderRepository repository.FolderRepository
}

func NewFolderService(folderRepository repository.FolderRepository) *folderService {
	return &folderService{
		folderRepository: folderRepository,
	}
}

func (s *folderService) CreateFolder(folder *model.Folder) (*model.Folder, error) {
	createdFolder, err := s.folderRepository.CreateFolder(folder)
	if err != nil {
		return nil, err
	}
	return createdFolder, nil
}

func (s *folderService) GetFolderByParentFolderID(parentFolderID *string, userID string, path string) ([]*model.Folder, error) {
	folders, err := s.folderRepository.GetFolderByParentFolderID(parentFolderID, userID, path)
	if err != nil {
		return nil, err
	}
	return folders, nil
}

func (s *folderService) GetRootFolder(userID string, path string) ([]*model.Folder, error) {
	folders, err := s.folderRepository.GetRootFolder(userID, path)
	if err != nil {
		return nil, err
	}
	return folders, nil
}

func (s *folderService) GetFolderByID(id string) (*model.Folder, error) {
	folder, err := s.folderRepository.GetFolderByID(id)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (s *folderService) UpdateFolder(folder *model.Folder, id string) (*model.Folder, error) {
	updatedFolder, err := s.folderRepository.UpdateFolder(folder, id)
	if err != nil {
		return nil, err
	}
	return updatedFolder, nil
}

func (s *folderService) DeleteFolder(id string) error {
	err := s.folderRepository.DeleteFolder(id)
	if err != nil {
		return err
	}
	return nil
}