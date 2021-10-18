#!/bin/bash
set -Eeuo pipefail
set -x
shopt -s globstar
cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/../.."

declare -a GO_MANUAL_DEPS=(
  "github.com/jszwec/csvutil"
  "github.com/creack/goselect"
  "go.bug.st/serial"
)

#if [ ! -d "/var/cache/pbuilder/base.cow" ]; then
#  sudo cowbuilder --create
#fi

mkdir -p /tmp/artifacts

pushd /tmp
for pkg in "${GO_MANUAL_DEPS[@]}"; do
  # setup build directory
  rm -rf build || true
  mkdir -p build
  pushd build

  # generate debian control files
  dh-make-golang -allow_unknown_hoster -type "library" "$pkg"
  pushd */
  # git add .
  # git commit -m "add debian packaging metadata"

  # install dependencies
  mk-build-deps --root-cmd sudo --install --tool "apt-get -o Debug::pkgProblemResolver=yes --no-install-recommends -y"

  # build
  # gbp buildpackage --git-pbuilder
  dpkg-buildpackage --build=binary --no-sign

  # collect artifacts
  find .. -type f -maxdepth 1 -exec cp -afv -- \{\} /tmp/artifacts \;

  # install artifacts
  sudo dpkg -i ../*.deb
  sudo apt-get install -fy || true
  popd
  popd
done
popd

# install artifacts
# sudo dpkg -i /tmp/artifacts/*.deb

# generate debian control files
pushd /tmp
dh-make-golang github.com/Jamesits/serial
popd

# overlay files
cp -afv /tmp/serial/debian .
cp -afv contrib/debian/overrides/* debian/

# install dependencies
mk-build-deps --root-cmd sudo --install --tool "apt-get -o Debug::pkgProblemResolver=yes --no-install-recommends -y"

# build
dpkg-buildpackage --build=binary --no-sign

# collect artifacts
find .. -type f -maxdepth 1 -exec cp -afv -- \{\} /tmp/artifacts \;
