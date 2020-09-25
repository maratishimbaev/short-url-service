FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN make build

CMD sleep 10 && ./build