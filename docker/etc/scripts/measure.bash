#!/bin/bash

set -eu

function mklabel() {
	local exid=$1
	local unixtime=$(date +%s)

	printf "lola-%s-%s" "${exid}" "${unixtime}"
}

function send_cmd () {
	cmd_url = $1
	shift
	wget $* -nv -O /dev/null ${cmd_url}
}
# remote measurements don't include capturing traffic
#IFACE=${IFACE:-eth0}
IFACE=
EXID=${EXID:-baseline}
HOST=${HOST:-iperf}

max=100
load=75
while [ $load -lt $max ]; do
	exid="${EXID}-${load}"
	label=$(mklabel "${exid}")
	capfn="${label}.pcap"

	# start servers
	send_cmd http://${HOST}-server:9000/hooks/start-servers \
			 --header "X-CONF: ${exid}.env"

	sleep 1

	# start clients
	send_cmd http://${HOST}-client:9000/hooks/start-clients \
			 --header "X-CONF: ${exid}.env" \
			 --header "X-LABEL: ${label}" \
			 --header "X-DB: ${EXID}"

	if [ -n "${IFACE}" ]; then
		# start capture for 60s
		tshark -i ${IFACE} -s 128 -w ${capfn} -f 'tcp or udp' -a duration:60
		# try to save as much space as possible
		bzip2 -9 ${capfn}
		sleep 5	# allow some time for flows to drain
	else
		sleep 65
	fi
	# cleanup (and, possibly, go again)
	send_cmd http://${HOST}-server:9000/hooks/stop-servers
	send_cmd http://${HOST}-client:9000/hooks/stop-clients
	sleep 5

	send_cmd http://${HOST}-client:9000/hooks/clean-stats \
			 --header "X-LABEL: ${label}"

	load=$((load + 5))
done
