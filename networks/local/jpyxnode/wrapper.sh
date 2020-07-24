#!/usr/bin/env sh

BINARY=/jpyxd/linux/${BINARY:-jpyxd}
echo "binary: ${BINARY}"
ID=${ID:-0}
LOG=${LOG:-jpyxd.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'jpyxd' E.g.: -e BINARY=jpyxd_my_test_version"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export JPYXDHOME="/jpyxd/node${ID}/jpyxd"

if [ -d "$(dirname "${JPYXDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${JPYXDHOME}" "$@" | tee "${JPYXDHOME}/${LOG}"
else
  "${BINARY}" --home "${JPYXDHOME}" "$@"
fi
