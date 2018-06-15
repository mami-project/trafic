#!/bin/bash
#
# A simple screen to calibrate the link
# in server: iperf3 -s -1
# in client: iperf3 -c iperf3-server --reverse --time 30 --logfile <logfile>
#
# Wait one minute and kill any iperf process that eventually is alive
# (shouldn't kill anything)
#
set -eu

	# start server
wget --header "X-CONF: calibrate.env" \
	 -O /dev/null \
	 http://iperf-server:9000/hooks/calibrate

sleep 1

# start client
wget --header "X-CONF: calibrate.env" \
	 -O /dev/null \
	 http://iperf-client:9000/hooks/calibrate

sleep 60

wget --header "X-CONF: calibrate.env" \
	 -O /dev/null \
	 http://iperf-client:9000/hooks/stop-clients
wget --header "X-CONF: calibrate.env" \
	 -O /dev/null \
	 http://iperf-server:9000/hooks/stop-servers
