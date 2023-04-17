# # Builder stage with air
# FROM golang:1.19-alpine AS air_builder
# RUN apk add --no-cache git
# RUN go install github.com/cosmtrek/air@latest

# # Builder stage for the Go application
# FROM golang:1.19-alpine AS builder
# WORKDIR /app
# COPY . .
# RUN go build -o main main.go

# # Final stage with the compiled application and air binary
# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY --from=air_builder /go/bin/air /usr/local/bin/air
# COPY .air.toml /app/.air.toml
# CMD ["air", "run"]

# 上記のよう思考錯誤したが、マルチステージビルド & airはできなさそう
# 理由
# - CMD ["air", "run"]の実行にはgoコマンドがいる
# - そのためには、結局FROM alpine:latestをFROM golang:1.19-alpineに変更する必要がある
# そしたらイメージは別に軽くならない
# というか、そもそもairは開発用=> 開発用とデプロイ用でDockerfileをわける
# -----------------------------------------


# 開発用
  FROM golang:1.19-alpine
  WORKDIR /go/src
  COPY go.mod go.sum ./
  RUN go mod download
  COPY . .
  RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air
  CMD ["air", "-c", ".air.toml"]

# デプロイ用
  # # Builder stage for the Go application
  # FROM golang:1.19-alpine AS builder
  # WORKDIR /app
  # COPY . .
  # RUN go build -o main main.go

  # # Final stage with the compiled application
  # FROM alpine:latest
  # WORKDIR /app
  # COPY --from=builder /app/main .
  # CMD ["/app/main"]
