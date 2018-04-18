#!/bin/bash

set -exu

base=$(dirname $0)
. $base/flows.env

schedule servers \
	--log-tag=S \
	--flows-dirs=${FLOWS}
