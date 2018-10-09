#!/bin/bash

set -exu

CONF=${CONF:-quic.env}

base=$(dirname $0)
. ${base}/${CONF}

PUB_IP=$(ip address show dev ${PUB_IF} | awk '/inet /{print $2}' | awk -F/ '{print $1}')
printf "Will run QUIC server on %s:%d\n" ${PUB_IP} ${QPORT}

flowsim server -Q -I ${PUB_IP} -p ${QPORT} ${SERVER_FLAGS}
