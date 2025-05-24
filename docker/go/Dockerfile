FROM golang:1.24-alpine

# 作業ディレクトリを設定
WORKDIR /app

# 必要なパッケージをインストール
RUN apk update && \
    apk add --no-cache git && \
    rm -rf /var/cache/apk/*

# Air（ホットリロード用）をインストール
RUN go install github.com/air-verse/air@latest

# delveデバッガーをインストール
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# ポート8080（アプリ）と2345（デバッガー）を公開
EXPOSE 8080 2345

# 開発時のコマンド（airによるホットリロード）
CMD ["air", "-c", ".air.toml"]
