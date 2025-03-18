package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/route"
)

func main() {
	// Airのホットリロード環境では /go/src/app/.env に配置される
	envPath := ".env"
	if os.Getenv("ENV") == "docker" {
		// Docker環境では絶対パスを使用
		envPath = "/go/src/app/.env"
	} else if _, err := os.Stat(".env"); err == nil {
		// カレントディレクトリに.envがある場合
		envPath = ".env"
	} else {
		// 親ディレクトリの.envを読み込む
		envPath = filepath.Join("..", ".env")
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Warning: Error loading .env file from %s: %v", envPath, err)
		log.Println("Using environment variables from Docker Compose")
	}

	// 環境変数のデバッグ出力
	log.Printf("JWT_SECRET 環境変数: [%s]", os.Getenv("JWT_SECRET"))
	log.Printf("DB_HOST 環境変数: [%s]", os.Getenv("DB_HOST"))
	log.Printf("ENV 環境変数: [%s]", os.Getenv("ENV"))

	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!!!")
	// })
	route.SetupRoutes(e)

	e.Start(":8080")
}
