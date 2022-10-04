fmt:
	gofmt -l -e .

test: fmt
	go test ./...

checks:
	staticcheck
	golangci-lint run ./...
