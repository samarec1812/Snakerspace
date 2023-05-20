.PHONY: lint
lint:
	$(info Run go linters in project...)
	golangci-lint run ./... -c ./.golangci.yml

.PHONY: build
build:
	$(info Build project...)
	go mod tidy
	go build -o bin/app ./cmd/article/main.go