#!/bin/bash
set -Eeuo pipefail
shopt -s globstar
cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/../.."

>&2 echo "[*] Setting up RPM environment..."
rpmdev-setuptree

# archive source
>&2 echo "[*] Archiving source..."
# git archive --format=tar.gz --prefix=serial-0.0.0/ -o "${HOME}/rpmbuild/SOURCES/serial-0.0.0.tar.gz" HEAD
tar -cf "${HOME}/rpmbuild/SOURCES/serial-0.0.0.tar.gz" --transform "s/^\./serial-0.0.0/" .

>&2 echo "[*] Building RPM..."
rpmbuild -ba contrib/rpm/serial.spec

>&2 echo "[+] Done."
