#!/usr/bin/env bash

DIR=$(pwd)
BIN_DIR=${DIR}/bin

# Application metadata
APP_VERSION=$(git describe --tags --always)

rm -rf "${BIN_DIR}"
mkdir -p "${BIN_DIR}"

# Iterates over each subdirectory in ~/cmd that contains a main.go file
# and builds the binary. Each binary is placed within ~/bin.
APP_NAME=$(basename "${DIR}")
APP_DESCRIPTION=$(cat ./DESCRIPTION)
COMPILED=$(date +%s)

OUTPUT=${BIN_DIR}/${APP_NAME}
CGO_ENABLED=0 go build -ldflags \
  "-w -s "\
"-X 'pkg.dsb.dev/environment.Version=${APP_VERSION}'"\
"-X 'pkg.dsb.dev/environment.compiled=${COMPILED}'"\
"-X 'pkg.dsb.dev/environment.ApplicationName=${APP_NAME}'"\
"-X 'pkg.dsb.dev/environment.ApplicationDescription=${APP_DESCRIPTION}'" \
  -o "${OUTPUT}" "${DIR}"
