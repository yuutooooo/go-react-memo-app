package dto

import (
	"time"

	"github.com/yourusername/go-react-memo-app/internal/domain/model"
)

type FolderResponse struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Path           string    `json:"path"`
	ParentFolderID string    `json:"parentFolderID"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type FolderResponseList struct {
	Folders []*FolderResponse `json:"folders"`
}

func CreateFolderResponseMany(folders []*model.Folder) *FolderResponseList {
	folderDTOs := make([]*FolderResponse, len(folders))
	for i, folder := range folders {
		var parentID string
		if folder.ParentFolderID() != nil {
			parentID = *folder.ParentFolderID()
		}

		folderDTOs[i] = &FolderResponse{
			ID:             folder.ID(),
			Name:           folder.Name(),
			Path:           folder.Path(),
			ParentFolderID: parentID,
			CreatedAt:      folder.CreatedAt(),
			UpdatedAt:      folder.UpdatedAt(),
		}
	}
	return &FolderResponseList{
		Folders: folderDTOs,
	}
}

func GetFolderByIDResponse(folder *model.Folder) *FolderResponse {
	var parentID string
	if folder.ParentFolderID() != nil {
		parentID = *folder.ParentFolderID()
	}
	return &FolderResponse{
		ID:             folder.ID(),
		Name:           folder.Name(),
		Path:           folder.Path(),
		ParentFolderID: parentID,
		CreatedAt:      folder.CreatedAt(),
		UpdatedAt:      folder.UpdatedAt(),
	}
}

type FolderNoteTree struct {
	Folder   *FolderResponse   `json:"folder"`
	Notes    []*NoteResponse   `json:"notes"`
	Children []*FolderNoteTree `json:"children"`
}

// ここに変換関数を追加
func ConvertToFolderResponse(folder model.Folder) FolderResponse {
	var parentID string
	if folder.ParentFolderID() != nil {
		parentID = *folder.ParentFolderID()
	}

	return FolderResponse{
		ID:             folder.ID(),
		Name:           folder.Name(),
		Path:           folder.Path(),
		ParentFolderID: parentID,
		CreatedAt:      folder.CreatedAt(),
		UpdatedAt:      folder.UpdatedAt(),
	}
}

// NoteをNoteResponseに変換する関数
func ConvertToNoteResponse(note model.Note) NoteResponse {
	return NoteResponse{
		ID:        note.ID(),
		Title:     note.Title(),
		Content:   note.Content(),
		FolderID:  note.FolderID(),
		CreatedAt: note.CreatedAt(),
		UpdatedAt: note.UpdatedAt(),
	}
}

// フォルダとノートのツリーを構築
func CreateFolderNoteTree(folders []*model.Folder, notes []*model.Note) ([]*FolderNoteTree, error) {
	// キー: id
	// 値: フォルダ
	// 親IDごとにフォルダをマッピング
	folderByParentID := make(map[string][]*model.Folder)
	for _, folder := range folders {
		parentID := ""
		if folder.ParentFolderID() != nil {
			parentID = *folder.ParentFolderID()
		}
		folderByParentID[parentID] = append(folderByParentID[parentID], folder)
	}

	// キー: id
	// 値: ノート
	// フォルダIDごとにノートをマッピング
	notesByFolderID := make(map[string][]*model.Note)
	for _, note := range notes {
		folderID := note.FolderID()
		notesByFolderID[folderID] = append(notesByFolderID[folderID], note)
	}

	// ルートレベルのツリーを構築
	rootTree := buildFolderTree(folderByParentID, notesByFolderID, "")

	return rootTree, nil
}

// 再帰的にフォルダツリーを構築する補助関数
func buildFolderTree(folderByParentID map[string][]*model.Folder, notesByFolderID map[string][]*model.Note, parentID string) []*FolderNoteTree {
	result := make([]*FolderNoteTree, 0)

	// 現在の親IDに属するフォルダを取得
	folders, exists := folderByParentID[parentID]
	if !exists {
		return result
	}

	// 各フォルダについて処理
	for _, folder := range folders {
		// フォルダの応答DTOを作成
		folderResponse := ConvertToFolderResponse(*folder)

		// このフォルダに属するノートを取得
		folderNotes := notesByFolderID[folder.ID()]
		noteResponses := make([]*NoteResponse, 0)
		for _, note := range folderNotes {
			noteResponse := ConvertToNoteResponse(*note)
			noteResponses = append(noteResponses, &noteResponse)
		}

		// 子フォルダを再帰的に構築
		children := buildFolderTree(folderByParentID, notesByFolderID, folder.ID())

		// 現在のフォルダ、ノート、子フォルダでツリーノードを作成
		treeNode := FolderNoteTree{
			Folder:   &folderResponse,
			Notes:    noteResponses,
			Children: children,
		}

		result = append(result, &treeNode)
	}

	return result
}
