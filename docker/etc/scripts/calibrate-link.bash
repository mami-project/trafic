#!/bin/bash

set -eu

function mk-label() {
	local exid=$1
	local ext=$2
	local unixtime=$(date +%s)

	printf "lola-%s-%s%s" "${exid}" "${unixtime}" "${ext}"
}
cwd=`dirname $0`
iperf3 -c iperf-server --reverse --time 30 --logfile ${cwd}/$(mk-label "calibrate").txt
