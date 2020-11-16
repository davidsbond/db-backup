.DEFAULT_GOAL := build

# Downloads and vendors dependencies
modules:
	go mod tidy
	go mod vendor

# Formats all go source code
format:
	grep -L -R "Code generated .* DO NOT EDIT" --exclude-dir=.git --exclude-dir=vendor --include="*.go" | \
	xargs -n 1 gofumports -w -local github.com/davidsbond/homelab

# Runs go tests
test:
	go test -race ./...

# Installs go tooling
install-tools:
	go install \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		mvdan.cc/gofumpt/gofumports

# Lints go source code
lint:
	golangci-lint run --enable-all

# Compiles go source code
build:
	./scripts/build.sh

docker:
	./scripts/docker.sh

has-changes:
	git add .
	git diff --staged --exit-code
