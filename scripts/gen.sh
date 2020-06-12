#!/bin/bash
set -e

BASE_PACKAGE="$1"
SYMBOLS_PATH="pkg/internal/loader/symbols/"

_usage() {
	echo "USAGE: gen.sh PACKAGE [SKIP_PATTERN...]"
}

if [[ -z "$BASE_PACKAGE" ]] ; then
	_usage
	exit 1
fi

shift

_grep_args=("-v")

for arg in "$@" ; do
	_grep_args+=("-e" "$arg")
done

pushd "$SYMBOLS_PATH"
while IFS= read -r package
do
	echo "generating symbols for $package ..."
	goexports "$package"
done < <(go list "$BASE_PACKAGE" | grep "${_grep_args[@]}")
popd
