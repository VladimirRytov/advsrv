#!/bin/bash

errorHandler() {
  echo ERROR: ${BASH_COMMAND} failed with error code $?
  exit 1
}

dnf install musl-devel musl-gcc golang git mingw64-gcc -y
git config --system --add safe.directory "*"
