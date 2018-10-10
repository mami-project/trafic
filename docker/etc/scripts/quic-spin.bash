#!/bin/bash

set -eum

# measurements not implemented yet
IFACE=${IFACE:-None}
HOST=${HOST:-iperf}
EXID=${EXID:-baseline}			# Use the baseline profiles for this

CAPTIME=90

function mklabel() {
	local exid=$1
	local unixtime=$(date +%s)

	printf "%s-%s" "${exid}" "${unixtime}"
}

load=85							#  TODO loop with all loads
label=$(mklabel "QUIC")
capfn="${label}.pcap"
exid="${EXID}-${load}"

# start background traffic

# start servers
wget --header "X-CONF: ${exid}.env" \
	 -O /dev/null \
	 http://iperf-server:9000/hooks/start-servers
sleep 1

# start clients
wget --header "X-CONF: ${exid}.env" \
	 --header "X-LABEL: ${label}" \
	 --header "X-DB: ${EXID}" \
	 -O /dev/null \
	 http://iperf-client:9000/hooks/start-clients
sleep 2

# start QUIC server
wget -O /dev/null \
	 --header "X-CONF: quic.env" \
	 http://${HOST}-server:9000/hooks/quic-server
sleep 1

if [ "${IFACE}" != "None" ]; then
	# start capture for 70s
	# should be plenty of time for all clients to start and finish
	#
	tshark -i ${IFACE} -s 128 -w ${capfn} -f 'tcp or udp' -a duration:${CAPTIME} & WAITPID=$!
	sleep 1
fi

# start QUIC client
wget -O /dev/null \
	 --header "X-CONF: quic.env" \
	 http://${HOST}-client:9000/hooks/quic-client

if [ "${IFACE}" != "None" ]; then
	wait $WAITPID
	# try to save as much space as possible
	bzip2 -9 ${capfn}
else
	sleep ${CAPTIME}
fi
wget -O /dev/null http://${HOST}-server:9000/hooks/stop-servers
wget -O /dev/null http://${HOST}-client:9000/hooks/stop-clients
