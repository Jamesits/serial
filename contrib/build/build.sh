#!/bin/bash
set -Eeuo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/../.."

# set GOROOT to the actual goroot, else you will have strange errors complaining cannot load bufio
# fix GOPATH if it doesn't exist
export GOPATH=${GOPATH:-/tmp/gopath}
OUT_FILE=${OUT_FILE:-serial}

GIT_COMMIT=$(git rev-list -1 HEAD | cut -c -8)
CURRENT_TIME=$(date -u "+%Y-%m-%d %T UTC")
COMPILE_HOST=$(hostname --fqdn)
GIT_STATUS=""
if output=$(git status --porcelain) && [ -z "$output" ]; then
	GIT_STATUS="clean"
else
	GIT_STATUS="dirty"
fi

mkdir -p build
! mkdir -p "$GOPATH"

# go get -d ./...
export GO111MODULE=on
go mod download
go mod verify

# build
find cmd/ -mindepth 1 -maxdepth 1 -type d -print0 | while IFS= read -r -d '' line; do
  echo "Compiling $line..."
  go build -ldflags "-s -w -X \"main.versionGitCommitHash=$GIT_COMMIT\" -X \"main.versionCompileTime=$CURRENT_TIME\" -X \"main.versionCompileHost=$COMPILE_HOST\" -X \"main.versionGitStatus=$GIT_STATUS\"" -o "build/" "github.com/Jamesits/serial/$line"
done

ls -alh build/

# upx
if command -v upx; then
	! upx "build/"*
else
	echo "UPX not installed, compression skipped"
fi

ls -alh build/

# set exit code even if the previous command fails
exit 0
