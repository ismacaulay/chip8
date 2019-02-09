FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git alpine-sdk

RUN mkdir -p /go/src/github.com/ismacaulay/chip8
WORKDIR /go/src/github.com/ismacaulay/chip8
COPY . .

RUN go get ./... \
    && go get github.com/golang/mock/gomock

WORKDIR /go/src/github.com/ismacaulay/chip8

CMD ["go", "build" "."]
