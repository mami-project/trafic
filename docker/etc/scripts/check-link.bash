#!/bin/bash

set -eu

ip route | grep default | grep 192.168
ip route | grep 10.3

ping -c 5 iperf-server
