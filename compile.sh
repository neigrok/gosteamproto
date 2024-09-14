#!/bin/bash

set -e

# Default values
PROTO_BASE_DIR=""
OUT_BASE_DIR=""

# Find protoc include directory
PROTOC_INCLUDE=$(dirname $(which protoc))/../include

# Parse named arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --proto-base-dir=*)
            PROTO_BASE_DIR="${1#*=}"
            shift
            ;;
        --out-base-dir=*)
            OUT_BASE_DIR="${1#*=}"
            shift
            ;;
        *)
            echo "Unknown parameter: $1"
            exit 1
            ;;
    esac
done

# Validate required arguments
if [ -z "$PROTO_BASE_DIR" ] || [ -z "$OUT_BASE_DIR" ]; then
    echo "Error: Both --proto-base-dir and --out-base-dir are required"
    exit 1
fi

# Process each subdirectory in the proto base directory
for proto_dir in "$PROTO_BASE_DIR"/*/ ; do
    if [ ! -d "$proto_dir" ] || [[ "$proto_dir" == *"/google/"* ]]; then
        continue
    fi

    # Get the relative package name from the directory
    pkg_name=$(basename "$proto_dir")

    echo "Processing package: $pkg_name"

    # Create the output directory if it does not exist
    mkdir -p "$OUT_BASE_DIR"

    # Map every proto file in the directory to its Go package using M arguments
    MAPPING_ARGS=""
    for proto_file in $(find "$proto_dir" -name "*.proto"); do
        # Get the proto file name
        proto_filename=$(basename "$proto_file")

        # Create the M argument for each .proto file (mapping proto file to Go import path)
        MAPPING_ARGS+=" --go_opt=M${proto_filename}=./${pkg_name}"
    done

    # Compile all .proto files in this directory
    for proto_file in $(find "$proto_dir" -name "*.proto"); do
        echo "Processing $proto_file..."

        protoc --experimental_allow_proto3_optional \
               --proto_path="$proto_dir" \
               --proto_path="$PROTO_BASE_DIR" \
               --proto_path="$PROTOC_INCLUDE" \
               --go_out="$OUT_BASE_DIR" \
               ${MAPPING_ARGS} \
               "$proto_file"
    done
done

echo "Go structs generation completed."
