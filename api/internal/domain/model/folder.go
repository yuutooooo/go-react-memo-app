package model

import "time"

type Folder struct {
	id             string
	name           string
	path           string
	createdAt      time.Time
	updatedAt      time.Time
	userID         string
	parentFolderID *string
}

// NewFolder Folderエンティティを作成
func NewFolder(name, path, userID string, parentFolderID *string) *Folder {
	// parentFolderIDの処理: 空文字列の場合はnilにする
	var parentID *string
	if parentFolderID != nil && *parentFolderID != "" {
		parentID = parentFolderID
	}

	return &Folder{
		name:           name,
		path:           path,
		userID:         userID,
		parentFolderID: parentID,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
	}
}

// Getters
func (f *Folder) ID() string              { return f.id }
func (f *Folder) Name() string            { return f.name }
func (f *Folder) Path() string            { return f.path }
func (f *Folder) UserID() string          { return f.userID }
func (f *Folder) ParentFolderID() *string { return f.parentFolderID }
func (f *Folder) CreatedAt() time.Time    { return f.createdAt }
func (f *Folder) UpdatedAt() time.Time    { return f.updatedAt }

// Setters
func (f *Folder) SetID(id string) {
	f.id = id
}
func (f *Folder) SetName(name string) {
	f.name = name
}
func (f *Folder) SetPath(path string) {
	f.path = path
}
func (f *Folder) SetUserID(userID string) {
	f.userID = userID
}
func (f *Folder) SetParentFolderID(parentFolderID *string) {
	f.parentFolderID = parentFolderID
}
func (f *Folder) SetCreatedAt(createdAt time.Time) {
	f.createdAt = createdAt
}
func (f *Folder) SetUpdatedAt(updatedAt time.Time) {
	f.updatedAt = updatedAt
}
