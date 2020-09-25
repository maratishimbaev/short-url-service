FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go build -o build cmd/main.go

CMD sleep 10 && ./build