#!/bin/env bash

apt update && apt install curl file -y
curl -L  https://github.com/AppImage/AppImageKit/releases/download/continuous/appimagetool-x86_64.AppImage -o appImage 
chmod +x appImage
./appImage --appimage-extract
mkdir -p advsrv.AppDir/usr/{sbin,lib,share/applications}
install -m 0750 linux/advsrv  advsrv.AppDir/usr/sbin

cp squashfs-root/usr/bin/AppRun advsrv.AppDir/
cp assets/icons/advertising256.png advsrv.AppDir/advsrv.png
install -m 0750 build/linux/advsrv.desktop advsrv.AppDir/

squashfs-root/AppRun advsrv.AppDir
err=$?
if [ $err -ne 0 ];
then 
  echo "An error occured while creating AppImage package."
  exit $err
fi