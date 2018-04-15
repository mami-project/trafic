#!/bin/bash

set -exu

. /root/share/scripts/flows.env

schedule servers \
	--log-tag=S \
	--flows-dirs=${FLOWS}
