#!/usr/bin/env sh

BINARY=/regen/${BINARY:-regen}
ID=${ID:-0}
LOG=${LOG:-regen.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'regen'"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export REGENHOME="/regen/node${ID}/regen"

if [ -d "$(dirname "${REGENHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${REGENHOME}" "$@" | tee "${REGENHOME}/${LOG}"
else
  "${BINARY}" --home "${REGENHOME}" "$@"
fi
