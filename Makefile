.PHONY: compile-proto
compile-proto:
	@echo "==> Checking buf dependencies..."
ifeq (, $(shell command -v buf 2> /dev/null))
	@echo "==> Setup: Buf not installed, please follow the instructions on https://docs.buf.build/installation"
endif
	@echo "==> compiling proto..."
	@echo "===> generating grpc code..."
	@go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@buf generate
