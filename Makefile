.PHONY: lint
lint:
	$(info Run go linters in project...)
	golangci-lint run ./... -c ./.golangci.yml

.PHONY: build
build:
	$(info Build project...)
	go build -o bin/app ./cmd/note/main.go