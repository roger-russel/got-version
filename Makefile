.PHONY: test
test:
	@go test -cover -coverprofile=./coverage.txt -covermode=atomic -coverpkg=./... ./...

.PHONY: coverage
coverage: test
	@go tool cover -html=./coverage.txt -o coverage.html
	@google-chrome coverage.html

.PHONY: lint
lint:
	@golangci-lint run
