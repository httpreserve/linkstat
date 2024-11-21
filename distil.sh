#!/usr/bin/env bash
set -eux

MOONSHINE="linkstat"
DIR="release"
VERSION="${MOONSHINE}"
rm -rf "$DIR"
mkdir -p "$DIR"
export GOOS=windows
export GOARCH=386
go build
mv "$MOONSHINE".exe "${DIR}/${VERSION}"-win386.exe
export GOOS=windows
export GOARCH=amd64
go build
mv "$MOONSHINE".exe "${DIR}/${VERSION}"-win64.exe
export GOOS=linux
export GOARCH=amd64
go build
mv "$MOONSHINE" "${DIR}/${VERSION}"-linux64
export GOOS=
export GOARCH=386
go build
mv "$MOONSHINE" "${DIR}/${VERSION}"-darwin386
export GOOS=darwin
export GOARCH=amd64
go build
mv "$MOONSHINE" "${DIR}/${VERSION}"-darwinAmd64
export GOOS=
export GOARCH=
