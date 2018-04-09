#!/bin/bash

set -exu

#FLOWS=/root/share/flows/lola-mix/realtime-media
#FLOWS=/root/share/flows/lola-mix/web-short
FLOWS=/root/share/flows/lola-mix/tcp-high-throughput

schedule servers \
	--log-tag=S \
	--flows-dirs=${FLOWS}
