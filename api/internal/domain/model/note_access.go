package model

import "time"

// NoteAccess ノートへのアクセス権限エンティティ
type NoteAccess struct {
	noteID    string
	userID    string
	createdAt time.Time
}

// NewNoteAccess NoteAccessエンティティを作成
func NewNoteAccess(noteID, userID string) *NoteAccess {
	return &NoteAccess{
		noteID:    noteID,
		userID:    userID,
		createdAt: time.Now(),
	}
}

// Getters
func (na *NoteAccess) NoteID() string       { return na.noteID }
func (na *NoteAccess) UserID() string       { return na.userID }
func (na *NoteAccess) CreatedAt() time.Time { return na.createdAt }
