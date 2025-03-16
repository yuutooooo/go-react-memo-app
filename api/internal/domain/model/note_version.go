package model

import (
	"encoding/json"
	"time"
)

type NoteVersion struct {
	id        string
	noteID    string
	content   json.RawMessage
	createdAt time.Time
	updatedAt time.Time
}

// NewNoteVersion NoteVersionエンティティを作成
func NewNoteVersion(noteID string, content json.RawMessage) *NoteVersion {
	return &NoteVersion{
		noteID:    noteID,
		content:   content,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

// Getters
func (nv *NoteVersion) ID() string               { return nv.id }
func (nv *NoteVersion) NoteID() string           { return nv.noteID }
func (nv *NoteVersion) Content() json.RawMessage { return nv.content }
func (nv *NoteVersion) CreatedAt() time.Time     { return nv.createdAt }
func (nv *NoteVersion) UpdatedAt() time.Time     { return nv.updatedAt }

// Setters
func (nv *NoteVersion) SetContent(content json.RawMessage) {
	nv.content = content
	nv.updatedAt = time.Now()
}
