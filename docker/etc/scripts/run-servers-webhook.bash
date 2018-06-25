#!/bin/bash

set -eux

HOOKS="${HOME}/share/scripts/server-hooks.json"

webhook -verbose -hooks "${HOOKS}"
