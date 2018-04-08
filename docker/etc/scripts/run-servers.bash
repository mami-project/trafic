#!/bin/bash

set -exu

FLOWS=/root/share/flows/lola-mix/realtime-media
#FLOWS=${FLOWS},/root/share/flows/lola-mix/web-short

schedule servers \
	--log-tag=S \
	--flows-dirs=${FLOWS}
