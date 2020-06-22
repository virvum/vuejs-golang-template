COLOR_PIPE = 2>&1 | sed "s/.*/$$(printf '\033[31m&\033[0m')/"
GO_DIRS = . backend utils

.PHONY: build
build:
	jq ".version = \"$$(git describe 2>/dev/null || { echo -n '0.0.1-'; git rev-parse --short HEAD; })\"" < frontend/package.json > frontend/package.json.new
	mv frontend/package.json.new frontend/package.json
	npm run build | cat
	cd backend; GOOS=linux GOARCH=amd64 make build
	cd backend; GOOS=darwin GOARCH=amd64 make build
	cd backend; GOOS=windows GOARCH=amd64 make build

.PHONY: readme
readme:
	go run utils/generate-readme.go

.PHONY: npmlint
npmlint:
	npm run lint

.PHONY: godoc
godoc:
	GOROOT=$(pwd) GOPATH=$(pwd) godoc -http=:6060

.PHONY: gofmt
gofmt:
	@echo gofmt
	@find $(GO_DIRS) -maxdepth 1 -type f -name '*.go' | xargs gofmt -s -e -d -w $(COLOR_PIPE)

.PHONY: govet
govet:
	@echo govet
	@find $(GO_DIRS) -maxdepth 1 -type f -name '*.go' | xargs dirname | sort -u | xargs go vet $(COLOR_PIPE)

.PHONY: golint
golint:
	@echo golint
	@golint ./... $(COLOR_PIPE)

.PHONY: gocyclo
gocyclo:
	@echo gocyclo
	@find $(GO_DIRS) -maxdepth 1 -type f -name '*.go' | xargs gocyclo -over 15 $(COLOR_PIPE)

.PHONY: gotest
test:
	@echo go test
	@for d in $(GO_DIRS); do pushd $$d >/dev/null; go test $(COLOR_PIPE); popd >/dev/null; done

.PHONY: gomodtidy
gomodtidy:
	@echo go mod tidy
	@for d in $(GO_DIRS); do \
		pushd $$d >/dev/null; \
		test -f go.mod && go mod tidy $(COLOR_PIPE); \
		popd >/dev/null; \
	done
