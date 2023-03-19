#!/usr/bin/env bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
BDIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"



pushd ${BDIR}/cmd/mint >> /dev/null

go build  -o "${BDIR}/bin/mint"

popd >> /dev/null