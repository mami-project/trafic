#!/bin/bash

set -eu

EXID=${EXID:-flows}
CONF=${CONF:-${EXID}.env}

base=$(dirname $0)
. ${base}/${CONF}
. ${base}/fun.bash

label=$(mklabel "$EXID")
STATS=${STATS:-/root/share/stats/$label}

echo schedule clients \
	--stats-dir="${STATS}" \
	--log-tag=C \
	--flows-dirs=${FLOWS} \
	--influxdb-enabled \
	--influxdb-endpoint=http://influxdb:8086 \
	--influxdb-db=${label}
