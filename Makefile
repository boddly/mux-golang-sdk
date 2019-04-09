lint:
	@echo "Running ${@}"
	@golangci-lint --disable=megacheck --disable=typecheck run

test:
	@echo "Running ${@}"
	@go clean ./...
	@go test ./...

integration:
	@echo "Running ${@}"
	@go test ./... -tags=integration