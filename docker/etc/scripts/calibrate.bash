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
HOST=${HOST:-iperf}
# start server
wget -O /dev/null \
	 http://${HOST}-server:9000/hooks/calibrate
sleep 1
# start client
wget -O /dev/null \
	 --header "X-REVERSE: FALSE" \
	 http://${HOST}-client:9000/hooks/calibrate

sleep 40
wget -O /dev/null \
	 http://${HOST}-server:9000/hooks/calibrate
sleep 1
# start client
wget -O /dev/null \
	 --header "X-REVERSE: TRUE" \
	 http://${HOST}-client:9000/hooks/calibrate

sleep 40

wget -O /dev/null \
	 http://${HOST}-client:9000/hooks/stop-clients
wget -O /dev/null \
	 http://${HOST}-server:9000/hooks/stop-servers
