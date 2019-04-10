lint:
	@echo "Running ${@}"
	@golangci-lint --disable=megacheck --disable=typecheck run

test:
	@echo "Running ${@}"
	@go clean ./...
	@go test ./...
	@go test -v -covermode=count -coverprofile=coverage.out

integration:
	@echo "Running ${@}"
	@go test ./... -tags=integration