.PHONY: setup test test-py test-go fmt help

help:
	@echo "make setup    - Install Python dependencies"
	@echo "make test     - Run Python and Go tests"
	@echo "make test-py  - Run pytest"
	@echo "make test-go  - Run go test"
	@echo "make fmt      - Format Go code"

setup:
	uv sync

test: test-py test-go

test-py:
	uv run pytest

test-go:
	go test ./...

fmt:
	go fmt ./...
