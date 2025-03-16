package main

import (
	"log"

	"github.com/joho/godotenv"
	gormmodel "github.com/yourusername/go-react-memo-app/internal/domain/model/gormModel"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/config"
)

// envファイルの位置がそれぞれのファイルによって異なるため、関数で切り出さない
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
	db.AutoMigrate(&gormmodel.User{}, &gormmodel.Folder{}, &gormmodel.Note{}, &gormmodel.Image{}, &gormmodel.NoteVersion{}, &gormmodel.NoteAccess{})
}
