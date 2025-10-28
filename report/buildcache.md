## god

```
[+] Building 2.5s (15/15) FINISHED                                                                                    docker:desktop-linux
 => [hello internal] load build definition from Dockerfile.god                                                                        0.0s
 => => transferring dockerfile: 821B                                                                                                  0.0s
 => [hello internal] load .dockerignore                                                                                               0.0s
 => => transferring context: 2B                                                                                                       0.0s
 => [hello] resolve image config for docker.io/docker/dockerfile:1.4                                                                  0.8s
 => CACHED [hello] docker-image://docker.io/docker/dockerfile:1.4@sha256:9ba7531bd80fb0a858632727cf7a112fbfd19b17e94c4e84ced81e24ef1  0.0s
 => [hello internal] load metadata for docker.io/library/golang:1.21                                                                  0.6s
 => [hello build 1/8] FROM docker.io/library/golang:1.21@sha256:4746d26432a9117a5f58e95cb9f954ddf0de128e9d5816886514199316e4a2fb      0.0s
 => [hello internal] load build context                                                                                               0.0s
 => => transferring context: 1.78kB                                                                                                   0.0s
 => CACHED [hello build 2/8] WORKDIR /app/hello                                                                                       0.0s
 => CACHED [hello build 3/8] COPY common /app/common                                                                                  0.0s
 => CACHED [hello build 4/8] COPY hello/go.mod hello/go.sum* ./                                                                       0.0s
 => CACHED [hello build 5/8] RUN --mount=type=cache,target=/go/pkg/mod     go mod tidy                                                0.0s
 => CACHED [hello build 6/8] RUN --mount=type=cache,target=/go/pkg/mod     go mod download                                            0.0s
 => [hello build 7/8] COPY hello .                                                                                                    0.0s
 => [hello build 8/8] RUN --mount=type=cache,target=/go/pkg/mod     --mount=type=cache,target=/root/.cache/go-build     CGO_ENABLED=  0.9s

    →Goが必要とする依存モジュールは /go/pkg/mod にすでに入ってるので、再ダウンロードしない
    →直前のビルドで作ったオブジェクトファイルが /root/.cache/go-build に残ってるので、差分だけ再コンパイルする

 => [hello] exporting to image                                                                                                        0.0s
 => => exporting layers                                                                                                               0.0s
 => => writing image sha256:16818e09f02214a01e1cc57035fefeef4bab8899f4d225d712a46a052ee56277                                          0.0s
 => => naming to docker.io/library/hello-god:dev
```

## fast

```
[+] Building 9.0s (12/12) FINISHED                                                                                                                                                    docker:desktop-linux
 => [hello internal] load .dockerignore                                                                                                                                                               0.0s
 => => transferring context: 2B                                                                                                                                                                       0.0s
 => [hello internal] load build definition from Dockerfile.fast                                                                                                                                       0.0s
 => => transferring dockerfile: 516B                                                                                                                                                                  0.0s
 => [hello internal] load metadata for docker.io/library/golang:1.21                                                                                                                                  0.3s
 => [hello 1/7] FROM docker.io/library/golang:1.21@sha256:4746d26432a9117a5f58e95cb9f954ddf0de128e9d5816886514199316e4a2fb                                                                            0.0s
 => [hello internal] load build context                                                                                                                                                               0.0s
 => => transferring context: 1.79kB                                                                                                                                                                   0.0s
 => CACHED [hello 2/7] WORKDIR /app/hello                                                                                                                                                             0.0s
 => CACHED [hello 3/7] COPY common /app/common                                                                                                                                                        0.0s
 => CACHED [hello 4/7] COPY hello/go.mod hello/go.sum* ./                                                                                                                                             0.0s
 => CACHED [hello 5/7] RUN go mod download                                                                                                                                                            0.0s
 => [hello 6/7] COPY hello .                                                                                                                                                                          0.3s
 => [hello 7/7] RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/hello-app main.go                                                                                                          7.8s
 => [hello] exporting to image                                                                                                                                                                        0.5s
 => => exporting layers                                                                                                                                                                               0.5s
 => => writing image sha256:93d6e79d3c396ce23a35a6c40e6b3c62207fef63dc0fa6af7bce697ebd89ad9f                                                                                                          0.0s
 => => naming to docker.io/library/hello-fast:dev                                                                                                                                                     0.0s
```
