.PHONY: packages

packages:
	@dep ensure

.PHONY: coverage test
test:
	@go test -cover -coverprofile=./coverage.dirty.txt -covermode=atomic -coverpkg=all ./...
	@$(shell echo "mode: " > ./coverage.soup.txt)
	@$(shell go list ./... >> ./coverage.soup.txt)
	@$(shell cat ./coverage.dirty.txt | grep -f ./coverage.soup.txt > coverage.txt)
	@rm ./coverage.soup.txt ./coverage.dirty.txt

coverage: test
	@go tool cover -html=./coverage.txt -o coverage.html
	@google-chrome coverage.html
