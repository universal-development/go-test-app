SHELL := /bin/bash

GO_FILES?=$$(find . -name '*.go' | grep -v vendor)

.PHONY: build

build:
	@echo "Building app"
	go build

container: build
	@echo "Building container"
	$(eval TAG:=$(shell git describe --tags --abbrev=12 --dirty --broken))
	docker build . -t "go-test-app:$(TAG)"
	docker tag "go-test-app:$(TAG)" "go-test-app:local"

fmt:
	@echo "Formatting files"
	gofmt -w $(GO_FILES)
	goimports -w $(GO_FILES)

pre-commit-hook:
	@ln -s scripts/pre-commit .git/hooks/pre-commit
	@echo "hook installed."

