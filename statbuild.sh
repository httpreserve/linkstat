#!/usr/bin/env bash
set -eux

STAT="linkstat"
DIR="release"
rm -rf "$DIR"
mkdir "$DIR"
env GOOS=windows GOARCH=386 go build
mv "$STAT".exe "${DIR}/${STAT}"-win386.exe
env GOOS=windows GOARCH=amd64 go build
mv "$STAT".exe "${DIR}/${STAT}"-win64.exe
env GOOS=linux GOARCH=amd64 go build
mv "$STAT" "${DIR}/${STAT}"-linux64
env GOOS=darwin GOARCH=386 go build
mv "$STAT" "${DIR}/${STAT}"-darwin386
env GOOS=darwin GOARCH=amd64 go build
mv "$STAT" "${DIR}/${STAT}"-darwinAmd64
