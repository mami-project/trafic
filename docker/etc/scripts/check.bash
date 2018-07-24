#!/bin/bash
set -eu
HOST=${HOST:-iperf}
wget -O /dev/null \
	 http://${HOST}-client:9000/hooks/check-link
