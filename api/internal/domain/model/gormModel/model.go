package gormmodel

import (
	"time"

	"gorm.io/datatypes"
)

// User ユーザー
type User struct {
	ID           string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name         string
	Email        string `gorm:"unique"`
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Folders      []Folder     `gorm:"foreignKey:UserID"`
	Notes        []Note       `gorm:"foreignKey:UserID"`
	NoteAccesses []NoteAccess `gorm:"foreignKey:UserID"`
}

// Folder フォルダ（ディレクトリ構造）
type Folder struct {
	ID             string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name           string
	Path           string // 例: "/root/仕事/プロジェクトA"
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserID         string
	User           User     `gorm:"foreignKey:UserID"`
	Notes          []Note   `gorm:"foreignKey:FolderID"`
	ParentFolderID *string  `gorm:"type:uuid"` // 親フォルダID（nullable）
	ParentFolder   *Folder  `gorm:"foreignKey:ParentFolderID"`
	Subfolders     []Folder `gorm:"foreignKey:ParentFolderID"` // 子フォルダ群
}

// Note メモ
type Note struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title     string
	Content   string // Markdown 形式
	CreatedAt time.Time
	UpdatedAt time.Time
	FolderID  string
	Folder    Folder `gorm:"foreignKey:FolderID"`
	UserID    string
	User      User          `gorm:"foreignKey:UserID"`
	Images    []Image       `gorm:"foreignKey:NoteID"`
	Versions  []NoteVersion `gorm:"foreignKey:NoteID"`
	Accesses  []NoteAccess  `gorm:"many2many:note_accesses;"`
}

// Image 画像
type Image struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	URL       string
	CreatedAt time.Time
	UpdatedAt time.Time
	NoteID    string
	Note      Note `gorm:"foreignKey:NoteID"`
}

// NoteVersion 編集履歴（差分）
type NoteVersion struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	NoteID    string
	Note      Note           `gorm:"foreignKey:NoteID"`
	Content   datatypes.JSON // 差分情報（JSON形式）
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NoteAccess ノートへのアクセス権限
// 中間テーブルとして、複数のユーザーが特定のノートにアクセスできることを表現
type NoteAccess struct {
	NoteID string `gorm:"primaryKey;type:uuid"`
	UserID string `gorm:"primaryKey;type:uuid"`
	// 例えばアクセスレベル（"read", "write" など）を追加する場合:
	// Role      string
	CreatedAt time.Time
}
