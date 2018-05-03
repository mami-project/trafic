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
