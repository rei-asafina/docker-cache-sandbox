# syntax=docker/dockerfile:1.4

FROM golang:1.21 AS build

WORKDIR /app/hello

# 1. ローカルのモジュール（replace先など）
COPY common /app/common

# 2. go.mod/go.sum だけコピー
COPY hello/go.mod hello/go.sum* ./

# 3. go.mod/go.sum を安定化させる
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod tidy

# 4. 依存をすべてダウンロード
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# 5. アプリ本体コピー
COPY hello .

# 6. ビルド（依存キャッシュ＋ビルドキャッシュの両方を持続）
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /app/hello-app main.go

CMD ["/app/hello-app"]
