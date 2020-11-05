.PHONY: vendor proto-gen protoc-gen-gocosmos test

protoc-gen-gocosmos:
	@echo "Installing protoc-gen-gocosmos..."
	@go install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos

vendor:
	go mod vendor

proto-gen: vendor
	./protocgen.sh

test:
	@go test -mod=readonly -race  ./...