package model

import "time"

// Image 画像エンティティ
type Image struct {
	id        string
	url       string
	createdAt time.Time
	updatedAt time.Time
	noteID    string
}

// NewImage Imageエンティティを作成
func NewImage(url, noteID string) *Image {
	return &Image{
		url:       url,
		noteID:    noteID,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

// Getters
func (i *Image) ID() string           { return i.id }
func (i *Image) URL() string          { return i.url }
func (i *Image) NoteID() string       { return i.noteID }
func (i *Image) CreatedAt() time.Time { return i.createdAt }
func (i *Image) UpdatedAt() time.Time { return i.updatedAt }

// Setters
func (i *Image) SetURL(url string) {
	i.url = url
	i.updatedAt = time.Now()
}
