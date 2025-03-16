# Go React メモアプリケーション

このリポジトリはnotionのメモ機能をクローンすることがイメージのメモアプリ
ディレクトリ構造とその中に入っているファイルをメモで編集できる。
マークダウンでメモを書く。
他人との共有機能も実装。
メモをpdfで出力する機能
メモの中には写真やpdfなどのファイルをアップロードすることができるようにしておく。

## 技術スタック

### バックエンド
- 言語: Go
- フレームワーク: Echo
- データベース: Neon (PostgreSQL)
- ファイルストレージ: Cloudflare R2
- デプロイ: Fly.io
-  ORM: gorm

### フロントエンド
- 言語: TypeScript
- フレームワーク: React
- デプロイ: Vercel

## プロジェクト構造

```
.
├── api/            # Goバックエンド
│   ├── main.go    # エントリーポイント
│   └── go.mod     # Go依存関係
│
├── frontend/       # Reactフロントエンド
│   ├── src/       # ソースコード
│   ├── public/    # 静的ファイル
│   └── package.json
│
└── docker-compose.yml  # ローカル開発環境
```

## セットアップ

### 必要条件
- Go 1.21以上
- Node.js 18以上
- Docker & Docker Compose
- Neon アカウント
- Cloudflare R2 アカウント
- Fly.io アカウント
- Vercel アカウント

### ローカル開発環境のセットアップ

1. リポジトリのクローン
```bash
git clone <repository-url>
cd go-react-memo-app
```

2. 環境変数の設定
```bash
# バックエンド (.env)
cp api/.env.example api/.env
# 必要な環境変数を設定

# フロントエンド (.env.local)
cp frontend/.env.example frontend/.env.local
# 必要な環境変数を設定
```

3. アプリケーションの起動
```bash
# 開発環境の起動
docker-compose up -d
```

## デプロイ

### バックエンド (Fly.io)

1. Fly.ioのCLIをインストール
2. アプリケーションの作成とデプロイ
```bash
cd api
fly launch
fly deploy
```

### フロントエンド (Vercel)

1. Vercelにプロジェクトをインポート
2. 環境変数の設定
3. デプロイ

## 主な機能

- メモの作成、編集、削除
- ファイルの添付（Cloudflare R2を使用）
- リアルタイム同期
- タグ付けとカテゴリ分け
- 検索機能


## /apiディレクトリの構成
```
├── Dockerfile
├── go.mod
├── go.sum
├── internal # ここに具体的なコードを書いていく
│   ├── domain # ドメイン層
│   │   ├── model # モデルの定義
│   │   ├── repository # リポジトリのインターフェースを定義
│   │   └── service # ドメインサービスの実装
│   ├── infrastructure # 外部（アプリケーション以外のもの）とのやり取り
│   │   ├── config # DBの接続など
│   │   ├── external # 外部サービスとのやり取りなど（このアプリケーションではデファイルデータの保存など）
│   │   └── persistence # ここで実際にデータベースとのやり取りをする。
│   ├── interface # インターフェース層
│   │   ├── controller # コントローラー層
│   │   ├── dto # データ変換などを担当 ここでバリデーションを実装する。
│   │   └── presenter # レスポンスなどを変換
│   └── usecase
├── main.go # プログラムのエントリーポイント
├── migration # マイグレーションを実行するためのエントリーポイント
│   └── main.go
└── tmp
    ├── air.log
    └── main
```

## DBマイグレーション方法
`/api/migration`に移動して`go run main.go`を実行

## DBデータシード方法
`/api/seed`に移動して`go run seed.go`を実行

以上の手順を踏んで
http://localhost:5050
にアクセスして、
mail: admin@example.com
pass: admin
を入力してサーバーを登録するとデータが登録されているのが確認できる。