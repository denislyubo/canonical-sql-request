install:
	go mod tidy

vet:
	go vet ./...

install-static-check:
	go install honnef.co/go/tools/cmd/staticcheck@latest

static-check:
	staticcheck ./...

install-go-lint:
	go install golang.org/x/lint/golint@latest

lint:
	golint ./...

build:
	go build ./...

test:
	go test --race ./...