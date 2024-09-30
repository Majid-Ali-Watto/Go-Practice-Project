#!/bin/bash

# Extract the current version from version.go
current_version=$(grep '^var Version' version.go | sed -E 's/var Version = "(.*)"/\1/')
echo "Current version: $current_version"

# Split version into major, minor, patch
IFS='.' read -r major minor patch <<< "${current_version#v}"

# Increment the patch version
patch=$((patch + 1))

# If patch exceeds 9, reset patch to 0 and increment minor
if [ "$patch" -ge 10 ]; then
  patch=0
  minor=$((minor + 1))
fi

# If minor exceeds 9, reset minor to 0 and increment major
if [ "$minor" -ge 10 ]; then
  minor=0
  major=$((major + 1))
fi

# Create new version string
new_version="$major.$minor.$patch"
echo "New version: $new_version"

# Update version.go with the new version
sed -i "s/^var Version = \".*\"/var Version = \"$new_version\"/" version.go
echo "Updated version.go with new version $new_version"

# Add updated version.go to staging
git add version.go

# Append new version to commit message
# Get the latest commit message and append the new version
commit_message=$(git log -1 --pretty=%B)
git commit --amend -m "$commit_message - Bump version to $new_version"
