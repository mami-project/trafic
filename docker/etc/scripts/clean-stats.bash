#!/bin/bash

set -eu

LABEL=${LABEL:-"lola-flows"}
STATS=${STATS:-${HOME}/share/stats/$LABEL}

rm -rvf ${STATS}
