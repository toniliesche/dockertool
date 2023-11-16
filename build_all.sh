#!/bin/bash

ARCHS="amd64 arm64"
OSS="darwin linux windows"

if [ $# -gt 0 ]; then
  PROFILE=$1
else
  PROFILE=generic
fi

for ARCH in ${ARCHS}; do
  for OS in ${OSS}; do
    make build GOOS=${OS} GOARCH=${ARCH} PROFILE=${PROFILE}
  done
done