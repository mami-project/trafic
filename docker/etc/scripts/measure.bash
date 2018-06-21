#!/bin/bash

set -eu

function mk-label() {
	local exid=$1
	local unixtime=$(date +%s)

	printf "lola-%s-%s" "${exid}" "${unixtime}"
}

IFACE=${IFACE:-eth0}
EXID=${EXID:-baseline}

for load in 75 80 85 90 95
do
	exid="${EXID}-${load}"
	label=$(mk-label "${exid}")
	# capfn="${label}.pcap"

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

	sleep 65	# allow some time for flows to drain

	# cleanup (and, possibly, go again)
	wget http://iperf-server:9000/hooks/stop-servers -O /dev/null
	wget http://iperf-client:9000/hooks/stop-clients -O /dev/null

	# send measurements to the postprocessing machine and clean up
	wget http://iperf-client:9000/hooks/send-stats -O /dev/null
done
