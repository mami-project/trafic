#!/bin/bash

set -exu

CONF=${CONF:-udp-fill.env}

base=$(dirname $0)
. ${base}/${CONF}

iperf3 --client ${SERVER} --udp --port ${PORT} --parallel ${THREADS} --time ${TIME} ${XTRAS}
