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

echo ">> on the RAN side run:"
echo "  CONF=${EXID}.env LABEL=${label} ./run-clients.bash"
echo ">> press <ENTER> when ready to capture"
read
tshark -i ${IFACE} -s 128 -w ${outfn} -a duration:60 -f 'tcp or udp'
