#!/bin/bash

set -exu

CONF=${CONF:-quic.env}

base=$(dirname $0)
. ${base}/${CONF}

printf "Will talk to QUIC server on %s:%d\n" ${SERVER_IP} ${QPORT}

flowsim client -Q -I ${SERVER_IP} -p ${QPORT} -N ${REQUEST} -n ${TIMES} -t ${DELAY}
