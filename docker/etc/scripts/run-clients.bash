#!/bin/bash

set -exu

FLOWS=/root/share/flows/lola-mix/realtime-media
#FLOWS=${FLOWS},/root/share/flows/lola-mix/web-short
#FLOWS=${FLOWS},/root/share/flows/lola-mix/scavenger
#FLOWS=${FLOWS},/root/share/flows/lola-mix/tcp-high-throughput

schedule clients \
	--stats-dir=/root/share/stats \
	--log-tag=C \
	--flows-dirs=${FLOWS} \
	--influxdb-enabled \
	--influxdb-endpoint=http://influxdb:8086 \
	--influxdb-db=lola
