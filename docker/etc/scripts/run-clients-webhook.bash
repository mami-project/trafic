#!/bin/bash

set -eux

HOOKS="${HOME}/share/scripts/client-hooks.json"

webhook -verbose -hooks "${HOOKS}"
