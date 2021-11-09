#!/bin/sh

set -eu

usage() {
    echo "Syntax: ./scripts/increment-version.sh"
    exit 1
}

# Check if the version file exists.
if [ ! -f internal/pkg/version/version.go ]; then
    usage
fi

# Get the current version.
version=$(cat internal/pkg/version/version.go | grep 'VERSION = ' | grep -o '".*"$' | tr -d '"')

# Strip the optional -SNAPSHOT part.
if echo "$version" | grep -q "\-SNAPSHOT"; then
    version=$(echo "$version" | sed "s/-SNAPSHOT//g")
fi

# Split the version.
major=$(echo "$version" | egrep -o "^[0-9]+")
minor=$(echo "$version" | egrep -o "\.[0-9]+\." | egrep -o "[0-9]+")
bugfix=$(echo "$version" | egrep -o "\.[0-9]+$" | egrep -o "[0-9]+")

# Increment the bugfix version.
bugfix=$((bugfix+1))

# Build the new version.
newversion="$major.$minor.$bugfix-SNAPSHOT"

# Set the new version.
./scripts/set-version.sh "$newversion"