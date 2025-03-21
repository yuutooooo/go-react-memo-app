package dto

import (
	"time"

	"github.com/yourusername/go-react-memo-app/internal/domain/model"
)

type FolderResponseDTO struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Path           string    `json:"path"`
	ParentFolderID string    `json:"parentFolderID"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type FolderResponseListDTO struct {
	Folders []*FolderResponseDTO `json:"folders"`
}

func CreateFolderResponseMany(folders []*model.Folder) *FolderResponseListDTO {
	folderDTOs := make([]*FolderResponseDTO, len(folders))
	for i, folder := range folders {
		var parentID string
		if folder.ParentFolderID() != nil {
			parentID = *folder.ParentFolderID()
		}

		folderDTOs[i] = &FolderResponseDTO{
			ID:             folder.ID(),
			Name:           folder.Name(),
			Path:           folder.Path(),
			ParentFolderID: parentID,
			CreatedAt:      folder.CreatedAt(),
			UpdatedAt:      folder.UpdatedAt(),
		}
	}
	return &FolderResponseListDTO{
		Folders: folderDTOs,
	}
}

func GetFolderByIDResponse(folder *model.Folder) *FolderResponseDTO {
	var parentID string
	if folder.ParentFolderID() != nil {
		parentID = *folder.ParentFolderID()
	}
	return &FolderResponseDTO{
		ID:             folder.ID(),
		Name:           folder.Name(),
		Path:           folder.Path(),
		ParentFolderID: parentID,
		CreatedAt:      folder.CreatedAt(),
		UpdatedAt:      folder.UpdatedAt(),
	}
}