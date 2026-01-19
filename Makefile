.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build-lib
build-lib:
	CGO_ENABLED=1 go build -o libspliit.so -buildmode=c-shared ./cbindings
