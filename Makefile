# Parse aileron-gateway version.
AILERON := $(shell cat go.mod | grep github.com/aileron-gateway/aileron-gateway | awk '{ print $$1 "@" $$2 }')

# Build protoc command.
PROTO_CMD := protoc
PROTO_CMD += --proto_path ./proto
PROTO_CMD += --proto_path=$(shell go env GOMODCACHE)/$(AILERON)/proto/
PROTO_CMD += --plugin=protoc-gen-go=$(shell which protoc-gen-go)
PROTO_CMD += --go_out=./
PROTO_CMD += --go_opt=module="github.com/aileron-gateway/example-extension"

.PHONY: proto
proto:
	$(PROTO_CMD) $(shell find ./proto/ -type f -name "*.proto" -not -path "./proto/buf/*")
