PKG := ./...

.PHONY: all dep lint test race bench coverage clean

all: lint test race bench coverage clean

dep: ## Get the dependencies
	go get -v -d ${PKG}

lint: ## Lint the files
	go get -u golang.org/x/lint/golint
	golint -set_exit_status ${PKG}

test: dep ## Run unittests
	go test -v ${PKG}

race: dep ## Run data race detector
	go test -race -v ${PKG}

bench: dep ## Run benchmark
	go test -v ${PKG} -bench . -run ^$

coverage: dep ## Run code coverage
	go test -cover -coverprofile=coverage.out ${PKG}
	go tool cover -func=coverage.out

clean: ## Cleanup files created by make
	rm coverage.out

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
