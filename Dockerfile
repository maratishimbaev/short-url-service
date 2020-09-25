FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN apk add make && make build

CMD sleep 10 && ./build