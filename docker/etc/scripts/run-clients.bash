#!/bin/bash

set -exu

. /root/share/scripts/flows.env

schedule clients \
	--stats-dir=/root/share/stats \
	--log-tag=C \
	--flows-dirs=${FLOWS}
#	--influxdb-enabled \
#	--influxdb-endpoint=http://influxdb:8086 \
#	--influxdb-db=lola
