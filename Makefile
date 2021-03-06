local_package=github.com/chyroc/tui

.PHONY: lint
.PHONY: test

run:
	make lint
	make test

lint:
	@$(set_env) go fmt ./...
	@$(set_env) go vet ./...
	@$(set_env) goimports -local $(local_package) -w .
	@$(set_env) go mod tidy

test:
	$(set_env) test -z "$$(gofmt -l .)"
	$(set_env) test -z "$$(goimports -local $(local_package) -d .)"
	./.github/test.sh

html:
	go tool cover -html=c.out
