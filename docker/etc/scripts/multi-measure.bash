#!/bin/bash

set -eum

function main() {
	local i, max

	cwd=$(dirname $0)
	i=0
	max=${1:-10}
	while [ $i -lt $max ]; do
		i=$((i + 1))
		IFACE=eth0 ${cwd}/measure.bash ${i}
	done
}

main $*
