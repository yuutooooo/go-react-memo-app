package persistence

import (
	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	gormmodel "github.com/yourusername/go-react-memo-app/internal/domain/model/gormModel"
	"github.com/yourusername/go-react-memo-app/internal/domain/repository"
	"gorm.io/gorm"
)

func NewFolderRepository(db *gorm.DB) repository.FolderRepository {
	return &FolderRepositoryImpl{
		db: db,
	}
}

type FolderRepositoryImpl struct {
	db *gorm.DB
}

func (r *FolderRepositoryImpl) CreateFolder(folder *model.Folder) (*model.Folder, error) {
	gormFolder := convertToGormFolder(folder)
	if err := r.db.Create(gormFolder).Error; err != nil {
		return nil, err
	}
	folderModel := convertToFolderModel(*gormFolder)
	return folderModel, nil
}

func (r *FolderRepositoryImpl) GetFolderByParentFolderID(parentFolderID *string, userID string, path string) ([]*model.Folder, error) {
	var gormFolders []gormmodel.Folder
	if err := r.db.Where("parent_folder_id = ? AND user_id = ? AND path = ?", parentFolderID, userID, path).Find(&gormFolders).Error; err != nil {
		return nil, err
	}
	return convertToFolderModelList(gormFolders), nil
}

func (r *FolderRepositoryImpl) GetRootFolder(userID string, path string) ([]*model.Folder, error) {
	var gormFolders []gormmodel.Folder
	if err := r.db.Where("user_id = ? AND path = ?", userID, path).Find(&gormFolders).Error; err != nil {
		return nil, err
	}
	return convertToFolderModelList(gormFolders), nil
}

func (r *FolderRepositoryImpl) UpdateFolder(folder *model.Folder, id string) (*model.Folder, error) {
	gormFolder := convertToGormFolder(folder)
	if err := r.db.Where("id = ?", id).Save(gormFolder).Error; err != nil {
		return nil, err
	}
	folderModel := convertToFolderModel(*gormFolder)
	return folderModel, nil
}

func (r *FolderRepositoryImpl) DeleteFolder(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&gormmodel.Folder{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *FolderRepositoryImpl) GetFolderByID(id string) (*model.Folder, error) {
	var gormFolder gormmodel.Folder
	if err := r.db.Where("id = ?", id).First(&gormFolder).Error; err != nil {
		return nil, err
	}
	return convertToFolderModel(gormFolder), nil
}

func convertToGormFolder(folder *model.Folder) *gormmodel.Folder {
	var gormFolder gormmodel.Folder
	var parentID *string
	if folder.ParentFolderID() != nil {
		parentID = folder.ParentFolderID()
	}
	gormFolder.ID = folder.ID()
	gormFolder.Name = folder.Name()
	gormFolder.UserID = folder.UserID()
	gormFolder.Path = folder.Path()
	gormFolder.ParentFolderID = parentID
	return &gormFolder
}

func convertToFolderModel(gormFolder gormmodel.Folder) *model.Folder {
	var folder model.Folder
	folder.SetID(gormFolder.ID)
	folder.SetName(gormFolder.Name)
	folder.SetPath(gormFolder.Path)
	folder.SetUserID(gormFolder.UserID)
	folder.SetCreatedAt(gormFolder.CreatedAt)
	folder.SetUpdatedAt(gormFolder.UpdatedAt)
	// 親フォルダのフィールドのみnil許容
	if gormFolder.ParentFolderID != nil {
		folder.SetParentFolderID(gormFolder.ParentFolderID)
	}
	return &folder
}

func convertToFolderModelList(gormFolders []gormmodel.Folder) []*model.Folder {
	folders := make([]*model.Folder, len(gormFolders))
	for i, gormFolder := range gormFolders {
		folders[i] = convertToFolderModel(gormFolder)
	}
	return folders
}
