package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/route"
)

func main() {
	// Airのホットリロード環境では /go/src/app/.env に配置される
	envPath := ".env"
	if os.Getenv("ENV") == "docker" {
		// Docker環境では同じディレクトリ内の.envを使用
		envPath = ".env"
		log.Println("Docker環境: .env ファイルを現在のディレクトリから読み込みます")
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

	// CORSミドルウェアを設定
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000" // デフォルト値
		log.Printf("FRONTEND_URL環境変数が設定されていないため、デフォルト値を使用します: %s", frontendURL)
	}
	log.Printf("CORS設定: フロントエンドURL = %s", frontendURL)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
		MaxAge:           86400, // プリフライトリクエストの結果をキャッシュする秒数（24時間）
	}))

	route.SetupRoutes(e)

	e.Start(":8080")
}
