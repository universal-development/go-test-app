GO_FILES?=$$(find . -name '*.go' | grep -v vendor)

.PHONY: build

build:
	@echo "Building app"
	go build

fmt:
	@echo "Formatting files"
	gofmt -w $(GO_FILES)
	goimports -w $(GO_FILES)

pre-commit-hook:
	@ln -s scripts/pre-commit .git/hooks/pre-commit
	@echo "hook installed."

