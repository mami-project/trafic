#!/bin/bash

set -eux

HOOKS="/root/share/scripts/server-hooks.json"

webhook -verbose -hooks "${HOOKS}"
