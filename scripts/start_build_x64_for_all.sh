#!/bin/bash
set -o errtrace
os=(windows linux linux_musl)

errorHandler() {
  echo ERROR: ${BASH_COMMAND} failed with error code $?
  exit 1
}

trap  errorHandler ERR

version="$1.$2"
cp assets/icons/file.png internal/handlers/filehandler/

echo "Creating dirs"
for dir in ${os[@]} 
do
  if [ ! -d $dir ]; then
    mkdir -p $dir || $("FATAL: Failed to create $dir" && exit 100)
  fi
  echo "====> build for $dir <===="

  bash ./scripts/build_x64_for_"$dir".sh $(pwd)/$dir $version &
done

wait
echo "====> all jobs done <===="