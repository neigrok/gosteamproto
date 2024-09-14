PROTO_BASE_DIR=raw
OUT_BASE_DIR=compiled

.PHONY: proto

proto:
	@echo "Compiling protobuf files..."
	./compile.sh --proto-base-dir=$(PROTO_BASE_DIR) --out-base-dir=$(OUT_BASE_DIR)
	@echo "Protobuf files compiled successfully"
