#!/bin/bash
set -Eeuo pipefail
cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/../.."

export CGO_CPPFLAGS="${CPPFLAGS:-}"
export CGO_CFLAGS="${CFLAGS:-}"
export CGO_CXXFLAGS="${CXXFLAGS:-}"
export CGO_LDFLAGS="${LDFLAGS:-}"

VERSION_PKG="github.com/Jamesits/serial/internal/cmd/version"

GIT_COMMIT="UNKNOWN"
CURRENT_TIME="$(date -u "+%Y-%m-%d %T UTC")"
COMPILE_HOST="$(cat /proc/sys/kernel/hostname)"
GIT_STATUS="unknown"
if command -v git > /dev/null; then
  GIT_COMMIT="$(git rev-list -1 HEAD | cut -c -8)"
  if output="$(git status --porcelain)" && [ -z "$output" ]; then
    GIT_STATUS="clean"
  else
    GIT_STATUS="dirty"
  fi
fi

mkdir -p build
! mkdir -p "$GOPATH"

export GO111MODULE=on
>&2 echo "[*] Downloading dependencies..."
go mod download
>&2 echo "[*] Verifying dependencies..."
go mod verify

# build
>&2 echo "[*] Building..."
go build -mod=readonly -modcacherw -trimpath -ldflags "-s -w -X \"${VERSION_PKG}.versionGitCommitHash=$GIT_COMMIT\" -X \"${VERSION_PKG}.versionCompileTime=$CURRENT_TIME\" -X \"${VERSION_PKG}.versionCompileHost=$COMPILE_HOST\" -X \"${VERSION_PKG}.versionGitStatus=$GIT_STATUS\"" -o "build/" ./cmd/...

# strip
if command -v strip; then
  >&2 echo "[-] Stripping..."
  strip --strip-unneeded "build/"*
else
  >&2 echo "[-] Stripping skipped."
fi

# upx
if command -v upx; then
	upx "build/"* && >&2 echo "[+] UPX succeeded." || >&2 echo "[-] UPX failed."
else
	>&2 echo "[-] UPX skipped."
fi

>&2 echo "[+] Done."
# set exit code even if the previous command fails
exit 0
