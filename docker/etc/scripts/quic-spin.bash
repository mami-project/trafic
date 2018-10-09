#!/bin/bash

set -eum

# measurements not implemented yet
IFACE=${IFACE:-None}
HOST=${HOST:-iperf}

CAPTIME=75

# start server
wget -nv -O /dev/null \
	 --header "X-CONF: quic.env" \
	 http://${HOST}-server:9000/hooks/quic-server
sleep 1

if [ "${IFACE}" != "None" ]; then
	# start capture for 70s
	# should be plenty of time for all clients to start and finish
	#
	echo This is not ready yet... experiment without capture
	if [ "${IFACE}" == "None" ]; then
		tshark -i ${IFACE} -s 128 -w ${capfn} -f 'tcp or udp' -a duration:${CAPTIME} & WAITPID=$!
		# start the clients
		sleep 1
		wget -nv -O /dev/null \
			 --header "X-CONF: quic.env" \
			 http://${HOST}-client:9000/hooks/quic-client
		# wait for the capture to finish
		wait $WAITPID
		# try to save as much space as possible
		bzip2 -9 ${capfn}
	else
		sleep 1
		wget -nv -O /dev/null \
			 --header "X-CONF: quic.env"
		http://${HOST}-client:9000/hooks/quic-client
		sleep ${CAPTIME}
	fi
else
	sleep 1
	wget -nv -O /dev/null \
		 --header "X-CONF: quic.env"
	http://${HOST}-client:9000/hooks/start-clients
	sleep ${CAPTIME}
fi
