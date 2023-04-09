#!/usr/bin/env bash
#
# This script builds the application from source for multiple platforms.
set -x
set +e
# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
LD_FLAGS="-s -w"
GOPATH=${GOPATH:-$(go env GOPATH)}

# Instruct gox to build statically linked binaries
#export CGO_ENABLED=0
#CGO_ENABLED=0
export GOFLAGS="-mod=readonly"

while [ -h "$SOURCE" ] ; 
do SOURCE="$(readlink "$SOURCE")"; 
done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
echo $DIR
cd "$DIR"
# Delete the old dir
echo "==> Recreate old directory..."
rm -f bin/*
mkdir -p bin/
mkdir -p pkg/
go mod download
BIN=$DIR/bin/
# Build!
cd $SOURCE
echo "==> Building..."
for APP_DIR in $(ls -d $DIR/cmd/*/); do 
   APP_NAME=$(basename $APP_DIR)
   cd $APP_DIR
   go build -o ${BIN}${APP_NAME} main.go
echo ${i%%/}; done
# Move all the compiled things to the $GOPATH/bin

# Done!
echo
echo "==> Results:"
ls $BIN