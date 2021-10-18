#!/bin/bash
set -Eeuo pipefail
set -x
cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/../.."

BUILD_DIR="/tmp/archlinux"

# make source tarball
mkdir -p "${BUILD_DIR}"
tar -cvJf "${BUILD_DIR}/source.tar.xz" .
cp -rv contrib/archlinux/* "${BUILD_DIR}"

# build
pushd "${BUILD_DIR}"
updpkgsums
makepkg
popd
