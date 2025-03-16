package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	gormmodel "github.com/yourusername/go-react-memo-app/internal/domain/model/gormModel"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/config"
	"gorm.io/datatypes"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}

func main() {
	db, err := config.NewDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	user1 := gormmodel.User{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "password1",
	}
	user2 := gormmodel.User{
		Name:     "Bob",
		Email:    "bob@example.com",
		Password: "password2",
	}
	db.Create(&user1)
	db.Create(&user2)

	// 2. フォルダ作成（各ユーザーにつき2個）
	// Alice のフォルダ
	folder1 := gormmodel.Folder{
		Name:   "Work",
		Path:   "/root/Work",
		UserID: user1.ID,
	}
	folder2 := gormmodel.Folder{
		Name:   "Personal",
		Path:   "/root/Personal",
		UserID: user1.ID,
	}
	db.Create(&folder1)
	db.Create(&folder2)

	// Bob のフォルダ
	folder3 := gormmodel.Folder{
		Name:   "Projects",
		Path:   "/root/Projects",
		UserID: user2.ID,
	}
	folder4 := gormmodel.Folder{
		Name:   "Hobbies",
		Path:   "/root/Hobbies",
		UserID: user2.ID,
	}
	db.Create(&folder3)
	db.Create(&folder4)

	// 3. サブフォルダ作成（リレーション確認用：Aliceのフォルダ内にサブフォルダ）
	subfolder1 := gormmodel.Folder{
		Name:           "ProjectA",
		Path:           "/root/Work/ProjectA",
		UserID:         user1.ID,
		ParentFolderID: &folder1.ID,
	}
	db.Create(&subfolder1)

	// 4. ノート作成（各ユーザーにつき2個、フォルダとのリレーション）
	note1 := gormmodel.Note{
		Title:    "Meeting Notes",
		Content:  "Meeting notes content...",
		FolderID: folder1.ID,
		UserID:   user1.ID,
	}
	note2 := gormmodel.Note{
		Title:    "Shopping List",
		Content:  "Milk, Eggs, Bread",
		FolderID: folder2.ID,
		UserID:   user1.ID,
	}
	db.Create(&note1)
	db.Create(&note2)

	note3 := gormmodel.Note{
		Title:    "Project Ideas",
		Content:  "Idea 1, Idea 2",
		FolderID: folder3.ID,
		UserID:   user2.ID,
	}
	note4 := gormmodel.Note{
		Title:    "Hobby Plans",
		Content:  "Plan A, Plan B",
		FolderID: folder4.ID,
		UserID:   user2.ID,
	}
	db.Create(&note3)
	db.Create(&note4)

	// 5. 画像作成（各ユーザーのノートに1枚ずつ関連付け）
	image1 := gormmodel.Image{
		URL:    "http://example.com/image1.png",
		NoteID: note1.ID,
	}
	image2 := gormmodel.Image{
		URL:    "http://example.com/image2.png",
		NoteID: note3.ID,
	}
	db.Create(&image1)
	db.Create(&image2)

	// 6. ノートの編集履歴（各ノートに1件、シンプルな差分情報）
	version1 := gormmodel.NoteVersion{
		NoteID:  note1.ID,
		Content: datatypes.JSON([]byte(`{"changes": "Initial version"}`)),
	}
	version2 := gormmodel.NoteVersion{
		NoteID:  note3.ID,
		Content: datatypes.JSON([]byte(`{"changes": "Initial version"}`)),
	}
	db.Create(&version1)
	db.Create(&version2)

	// 7. ノートへのアクセス権限（中間テーブルでリレーション確認）
	// 例として、Alice の note1 を Bob に共有し、Bob の note3 を Alice に共有
	noteAccess1 := gormmodel.NoteAccess{
		NoteID:    note1.ID,
		UserID:    user2.ID,
		CreatedAt: time.Now(),
	}
	noteAccess2 := gormmodel.NoteAccess{
		NoteID:    note3.ID,
		UserID:    user1.ID,
		CreatedAt: time.Now(),
	}
	db.Create(&noteAccess1)
	db.Create(&noteAccess2)

	fmt.Println("Seed data insertion complete")
}
