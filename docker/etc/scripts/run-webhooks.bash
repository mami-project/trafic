#!/bin/bash
dirn=$(dirname $0)
hosttype=$(hostname | sed 's/iperf-//g')
hosttype=$(echo ${hosttype}s | sed 's/ss/s/g')
bash -c ${dirn}/run-${hosttype}-webhook.bash
