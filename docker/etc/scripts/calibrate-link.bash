#!/bin/bash

set -eu

function mk-label() {
	local exid=$1
	local ext=$2
	local unixtime=$(date +%s)

	printf "lola-%s-%s%s" "${exid}" "${unixtime}" "${ext}"
}
REVERSE=${REVERSE:FALSE}
cwd=`dirname $0`
[ "$REVERSE" == "TRUE" ] && reverse="--reverse"
[ "$REVERSE" != "TRUE" ] && reverse=""

iperf3 -c iperf-server ${reverse} --time 30 --logfile ${cwd}/$(mk-label "calibrate").txt
