# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN apk update && apk upgrade
RUN apk add gcc musl-dev pkgconfig zeromq-dev


WORKDIR /app

COPY srcs/go.mod ./
COPY srcs/go.sum ./
RUN go mod download
COPY srcs/server.go ./

RUN go build -o server.out

EXPOSE 3000 5555

CMD ["./server.out"]
