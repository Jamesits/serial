#!/bin/bash
set -Eeuo pipefail
set -x
shopt -s globstar
cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/../.."

rpmdev-setuptree

# archive source
# git archive --format=tar.gz --prefix=serial-0.0.1/ -o "${HOME}/rpmbuild/SOURCES/serial-0.0.1.tar.gz" HEAD
tar -cvf "${HOME}/rpmbuild/SOURCES/serial-0.0.1.tar.gz" --transform "s/^\./serial-0.0.1/" .

rpmbuild -ba contrib/rpm/serial.spec
