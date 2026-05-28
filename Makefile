SHELL := /bin/zsh

GO ?= go
GOCACHE_DIR := /Users/platojobs/Desktop/Github/flutterffi/PFGolang/.gocache
GOMODCACHE_DIR := /Users/platojobs/Desktop/Github/flutterffi/PFGolang/.gomodcache
GO_RUN = GOCACHE=$(GOCACHE_DIR) GOMODCACHE=$(GOMODCACHE_DIR) $(GO)

.PHONY: help test bench foundations intermediate advanced modules lesson

help:
	@echo "Available targets:"
	@echo "  make test"
	@echo "  make bench"
	@echo "  make foundations"
	@echo "  make intermediate"
	@echo "  make advanced"
	@echo "  make modules"
	@echo "  make lesson LESSON=foundations/06_functions"

test:
	$(GO_RUN) test ./...

bench:
	$(GO_RUN) test -bench=. ./advanced/26_benchmarks

foundations:
	$(GO_RUN) run ./foundations/01_hello_world

intermediate:
	$(GO_RUN) run ./intermediate/13_pointers

advanced:
	$(GO_RUN) run ./advanced/54_deployment_ready_todo_service

modules:
	$(GO_RUN) run ./modules/02_flags_cli --name Learner

lesson:
	@if [ -z "$(LESSON)" ]; then echo "Usage: make lesson LESSON=foundations/06_functions"; exit 1; fi
	$(GO_RUN) run ./$(LESSON)
