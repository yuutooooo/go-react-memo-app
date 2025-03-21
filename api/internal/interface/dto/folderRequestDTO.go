package dto

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
)

type FolderRequestDTO struct {
	Name           string `json:"name"`
	Path           string `json:"path"`
	ParentFolderID string `json:"parent_folder_id"`
}

func NewFolderRequest(ctx echo.Context) (*FolderRequestDTO, error) {
	var newFolderRequestDTO FolderRequestDTO
	if err := ctx.Bind(&newFolderRequestDTO); err != nil {
		return nil, err
	}
	if err := newFolderRequestDTO.Validate(); err != nil {
		return nil, err
	}
	if newFolderRequestDTO.Path != "" {
		if !strings.HasPrefix(newFolderRequestDTO.Path, "/") {
			newFolderRequestDTO.Path = "/" + newFolderRequestDTO.Path
		}
		newFolderRequestDTO.Path = strings.TrimRight(newFolderRequestDTO.Path, "/")
	}
	return &newFolderRequestDTO, nil
}

func (f *FolderRequestDTO) Validate() error {
	if f.Name == "" {
		return errors.New("name is required")
	}
	if f.Path == "" {
		return errors.New("path is required")
	}
	return nil
}
