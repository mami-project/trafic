#!/bin/bash

set -eum

function mklabel() {
	local exid=$1
	local unixtime=$(date +%s)

	printf "lola-%s-%s" "${exid}" "${unixtime}"
}


# remote measurements don't include capturing traffic
IFACE=${IFACE:-None}
EXID=${EXID:-baseline}
HOST=${HOST:-iperf}

CAPTIME=70

max=100
load=75
while [ $load -lt $max ]; do
	exid="${EXID}-${load}"
	label=$(mklabel "${exid}")
	capfn="${label}.pcap"

	[ -n "$1" ] && printf "\n ----------------------------------\n Iteration %d " $1
	printf " Load %d%%\n" ${load}
	# start servers
	wget -nv -O /dev/null \
		 --header "X-CONF: ${exid}.env" \
		 http://${HOST}-server:9000/hooks/start-servers
	sleep 1

	# start clients
	if [ "${IFACE}" != "None" ]; then
		# start capture for 70s
		# should be plenty of time for all clients to start and finish
		#
		tshark -i ${IFACE} -s 128 -w ${capfn} -f 'tcp or udp' -a duration:${CAPTIME} & WAITPID=$!
		# start the clients
		sleep 1
		wget -nv -O /dev/null \
			 --header "X-CONF: ${exid}.env" \
			 --header "X-LABEL: ${label}" \
			 --header "X-DB: ${EXID}" \
			 http://${HOST}-client:9000/hooks/start-clients
		# wait for the capture to finish
		wait $WAITPID
		# try to save as much space as possible
		bzip2 -9 ${capfn}
	else
		sleep 1
		wget -nv -O /dev/null \
			 --header "X-CONF: ${exid}.env" \
			 --header "X-LABEL: ${label}" \
			 --header "X-DB: ${EXID}" \
			 http://${HOST}-client:9000/hooks/start-clients
		sleep ${CAPTIME}
	fi
	# cleanup (and, possibly, go again)
	wget -nv -O /dev/null \
		 http://${HOST}-server:9000/hooks/stop-servers
	wget -nv -O /dev/null \
		 http://${HOST}-client:9000/hooks/stop-clients
	sleep 5

	wget -nv -O /dev/null \
		 --header "X-LABEL: ${label}" \
		 http://${HOST}-client:9000/hooks/clean-stats

	sleep 5
	load=$((load + 5))
done
