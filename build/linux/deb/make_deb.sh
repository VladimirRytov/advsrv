#!/bin/env bash

errorHandler() {
  echo ERROR: ${BASH_COMMAND} failed with error code $?
  exit 1
}

trap errorHandler ERR

apt update && apt install build-essential md5deep -y
mkdir -p advsrv-$1-$2/DEBIAN advsrv-$1-$2/{etc/advsrv,usr/{sbin,share/doc/advsrv,lib/systemd/system}}

install -m 0750 linux/advsrv advsrv-$1-$2/usr/sbin/
cp build/linux/advsrv.service advsrv-$1-$2/usr/lib/systemd/system

deb_dir=$(pwd)/advsrv-$1-$2/
md5deep -r advsrv-$1-$2/usr > advsrv-$1-$2/DEBIAN/md5sums
sed "s|$deb_dir||g" -i advsrv-$1-$2/DEBIAN/md5sums
cp build/linux/deb/{control,postinst,postremove,preremove} advsrv-$1-$2/DEBIAN
cp build/linux/deb/copyright advsrv-$1-$2/usr/share/doc/advsrv
total_size=$(du -sk advsrv-$1-$2/usr | awk '{ print $1 }')
sed -e 's/${version}/'"$1"'/' -e 's/${release}/'"$2"'/' -e 's/${buildNumber}/'"$3"'/' -e 's/${size}/'"$total_size"'/' -i advsrv-$1-$2/DEBIAN/control

chmod 0775 advsrv-$1-$2/DEBIAN/{postremove,postinst,preremove}
fakeroot dpkg-deb --build advsrv-$1-$2