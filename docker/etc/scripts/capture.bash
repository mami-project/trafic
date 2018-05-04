#!/bin/bash

set -eu

function mklabel() {
	local exid=$1
	local unixtime=$(date +%s)

	printf "lola-%s-%s" "${exid}" "${unixtime}"
}

EXID=${EXID:-flows}
IFACE=${IFACE:-eth0}

label=$(mklabel "$EXID")
outfn="${label}.pcap"

echo ">> now run: CONF=${EXID}.env LABEL=${label} ./run-clients.bash"
tshark -i ${IFACE} -s 128 -w ${outfn} -a duration:60
