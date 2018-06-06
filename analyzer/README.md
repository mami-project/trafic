# Trafic analyser

## Dependencies

OS-independent

Python3
GNUPlot5 for post-processing

### Python 3 libraries:

#### requests: http://docs.python-requests.org/

Recommended installation: use `pip install requests`

#### python-rfc3339: https://github.com/tonyg/python-rfc3339.git

Installation: clone the github repository and install using

`python setup.py install`

## Programs included

### getids

List the flow IDs currently stored in the influxdb server
```
./getids
usage: getids [-h] [-d DATABASE] [-H HOST]

An analyser for the influxdb data

optional arguments:
  -h, --help            show this help message and exit
  -d DATABASE, --database DATABASE
                        The database we want to query (default is lola-
                        baseline)
  -H HOST, --host HOST  The database server (default is influxdb)
```

Use this program to obtain the flow IDs you will use in the `analyse-lola` command.

### analyse-lola

This is the main cruncher. It will take an influxdb database with your measurements, extract the information collected from the ABR video streams and make a simulation of the playback based on [1]

Usage:
```
$ ./analyse-lola -h
usage: analyse-lola [-h] [-d DATABASE] [-H HOST] [-R] FLOWID

An analyser for the influxdb data

positional arguments:
  FLOWID                Analyse this flow

optional arguments:
  -h, --help            show this help message and exit
  -d DATABASE, --database DATABASE
                        The database we want to query (default is lola-
                        baseline)
  -H HOST, --host HOST  The database server (default is influxdb)
  -R, --raw             Spit out the raw data and quit (default is to analyse
                        the data)
```

Normally, this program will grab data from the influxdb host and crunch them to a format that is then postprocessed with GNUplot to produce a graphical representation of the playback simulation. Overall results of the simulation with QoE parameters are included as GNUplot comments.

In addition, `analyse-lola` can also extract the flow from the database and print the raw data as JSON to stdout. This can be used to store the data locally for offline analysis.

### local-analyse

The raw datafiles can also be analysed locally with this program.
```
$ ./local-analyse -h
usage: local-analyse [-h] [-i INPUT]

An analyser for the influxdb data

optional arguments:
  -h, --help            show this help message and exit
  -i INPUT, --input INPUT
                        The raw JSON data from analyse-lola
```

### Sample scripts for GNUPlot

With getids, we determined that one of the ABR flows at 75% link load in a baseline experiment (i.e. no AQM, no QCI mapping) had the
flow ID `lola-baseline-75-1528175109`. The followig GNUPlot script will produce a joint plot of the flows and the progress of the video playback buffer accessing the influxdb server

```
#!/usr/bin/env gnuplot
set term aqua
set key outside
set xlabel "Time slots = 0.1 s"
set ylabel "Video bytes"

#set output "abr-75.eps"
set title "lola-baseline-75 (ABR)"
plot for [col=3:4] "<./analyse-lola lola-baseline-75-1528175109" using 1:col with dots title columnheader
```

If the raw data were stored locally in `lola-baseline-75-1528175109.json` the script can also use the offline analysis tool, The last line would then be
```
plot for [col=3:4] "<./local-analyse -i lola-baseline-75-1528175109.json" using 1:col with dots title columnheader
```

In both bases, the relative timestamp (expressed in tenths of seconds) is in column 1, the accumulated amount of received data is given in column 3 and the progress of the video playback buffer is given in column 4.

# References

[1] [Biernacki: "Performance of HTTP video streaming under different network conditions"](https://doi,org/10.1007/s11042-013-1424-x)
