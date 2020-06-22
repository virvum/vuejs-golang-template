APP_NAME = $(shell jq -r .name < ../frontend/package.json)
APP_VERSION = $(shell jq -r .version < ../frontend/package.json)
APP_HOMEPAGE = $(shell jq -r .homepage < ../frontend/package.json)

GOOS ?= linux
GOARCH ?= amd64
DIST_DIR = ../dist
BINARY_NAME = $(APP_NAME)-$(GOOS)-$(GOARCH)
LDFLAGS = -s -w \
	-X 'main.buildName=$(APP_NAME)' \
	-X 'main.buildVersion=$(APP_VERSION)' \
	-X 'main.buildHomePage=$(APP_HOMEPAGE)' \
	-X 'main.buildUserAgent=$(APP_NAME)/$(APP_VERSION) +$(APP_HOMEPAGE)' \
	-X 'main.buildDate=$$(date -u '+%Y-%m-%d %H:%M:%S %Z')'

.PHONY: run
run:
	go run -v -trimpath -ldflags "$(LDFLAGS)" -tags=run . -c ../config.yaml

.PHONY: build
build:
	GOOS= GOARCH= go run assets_generate.go
	CGO_ENABLED=0 GIN_MODE=release GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -trimpath -ldflags "$(LDFLAGS)" -tags=build -o "$(DIST_DIR)/$(BINARY_NAME)" .
