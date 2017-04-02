FROM golang:1.8-alpine

ADD . /go/src/github.com/yuniersoad/hangman-phrase-service

RUN go install github.com/yuniersoad/hangman-phrase-service

ENTRYPOINT /go/bin/hangman-phrase-service

EXPOSE 8081


