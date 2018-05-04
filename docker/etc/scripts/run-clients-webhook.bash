#!/bin/bash

set -eux

HOOKS="/root/share/scripts/client-hooks.json"

webhook -verbose -hooks "${HOOKS}"
