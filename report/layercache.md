## slow

```
[+] Building 38.2s (10/10) FINISHED                                              docker:desktop-linux
 => [internal] load build definition from Dockerfile.slow                                        0.0s
 => => transferring dockerfile: 348B                                                             0.0s
 => [internal] load .dockerignore                                                                0.0s
 => => transferring context: 2B                                                                  0.0s
 => [internal] load metadata for docker.io/library/golang:1.21                                   0.3s
 => [1/5] FROM docker.io/library/golang:1.21@sha256:4746d26432a9117a5f58e95cb9f954ddf0de128e9d5  0.0s
 => [internal] load build context                                                                0.0s
 => => transferring context: 1.86kB                                                              0.0s
 => CACHED [2/5] WORKDIR /app/hello                                                              0.0s
 => [3/5] COPY . .                                                                               0.1s
 => [4/5] RUN cd hello && go mod download                                                       32.0s
 　　→依存に変化がない場合でも、毎回ダウンロードしてレイヤーを生成する。
 => [5/5] RUN cd hello && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/hello-app main  5.3s
 => exporting to image                                                                           0.5s
 => => exporting layers                                                                          0.5s
 => => writing image sha256:5a4c7976b918f64f4f6afc7d055db33111b0357778fd4abf7aeca32d113a24e0     0.0s
 => => naming to docker.io/library/hello-slow                                                    0.0s
```

## fast

```
[+] Building 7.1s (12/12) FINISHED                                               docker:desktop-linux
 => [internal] load build definition from Dockerfile.fast                                        0.0s
 => => transferring dockerfile: 450B                                                             0.0s
 => [internal] load .dockerignore                                                                0.0s
 => => transferring context: 2B                                                                  0.0s
 => [internal] load metadata for docker.io/library/golang:1.21                                   0.3s
 => [1/7] FROM docker.io/library/golang:1.21@sha256:4746d26432a9117a5f58e95cb9f954ddf0de128e9d5  0.0s
 => [internal] load build context                                                                0.0s
 => => transferring context: 1.79kB                                                              0.0s
 => CACHED [2/7] WORKDIR /app/hello                                                              0.0s
 => CACHED [3/7] COPY common /app/common                                                         0.0s
 => CACHED [4/7] COPY hello/go.mod hello/go.sum* ./                                              0.0s
 => CACHED [5/7] RUN go mod download                                                             0.0s
 　　→依存に変化がない場合、レイヤーキャッシュを使用してダウンロードをスキップする。
 => [6/7] COPY hello .                                                                           0.0s
 => [7/7] RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/hello-app main.go           6.4s
 => exporting to image                                                                           0.3s
 => => exporting layers                                                                          0.3s
 => => writing image sha256:89b2ba224c58bd49dbb91772090a31056df7e2f589716ffadec34680e14f38be     0.0s
 => => naming to docker.io/library/hello-fast
```
