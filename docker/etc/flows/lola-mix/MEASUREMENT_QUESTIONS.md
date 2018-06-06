# Testing metodology, questions we want to answer

The experiments aim at comparing a baseline scenario where no traffic optimization is used (i.e., the control group) against the following experimental groups:
- A scenario where a LLT-like marking scheme is used; and
- A scenario where an AQM (e.g., PIE) is used before forwarding the traffic over the same (default) bearer.

The behaviour of each experimental group shall be assessed based on the measurable gain it produces over the baseline along two axes:
-	Its impact on the network (aggregate goodput); and
-	Its impact on endpoints (QoE and energy savings).

For a non zero-bit scheme like LLT, the latter can be furtherly split into:
-	Impact on cooperating endpoints; and
-	Impact on non-cooperating endpoints.

The two high-level questions that we want to answer, which map to two slightly different experiments, are:
- If an endpoint cooperates, does it gets an advantage?
- If a proportion of the population cooperates, do we get a better utilisation of resources?

In the following, we go into the details with regards to the configuration of the control and the experimental groups, the traffic mix composition and the set of reference flows – either belonging to the Lo or La classes – that will be measured.

# Methodology

1. Characterise the underlying \ac{ran} measuring the raw bandwidth delivered (with iperf3 in greedy mode)
  - Raw bandwidth delivered by the uplink to make sure it is not going to introduce any hidden limitation
  - Raw bandwidth delivered by the downlink to obtain the amount of background traffic
2. First measurement campain with background + ramping foreground on the \ac{ran} without any \ac{qos} treatment
  - Network aggregate goodput
  - \ac{qoe} experienced by the endpoints (on the foreground traffic flows)

