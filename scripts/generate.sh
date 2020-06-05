#!/bin/bash
set -ex

OUTPUT_DIR="$1"

for file in $(find . -name '*.go') ; do
    nextdir="$(dirname $file)"
    mkdir -p "$OUTPUT_DIR/$nextdir"
    cp -v "$file" "$OUTPUT_DIR/$nextdir/"
done
    
