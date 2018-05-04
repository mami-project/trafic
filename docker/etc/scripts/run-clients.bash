#!/bin/bash

set -exu

function mklabel() {
	local exid=$1
	local unixtime=$(date +%s)

	printf "lola-%s-%s" "${exid}" "${unixtime}"
}

EXID=${EXID:-vanilla}
CONF=${CONF:-flows.env}

base=$(dirname $0)
. ${base}/${CONF}

label=$(mklabel "$EXID")
STATS=${STATS:-/root/share/stats/$label}

schedule clients \
	--stats-dir="${STATS}" \
	--log-tag=C \
	--flows-dirs=${FLOWS} \
	--influxdb-enabled \
	--influxdb-endpoint=http://influxdb:8086 \
	--influxdb-db=${label}
