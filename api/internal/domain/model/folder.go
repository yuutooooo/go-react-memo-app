package model

import "time"

// Folder フォルダエンティティ
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
	return &Folder{
		name:           name,
		path:           path,
		userID:         userID,
		parentFolderID: parentFolderID,
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
func (f *Folder) SetName(name string) {
	f.name = name
	f.updatedAt = time.Now()
}
func (f *Folder) SetPath(path string) {
	f.path = path
	f.updatedAt = time.Now()
}
func (f *Folder) SetParentFolderID(parentFolderID *string) {
	f.parentFolderID = parentFolderID
	f.updatedAt = time.Now()
}
