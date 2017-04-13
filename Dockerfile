FROM golang:1.8-alpine

ADD . /go/src/github.com/yuniersoad/hangman-phrase-service

RUN apk add --no-cache --virtual .build-deps glide git \
  && cd /go/src/github.com/yuniersoad/hangman-phrase-service \
  && glide install \
  && go install github.com/yuniersoad/hangman-phrase-service \
  && apk del .build-deps

ENTRYPOINT /go/bin/hangman-phrase-service

EXPOSE 8081


