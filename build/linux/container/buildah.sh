#!/bin/env bash

errorHandler() {
  echo ERROR: ${BASH_COMMAND} failed with error code $?
  exit 1
}

trap errorHandler ERR

container=$(buildah from alpine)
containermount=$(buildah mount $container) 
cp linux_musl/advsrv $containermount/usr/sbin
buildah unmount $container
buildah config --port 4567 $container
buildah config --cmd '/usr/sbin/advsrv start'  $container
buildah commit $container advsrv:$1
buildah push advsrv:$1 docker-archive:advsrv.tar
gzip advsrv.tar