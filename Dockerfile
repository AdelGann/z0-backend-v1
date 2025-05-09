FROM golang:1.24.2

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy