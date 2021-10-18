#!/bin/bash
set -Eeuo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/../.."

# set GOROOT to the actual goroot, else you will have strange errors complaining cannot load bufio
# fix GOPATH if it doesn't exist
export GOPATH="${GOPATH:-/tmp/gopath}"
OUT_FILE="${OUT_FILE:-serial}"

export CGO_CPPFLAGS="${CPPFLAGS:-}"
export CGO_CFLAGS="${CFLAGS:-}"
export CGO_CXXFLAGS="${CXXFLAGS:-}"
export CGO_LDFLAGS="${LDFLAGS:-}"

GIT_COMMIT="$(git rev-list -1 HEAD | cut -c -8)"
CURRENT_TIME="$(date -u "+%Y-%m-%d %T UTC")"
COMPILE_HOST="$(cat /proc/sys/kernel/hostname)"
GIT_STATUS=""
if output="$(git status --porcelain)" && [ -z "$output" ]; then
	GIT_STATUS="clean"
else
	GIT_STATUS="dirty"
fi

mkdir -p build
! mkdir -p "$GOPATH"

# go get -d ./...
export GO111MODULE=on
go mod download -x
go mod verify

# build
go build -mod=readonly -modcacherw -trimpath -ldflags "-s -w -X \"main.versionGitCommitHash=$GIT_COMMIT\" -X \"main.versionCompileTime=$CURRENT_TIME\" -X \"main.versionCompileHost=$COMPILE_HOST\" -X \"main.versionGitStatus=$GIT_STATUS\"" -o "build/" ./cmd/...
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
