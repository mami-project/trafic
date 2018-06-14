#!/bin/bash

set -eu

function mklabel() {
	local exid=$1
	local unixtime=$(date +%s)

	printf "lola-%s-%s" "${exid}" "${unixtime}"
}

iperf3 -c iperf-server --reverse --time 30 --logfile $(mk-label "calibrate").txt
