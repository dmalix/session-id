default:
	@echo "Makefile: Please specify a target to make."
	@echo "------------------------------------------"
	@grep '^[a-zA-Z0-9_-]*:' Makefile | grep -v '^default:'

check:
	go fmt ./...
	go vet ./...
	go test
	go test -bench=.

dependencies-init:
	rm --dir --recursive --force vendor
	rm --force go.mod
	rm --force go.sum
	go mod init github.com/dmalix/session-id
	go mod tidy
	go mod vendor

dependencies-update:
	go mod tidy
	go mod vendor
