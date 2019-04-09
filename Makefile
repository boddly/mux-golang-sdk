lint: ## Analyze and find programs in source code
	@echo "Running ${@}"
	@golangci-lint run