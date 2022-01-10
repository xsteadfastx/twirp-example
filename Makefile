BUF := go run github.com/bufbuild/buf/cmd/buf
GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: clean
clean:
	rm -rf api/go
	rm -rf api/swagger

.PHONY: generate
generate: install-tools
	@echo "\n==== GOLANG ===="
	go generate

	@echo "\n==== PROTO ===="
	cd api/proto; $(BUF) -v generate
	@# find api/swagger -name "*.swagger.json" -exec file {} \;

.PHONY: lint
lint:
	@echo "\n==== GOLANG ===="
	$(GOLANGCI_LINT) run \
		--enable-all \
		--disable=exhaustivestruct,godox,varnamelen \
		--build-tags=integration

	@echo "\n==== PROTO ===="
	cd api/proto; $(BUF) lint

.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor

.PHONY: install-tools
install-tools:
	@echo "\n==== INSTALL TOOLS ===="
	go list -f '{{range .Imports}}{{.}} {{end}}' third_party/tools/tools.go | xargs go install -v
