#!/bin/bash

set -exu

CONF=${CONF:-flows.env}

base=$(dirname $0)
. ${base}/${CONF}

schedule servers \
	--log-tag=S \
	--flows-dirs=${FLOWS}
