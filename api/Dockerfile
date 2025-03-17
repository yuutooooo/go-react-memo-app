# ベースイメージは公式の golang イメージを利用
FROM golang:1.23

# Air をインストール
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# 作業ディレクトリを作成
WORKDIR /go/src/app

# go.mod と go.sum をコピー
COPY go.mod go.sum* ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# デフォルトコマンドは Air を起動（ホットリロード）
CMD ["air", "-c", ".air.toml"]
