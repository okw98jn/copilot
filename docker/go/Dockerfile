FROM golang:1.24-alpine

# 作業ディレクトリを設定
WORKDIR /app

# 必要なパッケージをインストール
RUN apk update && \
    apk add --no-cache git && \
    rm -rf /var/cache/apk/*

# 必要なツールをインストール
RUN go install github.com/air-verse/air@latest && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.3 && \
    go install github.com/google/wire/cmd/wire@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

# ポート8080（アプリ）と2345（デバッガー）を公開
EXPOSE 8080 2345

# 開発時のコマンド（airによるホットリロード）
CMD ["air", "-c", ".air.toml"]
