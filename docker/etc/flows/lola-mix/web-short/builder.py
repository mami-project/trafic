#!/usr/bin/env python3

from random import randint, uniform
from jinja2 import Environment

YAML = """
# 3 instances doing, on average, a ~1.2M transfer (the average Alexa top-1000 page size)

label: &l {{ label }}

port: &p {{ port }}

instances: &i 3

client:
  at:
{% for at in ats -%} 
{{ indent }}- {{ at }}s
{% endfor %}

  config:
    server-address: iperf-server
    server-port: *p
    title: *l
    bytes: 3738K
    # report-interval-s: 0.2
    parallel: *i

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
def make_ats(nats, delta):
    # pick random point in [0, delta-1] (using decisecond precision)
    t0 = float("{:.1f}".format(uniform(0, delta - 1)))
    # generate as many point as requested starting at t0 every delta
    return [t0 + offset for offset in range(0, nats * delta, delta)]

if __name__ == "__main__":
    HOW_MANY = 6
    BASE_PORT = 5004

    for i in range(0, HOW_MANY * 2, 2):
        port = BASE_PORT + i
        label = "tcp-web-short-%d" % port
        ats = make_ats(8, 8)

        do_conf(label, port, ats)
