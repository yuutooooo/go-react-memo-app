package model

import "time"

type Note struct {
	id        string
	title     string
	content   string
	createdAt time.Time
	updatedAt time.Time
	folderID  string
	userID    string
}

// NewNote Noteエンティティを作成
func NewNote(title, content, folderID, userID string) *Note {
	return &Note{
		title:     title,
		content:   content,
		folderID:  folderID,
		userID:    userID,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

// Getters
func (n *Note) ID() string           { return n.id }
func (n *Note) Title() string        { return n.title }
func (n *Note) Content() string      { return n.content }
func (n *Note) FolderID() string     { return n.folderID }
func (n *Note) UserID() string       { return n.userID }
func (n *Note) CreatedAt() time.Time { return n.createdAt }
func (n *Note) UpdatedAt() time.Time { return n.updatedAt }

// Setters
func (n *Note) SetTitle(title string) {
	n.title = title
	n.updatedAt = time.Now()
}
func (n *Note) SetContent(content string) {
	n.content = content
	n.updatedAt = time.Now()
}
func (n *Note) SetFolderID(folderID string) {
	n.folderID = folderID
	n.updatedAt = time.Now()
}
