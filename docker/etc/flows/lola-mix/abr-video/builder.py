#!/usr/bin/env python3

from random import randint, uniform
from jinja2 import Environment

YAML = """
label: &l {{ label }}

port: &p {{ port }}

client:
  at:
{% for at in ats -%} 
{{ indent }}- {{ at }}s
{% endfor %}

  config:
    server-address: iperf-server
    server-port: *p
    title: *l
    bytes: 1.8M
    # report-interval-s: 0.2

server:
  at:
    - 0s
  config:
    server-port: *p
"""

def do_conf(label, port, ats):
    output = Environment().from_string(YAML).render(label=label, port=port, ats=ats, indent=4*" ")
    with open(label + ".yaml", "w") as fh:
        fh.write(output)

# nats is the number of points
# segment_length the length in seconds of each HLS segment
def make_ats(nats, segment_length):
    # pick random point in [0, segment_length-1] (using decisecond precision)
    t0 = float("{:.1f}".format(uniform(0, segment_length - 1)))
    # generate as many point as requested starting at t0 every segment_length
    return [t0 + offset for offset in range(0, nats * segment_length, segment_length)]

if __name__ == "__main__":
    HOW_MANY = 16
    BASE_PORT = 6000

    for i in range(0, HOW_MANY * 2, 2):
        port = BASE_PORT + i
        label = "tcp-abr-hd-%d" % port
        ats = make_ats(6, 10)

        do_conf(label, port, ats)
