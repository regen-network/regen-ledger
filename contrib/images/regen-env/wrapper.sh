#!/usr/bin/env sh

set -euo pipefail
set -x

BINARY=/regen/${BINARY:-regen}
ID=${ID:-0}
LOG=${LOG:-regen.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'regen'"
	exit 1
fi


export REGENHOME="/data/node${ID}/regen"

if [ -d "$(dirname "${REGENHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${REGENHOME}" "$@" | tee "${REGENHOME}/${LOG}"
else
  "${BINARY}" --home "${REGENHOME}" "$@"
fi
