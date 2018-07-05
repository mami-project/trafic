#!/bin/bash
set -eum

#
# help: multi-measure.bash [iterations [gap]]
#   default:  
#      iterations=10
#      gap=10

function main() {
	local i, max, sleep

	cwd=$(dirname $0)
	i=0
	max=${1:-10}
	sleep=${2:-10}
	while [ $i -lt $max ]; do
		i=$((i + 1))
		IFACE=eth0 ${cwd}/measure.bash ${i}
		sleep ${sleep}
	done
}

main $*
