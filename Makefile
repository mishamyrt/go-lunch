# Library version
VERSION = v0.0.1

# Tool versions
GOLANGCI_LINT_VERSION = v1.55.2
REVIVE_VERSION = v1.3.4
GIT_CHGLOG_VERSION = v0.15.4

# Internal constants
GO_BIN_PATH := $(shell go env GOPATH)/bin

.PHONY: setup
setup:
	curl -sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(GO_BIN_PATH) $(GOLANGCI_LINT_VERSION)
	go install github.com/mgechev/revive@$(REVIVE_VERSION)
	go install github.com/git-chglog/git-chglog/cmd/git-chglog@$(GIT_CHGLOG_VERSION)

.PHONY: lint
lint:
	golangci-lint run
	revive -config revive.toml  ./...

.PHONY: test
test:
	go test -cover ./...

.PHONY: changelog
changelog:
# Create tag for changelog generator
	git tag "$(VERSION)"
	git-chglog -o CHANGELOG.md
	git commit -am 'chore: release $(VERSION)'
# Recreate tag after commit
	git tag -D "$(VERSION)"
	git tag "$(VERSION)"
	