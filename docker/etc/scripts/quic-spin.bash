#!/bin/bash

set -eum

# measurements not implemented yet
IFACE=${IFACE:-None}
HOST=${HOST:-iperf}

CAPTIME=90

function mklabel() {
	local exid=$1
	local unixtime=$(date +%s)

	printf "%s-%s" "${exid}" "${unixtime}"
}

label=$(mklabel "QUIC")
capfn="${label}.pcap"

# start background traffic
wget -O /dev/null \
	 http://${HOST}-server:9000/hooks/quic-background
sleep 1
# start TCP background client
wget -O /dev/null \
	 --header "X-TIME: $((CAPTIME - 5))" \
	 http://${HOST}-client:9000/hooks/calibrate

wget -O /dev/null \
	 --header "X-TIME: $((CAPTIME - 5))" \
	 http://${HOST}-client:9000/hooks/upd-fill.bash

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
