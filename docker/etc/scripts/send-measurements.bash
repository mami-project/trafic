#!/bin/bash
set -eu
# 
# Get all measurements, compress them, send them to analyse and cleanup
#
base=$(dirname $0)

STATS=${STATS:-/root/share/stats/}

cd $STATS
for d in lola-*; do 
	tar -cvjf $d.tar.bz2 $d
	scp $d.tar.bz2 student@analyse:
	rm -vf $d.tar.bz2
	rm -rvf $d
done
