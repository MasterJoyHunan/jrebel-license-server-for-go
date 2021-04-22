# 构建
FROM golang:1.14-alpine as builder
WORKDIR /home/jrebel
ENV GOPROXY=https://goproxy.cn
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o jrebel

# 打包
FROM alpine as runner
COPY --from=builder /home/jrebel/jrebel /home/jrebel/jrebel
WORKDIR /home/jrebel
CMD ["./jrebel", "-h", "0.0.0.0", "-p", "8877"]