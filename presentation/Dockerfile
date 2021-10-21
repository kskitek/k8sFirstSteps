FROM golang:1.15-alpine

RUN apk add git --no-cache

RUN go get golang.org/x/tools/cmd/present

WORKDIR /slides

ADD *.slide ./

CMD ["present", "-http", "0.0.0.0:8080", "-play=false"]
