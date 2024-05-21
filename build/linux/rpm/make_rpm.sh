#!/bin/env bash

errorHandler() {
  echo ERROR: ${BASH_COMMAND} failed with error code $?
  exit 1
}

trap errorHandler ERR

dnf install -y rpmdevtools rpmlint
HOME=$(pwd)
rpmdev-setuptree

cp linux/advsrv ~/rpmbuild/BUILD
cp build/linux/rpm/advsrv.spec ~/rpmbuild/SPECS/
sed -e 's/${version}/'"$1"'/' -e 's/${release}/'"$2"'/' -e 's/${buildNumber}/'"$3"'/' -i ~/rpmbuild/SPECS/advsrv.spec
rpmlint ~/rpmbuild/SPECS/advsrv.spec
rpmbuild -bb ~/rpmbuild/SPECS/advsrv.spec