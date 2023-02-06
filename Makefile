PROTO_OUT = pkg/pb
PROTO_DIR = proto

DIR ?= $(shell pwd)

.PHONY: compile
compile-proto-go:
	@rm -rf $(PROTO_OUT)
	@mkdir -p $(PROTO_OUT)
	protoc --go_out=${PROTO_OUT} --go_opt=paths=source_relative \
		--go-grpc_out=${PROTO_OUT} --go-grpc_opt=paths=source_relative \
		${PROTO_DIR}/*.proto
