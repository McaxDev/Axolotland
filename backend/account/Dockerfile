FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .
RUN go build -o account
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/account .
RUN apk update && apk add --no-cache libc6-compat
ENTRYPOINT [ "./account" ]
