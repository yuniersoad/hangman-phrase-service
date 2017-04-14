FROM golang:1.8-alpine

ENV target_dir /go/src/github.com/yuniersoad/hangman-phrase-service

RUN mkdir -p $target_dir

COPY ./glide.yaml $target_dir
COPY ./glide.lock $target_dir

RUN apk add --no-cache --virtual .build-deps glide git \
  && cd $target_dir \
  && glide install \
  && apk del .build-deps

COPY . $target_dir

RUN cd $target_dir \
  && go install 

ENTRYPOINT /go/bin/hangman-phrase-service

EXPOSE 8081


