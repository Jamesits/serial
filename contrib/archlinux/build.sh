#!/bin/bash
set -Eeuo pipefail
cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/../.."

BUILD_DIR="/tmp/archlinux"

# make source tarball
>&2 echo "[*] Archiving source..."
mkdir -p "${BUILD_DIR}"
tar -cvJf "${BUILD_DIR}/source.tar.xz" .
cp -rv contrib/archlinux/* "${BUILD_DIR}"

# build
>&2 echo "[*] Building..."
pushd "${BUILD_DIR}"
updpkgsums
makepkg
popd

>&2 echo "[+] Done."
