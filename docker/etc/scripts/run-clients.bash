#!/bin/bash

set -exu

base=$(dirname $0)
. $base/flows.env

schedule clients \
	--stats-dir=/root/share/stats \
	--log-tag=C \
	--flows-dirs=${FLOWS} \
	--influxdb-enabled \
	--influxdb-endpoint=http://influxdb:8086 \
	--influxdb-db=lola
