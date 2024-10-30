FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .
RUN go build -o verification
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/verification .
RUN apk update && apk add --no-cache libc6-compat
ENTRYPOINT [ "./verification" ]
