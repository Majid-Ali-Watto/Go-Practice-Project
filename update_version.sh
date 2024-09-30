#!/bin/bash

# Read the current version from version.go
current_version=$(grep -oP 'var Version = "\K[0-9]+\.[0-9]+\.[0-9]+' version.go)

# Split the version into major, minor, and patch
IFS='.' read -r major minor patch <<< "$current_version"

# Increment the patch version if it's less than 9
if [ "$patch" -lt 9 ]; then
    patch=$((patch + 1))
else
    # If the patch version is 9, reset it and increment minor if less than 9
    patch=0
    if [ "$minor" -lt 9 ]; then
        minor=$((minor + 1))
    else
        # If the minor version is 9, reset it and increment major
        minor=0
        major=$((major + 1))
    fi
fi

# New version format
new_version="$major.$minor.$patch"

# Check if the version has changed
if [ "$new_version" != "$current_version" ]; then
    # Update the version.go file with the new version
    sed -i "s/^var Version = \".*\"/var Version = \"$new_version\"/" version.go
    echo "Updated version.go with version $new_version"
else
    echo "Version remains the same: $current_version"
fi
