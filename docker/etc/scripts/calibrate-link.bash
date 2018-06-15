#!/bin/bash

set -eu

function mklabel() {
	local exid=$1
	local ext=$2
	local unixtime=$(date +%s)

	printf "lola-%s-%s%s" "${exid}" "${unixtime}" "${ext}"
}
pwd=$(dirname $0)
iperf3 -c iperf-server --reverse --time 30 --logfile ${pwd}/`mklabel "calibrate" ".txt"`
