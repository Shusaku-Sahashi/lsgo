.DEFAULT_GOAL := alias

.PHONY: run
run: # Run code
	go run ./...

.PHONY: build
build: # Build code
	go build ./...

.PHONY: alias
alias: # show all command
	@grep -E '^[a-zA-Z0-9]+\s*:.*#' ${MAKEFILE_LIST} | sed -E 's/^([a-zA-Z0-9_-]+):.*#(.*)/\1\t\2/'
