#!/bin/bash
set -Eeuo pipefail
set -x
shopt -s globstar
cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/../.."

declare -a GO_MANUAL_DEPS=(
  "github.com/jszwec/csvutil"
  "go.bug.st/serial"
)

if [ ! -d "/var/cache/pbuilder/base.cow" ]; then
  cowbuilder --create
fi

pushd /tmp
for pkg in "${GO_MANUAL_DEPS[@]}"; do
  rm -rf build || true
  mkdir -p build
  pushd build
  dh-make-golang "$pkg"
  ls -alh . ..
  pushd */
  git add .
  git commit -m "add debian packaging metadata"
  gbp buildpackage --git-pbuilder
  ls -alh . ..
  popd
  popd
done
popd

pushd /tmp
dh-make-golang github.com/Jamesits/serial
popd

cp -afv /tmp/serial/debian .
cp -afv contrib/debian/overrides/* .

mk-build-deps --root-cmd sudo --install --tool "apt-get -o Debug::pkgProblemResolver=yes --no-install-recommends -y"
dpkg-buildpackage --build=binary --no-sign
