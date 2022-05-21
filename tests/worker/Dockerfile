# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN apk update && apk upgrade
RUN apk add gcc musl-dev pkgconfig zeromq-dev

WORKDIR /app

COPY srcs/go.mod ./
COPY srcs/go.sum ./
RUN go mod download
COPY srcs/worker.go ./

RUN go build -o worker.out

CMD ["./worker.out"]
