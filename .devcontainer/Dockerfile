FROM golang:1.23.3-bookworm

RUN apt-get update && apt-get install -y \
    xdg-utils

RUN go env -w GOPROXY=direct
RUN go install github.com/go-delve/delve/cmd/dlv@latest
