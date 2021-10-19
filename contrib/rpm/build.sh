#!/bin/bash
set -Eeuo pipefail
set -x
shopt -s globstar
cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/../.."

rpmdev-setuptree

# archive source
# git archive --format=tar.gz --prefix=serial-0.0.0/ -o "${HOME}/rpmbuild/SOURCES/serial-0.0.0.tar.gz" HEAD
tar -cvf "${HOME}/rpmbuild/SOURCES/serial-0.0.0.tar.gz" --transform "s/^\./serial-0.0.0/" .

rpmbuild -ba contrib/rpm/serial.spec
