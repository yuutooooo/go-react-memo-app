package usecase

import (
	"log"

	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	"github.com/yourusername/go-react-memo-app/internal/domain/service"
	"github.com/yourusername/go-react-memo-app/internal/interface/dto"
)

type FolderUsecase struct {
	folderService service.FolderService
}

func NewFolderUsecase(folderService service.FolderService) FolderUsecase {
	return FolderUsecase{
		folderService: folderService,
	}
}

func (u *FolderUsecase) CreateFolder(folder *dto.FolderRequestDTO, userID string) ([]*model.Folder, error) {
	// ParentFolderIDの処理
	var parentIDPtr *string
	if folder.ParentFolderID != "" {
		parentIDPtr = &folder.ParentFolderID
	}

	log.Println("ParentFolderID (DTO):", folder.ParentFolderID)
	modelFolder := model.NewFolder(folder.Name, folder.Path, userID, parentIDPtr)
	log.Println("usecase")
	log.Println(*modelFolder.ParentFolderID())// この形式で参照したらいける

	// ポインタをデリファレンスして値を表示
	if parentID := modelFolder.ParentFolderID(); parentID != nil {
		log.Println("ParentFolderID (モデル):", *parentID)
	} else {
		log.Println("ParentFolderID (モデル): nil")
	}
	createdFolder, err := u.folderService.CreateFolder(modelFolder)
	if err != nil {
		return nil, err
	}
	// 作成されたフォルダと同じ階層のリストの取得
	if createdFolder.ParentFolderID() != nil {
		folderGroup, err := u.folderService.GetFolderByParentFolderID(createdFolder.ParentFolderID(), userID, createdFolder.Path())
		if err != nil {
			return nil, err
		}
		return folderGroup, nil
	} else {
		folderGroup, err := u.folderService.GetRootFolder(userID, createdFolder.Path())
		if err != nil {
			return nil, err
		}
		return folderGroup, nil
	}
}

func (u *FolderUsecase) GetFolderByID(id string) (*model.Folder, error) {
	folder, err := u.folderService.GetFolderByID(id)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (u *FolderUsecase) UpdateFolder(id string, folder *dto.FolderRequestDTO, userID string) ([]*model.Folder, error) {
	// ParentFolderIDの処理
	var parentIDPtr *string
	if folder.ParentFolderID != "" {
		parentIDPtr = &folder.ParentFolderID
	}

	log.Println("ParentFolderID (DTO):", folder.ParentFolderID)
	modelFolder := model.NewFolder(folder.Name, folder.Path, userID, parentIDPtr)
	log.Println("usecase")

	// ポインタをデリファレンスして値を表示
	if parentID := modelFolder.ParentFolderID(); parentID != nil {
		log.Println("ParentFolderID (モデル):", *parentID)
	} else {
		log.Println("ParentFolderID (モデル): nil")
	}
	updatedFolder, err := u.folderService.UpdateFolder(modelFolder, id)
	if err != nil {
		return nil, err
	}
	// 更新されたフォルダと同じ階層のリストの取得
	if updatedFolder.ParentFolderID() != nil {
		folderGroup, err := u.folderService.GetFolderByParentFolderID(updatedFolder.ParentFolderID(), userID, updatedFolder.Path())
		if err != nil {
			return nil, err
		}
		return folderGroup, nil
	} else {
		folderGroup, err := u.folderService.GetRootFolder(userID, updatedFolder.Path())
		if err != nil {
			return nil, err
		}
		return folderGroup, nil
	}
}

func (u *FolderUsecase) DeleteFolder(id string, userID string) ([]*model.Folder, error) {
	err := u.folderService.DeleteFolder(id)
	if err != nil {
		return nil, err
	}
	// 削除されたフォルダと同じ階層のリストの取得
	folderGroup, err := u.folderService.GetFolderByParentFolderID(nil, userID, "")
	if err != nil {
		return nil, err
	}
	return folderGroup, nil
}
