build:
	go build -o build cmd/main.go

test-cover:
	go test -covermode=atomic -coverpkg=./... -coverprofile cover.out ./...