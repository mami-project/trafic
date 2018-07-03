#!/bin/bash
#run clients and send stats to influxdb host
set -eu

LABEL=${LABEL:-"lola-flows"}
CONF=${CONF:-flows.env}
DB=${DB:-lola}
INFLUXDB=${INFLUXDB:-influxdb}

base=$(dirname $0)
. ${base}/${CONF}

STATS=${STATS:-${HOME}/share/stats/$LABEL}

schedule clients \
	--log-tag=C \
	--flows-dirs=${FLOWS} \
	--influxdb-enabled \
	--influxdb-endpoint=http://${INFLUXDB}:8086 \
	--influxdb-db=lola-${DB} \
	--influxdb-measurements=${LABEL}
