#!/bin/bash

set -exu

schedule servers \
	--log-tag=S \
	--flows-dirs=/root/share/flows/lola-mix/web-short,/root/share/flows/lola-mix/realtime-media
