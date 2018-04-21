PKG := github.com/akshayjshah/negotiate

help: ## Show rules and documentation
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

bin/golint: Gopkg.lock
	go build -o $@ ./vendor/github.com/golang/lint/golint

bin/goimports: Gopkg.lock
	go build -o $@ ./vendor/golang.org/x/tools/cmd/goimports

bin/megacheck: Gopkg.lock
	go build -o $@ ./vendor/honnef.co/go/tools/cmd/megacheck

fmt: bin/goimports ## Reformat this package with goimports
	./bin/goimports -w -local $(PKG) .

.PHONY: lint
lint: bin/goimports bin/golint bin/megacheck ## Run all linters
ifdef SKIP_LINT
	@echo "Skipping linters on" $(shell go version)
else
	@rm -rf lint.log
	@echo "Checking formatting..."
	@bin/goimports -d -local $(PKG) . 2>&1 | tee lint.log
	@echo "Checking vet..."
	@go vet . 2>&1 | tee -a lint.log
	@echo "Checking lint..."
	@golint . 2>&1 | tee -a lint.log
	@echo "Running megacheck..."
	@bin/megacheck . 2>&1 | tee -a lint.log
	@echo "Checking for unresolved FIXMEs..."
	@git grep -i fixme | grep -v -e vendor -e Makefile | tee -a lint.log
	@[ ! -s lint.log ]
endif

.PHONY: test
test: ## Run unit tests
	go test -v -race -cover -coverprofile=cover.out .
