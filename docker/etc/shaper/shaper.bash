#!/bin/bash
#
# LTE (UE<->SGi LAN)
# - UL: MAXRATE=50Mbit  LATENCY=92ms
# - DL: MAXRATE=100Mbit LATENCY=76ms

set -eu

IFACE=${IFACE?set IFACE to the network interface you want to shape, e.g.: eth0}
LATENCY=${LATENCY?set LATENCY to the upstream network latency, e.g.: 50ms}
MAXRATE=${MAXRATE?set MAXRATE to the maximum upstream throughput rate, e.g.: 40Mbit}
IFACE_MAXRATE=${IFACE_MAXRATE:-1000Mbps}

# drop any previous setting on the interface
echo ">> Dropping previous settings (if any) on ${IFACE}"
tc qdisc del dev ${IFACE} root || true

# shape upstream bandwidth and latency
echo ">> Setting upstream latency to ${LATENCY} and max throughput rate to ${MAXRATE} on ${IFACE}"
tc qdisc add dev ${IFACE} handle 1: root htb default 11
tc class add dev ${IFACE} parent 1: classid 1:1 htb rate ${IFACE_MAXRATE}
tc class add dev ${IFACE} parent 1:1 classid 1:11 htb rate ${MAXRATE}
tc qdisc add dev ${IFACE} parent 1:11 handle 10: netem delay ${LATENCY}

# show current configuration
echo ">> Shaping settings on ${IFACE}:"
tc qdisc show dev ${IFACE}
