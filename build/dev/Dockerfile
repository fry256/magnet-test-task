FROM golang:latest as build_base
WORKDIR /app
# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

# This image builds the weavaite server
FROM build_base AS builder
COPY . .
RUN go build -o magnet cmd/main.go
CMD ["/app/magnet"]
