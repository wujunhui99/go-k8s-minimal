# 阶段一：编译
FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .
# CGO_ENABLED=0 确保静态链接，防止在 alpine 运行报错
RUN CGO_ENABLED=0 go build -o server ./cmd/server

# 阶段二：运行
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
