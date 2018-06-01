# Trafic
Prototype of a traffic mix generator based on [iperf3](https://iperf.fr).

## Dependencies

A UNIX or UNIX-like OS

### Go

Follow instructions at https://golang.org/doc/install

Also, a few 3rd party packages trafic depends on:
```
go get -u gopkg.in/yaml.v2 github.com/spf13/cobra github.com/spf13/viper github.com/alecthomas/units
```

If you plan to use trafic to run the lola tests, webhook can also come in handy to automate at least part of the execution:
```
go get github.com/adnanh/webhook
```

### Iperf3

Fetch the latest release from [here](https://github.com/esnet/iperf/releases)

### Docker

[Docker](https://docs.docker.com/install/) and [Docker Compose](https://docs.docker.com/compose/install/).

## Build and install

```
go install ./...
```

***

[Trafic](https://en.wikipedia.org/wiki/Trafic) is not a typo, it's a film by [Jacques Tati](https://en.wikipedia.org/wiki/Jacques_Tati).

***

# Documentation

At its core, trafic is just a flow scheduler.

You describe one or more flows, for example specifying which transport protocol, transmission patterns, markings, etc. and let trafic run the client and server side of that flow at the specified time.

When a flow completes, the trafic client stores key performance indicators (KPI) associated with each scheduled flow (e.g., bandwidth, packet loss, RTT, jitter).

These KPIs are sampled at a configurable rate (e.g. every 200ms) and made available as (per-flow) CSV files.  They can also be sent to a [Influx](https://www.influxdata.com/time-series-platform/influxdb/) instance.

On top of that, trafic also provides the ability to define a traffic mix at high level, in terms of:
* How much bandwidth does the mix take;
* Which applications compose the mix (ABR video, web browsing, real-time A/V, etc.);
* Which percentage of bandwidth is used by each application.

## Modelling a flow

A flow is completely described by a YAML file.  The file has three parts:
* General;
* Client specific;
* Server specific.

Client and server are further subdivided into an `at` block containing scheduling information, and a role specific configuration.  (See the full list of [client](etc/client-blueprint.yaml) and [server](etc/server-blueprint.yaml) configurables.)


For example, the following describes an adaptive bit rate (ABR) video flow composed of three HD (960x720) video chunks, 10 seconds each.

:
```
label: &l abr-video

port: &p 5400

client:
  at:
    - 0s
    - 10s
    - 20s
  config:
    server-address: trafic-server.example.org
    server-port: *p
    title: *l
    bytes: 1.8M
    report-interval-s: 0.2

server:
  at:
    - 0s
  config:
    server-port: *p
    report-interval-s: 0.2
```

## Designing a traffic mix

## Running the mix

## Exploring KPIs






