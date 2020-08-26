# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test -count=1
GOMOD=$(GOCMD) mod
GOGET=$(GOCMD) get
GOUPGRADE=$(GOCMD) get -u
GOFORMAT=go list -f {{.Dir}} ./... | xargs gofmt -s -w -d

.PHONY: test
test: 
	@$(GOTEST) -timeout 3s -v ./...

.PHONY: fmt
fmt:
	@echo tidying and formatting go code
	@$(GOMOD) tidy
	@$(GOFORMAT)
