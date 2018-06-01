# Trafic
Prototype of a traffic mix generator based on [iperf3](https://iperf.fr).

## Dependencies

A UNIX or UNIX-like OS

### Go

Follow instructions at https://golang.org/doc/install

Also, a few 3rd party packages *trafic* depends on:
```
go get -u gopkg.in/yaml.v2 github.com/spf13/cobra github.com/spf13/viper github.com/alecthomas/units
```

If you plan to use *trafic* to run the lola tests, webhook can also come in handy to automate at least part of the execution:
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

**

[Trafic](https://en.wikipedia.org/wiki/Trafic) is not a typo, it's a film by [Jacques Tati](https://en.wikipedia.org/wiki/Jacques_Tati).

**

# Documentation

At its core, *trafic* is just a flow scheduler.

You describe one or more flows, for example specifying which transport protocol, transmission patterns, markings, etc. and let *trafic* run the client and server side of that flow at the specified time.

When a flow completes, the *trafic* client stores key performance indicators (KPI) associated with each scheduled flow (e.g., bandwidth, packet loss, RTT, jitter).

These KPIs are sampled at a configurable rate (e.g. every 200ms) and made available as (per-flow) CSV files.  If requested, they can also be sent to an [Influx](https://www.influxdata.com/time-series-platform/influxdb/) instance.

On top of that, *trafic* also provides the ability to define a traffic mix at high level, in terms of:
* How much bandwidth does the mix take;
* Which applications compose the mix (ABR video, web browsing, real-time A/V, etc.);
* Which percentage of bandwidth is used by each application.

## Modelling a flow

A flow is completely described by a YAML file.  The file has three parts:
* General;
* Client specific;
* Server specific.

The `client` and `server` sections are further subdivided into two parts: an `at` block containing scheduling information, and a role specific configuration.  (See the full list of [client](etc/client-blueprint.yaml) and [server](etc/server-blueprint.yaml) configurables.)

The default direction of a flow is server to client.

### Example: ABR video

To give an example, the following file describes an adaptive bit rate (ABR) video flow composed of three HD (960x720) video chunks, 10 seconds each:
```
label: &l abr-video

port: &p 5400

client:
  at:
    - 0s
    - 10s
    - 20s
  config:
    server-address: trafic-server.example.org.
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

The flow is given a label "abr-video", which is used internally by the scheduler to reference the flow and _must be unique_ among flows in the same mix.  Note that the flow is TCP unless otherwise specified.

Server is scheduled once, at the very start of the scheduler execution (`0s`).  It will bind port 5400, and will do its KPI sampling every 200ms.

A new client instance is scheduled to run every 10 seconds (at `0s`, `10s` and `20s`) and download 1.8MB worth of data, i.e. roughly a 10s HD video chunk.  Each client instance will connect to port `5400` on host `trafic-server.example.org`.  In this case as well, KPI sampling happens every 200ms.  

### Example: realtime audio

The following configuration models two instances of a mono-directional realtime audio flow (half of a typical Skype call), with regularly paced UDP packets bearing 126 bytes of RTP and media payload (`length`) aiming at constantly injecting 64Kbps (`target-bitrate`) in the network.  The two flows run, in parallel, for 60 seconds.

Things worth noting:
* An UDP flow needs to be marked explicitly: `udp: true`;
* A constant bitrate flow needs to be temporally bounded: `time-s: 60`;
* Parallelism of flows can be specified with the `parallel` keyword;

```
label: &l rt-audio

port: &p 5000

instances: &i 2

client:
  at:
    - 0s
  config:
    server-address: trafic-server.example.org.
    server-port: *p
    time-s: 60
    udp: true
    length: 126
    target-bitrate: 64K
    title: *l
    report-interval-s: 0.2
    parallel: *i

server:
  at:
    - 0s
  config:
    server-port: *p
    report-interval-s: 0.2
```

## Designing a traffic mix

A traffic mix is completely described by a YAML file.

The file has a header section defining the bandwidth to fill (`total-bandwidth`) and for how long the mix shall run (`total-time`).

```
# target aggregate bandwidth in bytes/sec (B/KB/MB/GB/TB)
total-bandwidth: 12.5MB

# how long the mix should run - expressed as duration (s, m, h, etc.)
total-time: 60s

# measure sampling tick
report-interval: 0.2s
```

The high level description of the application flows is done in the `flows` section.

Each flow defines its `kind`, i.e. the application it simulates.  The available pre-defined applications are:
* realtime-audio
* realtime-video
* scavenger
* greedy
* abr-video
* web-page

The amount of bandwidth this application consumes out of the total available (`total-bandwidth`) is given as a percentage using the `percent-bandwidth` keyword.

The ports used by the server are specified as a range using the `ports-range` keyword.

```
server: &srv trafic-server.example.org.

flows:
  - kind: realtime-audio
    percent-bandwidth: 1%
    ports-range: 5000-5099
    props:
      label: rt-audio
      server: *srv
```

## Running the mix

TODO

## Exploring KPIs

TODO




