#!/bin/bash

set -eum

# measurements not implemented yet
IFACE=${IFACE:-None}
HOST=${HOST:-iperf}

CAPTIME=75

# start server
wget -O /dev/null \
	 --header "X-CONF: quic.env" \
	 http://${HOST}-server:9000/hooks/quic-server
sleep 1

wget -O /dev/null \
	 --header "X-CONF: quic.env" \
	 http://${HOST}-client:9000/hooks/quic-client
sleep ${CAPTIME}
