#!/bin/sh

set -eu

usage() {
    echo "Syntax: ./scripts/set-version.sh 2.0.0"
    exit 1
}

# Check if the version argument was passed.
if [ "$#" -ne 1 ]; then
    usage
fi

# Check if the version file exists.
if [ ! -f version/version.go ]; then
    usage
fi

# Replace the version.
sed -i "s/VERSION = .*$/VERSION = \"$1\"/g" "version/version.go"