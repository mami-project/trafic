# Integrating traffic in an NFV environment

This file describes how to create the basic VM image used in the experiments and how to integrate it in an NFV environment.
## Alpine Linux-based VM
[Alpine Linux](https://www.alpinelinux.org) is a Linux distribution that adapts very well to small footprint environments and uses a hardened Linux kernel. Alpine is based on packages, that are very up-to-date.
We will start with an [image of the 3.7 version](http://dl-cdn.alpinelinux.org/alpine/v3.7/releases/x86_64/alpine-virt-3.7.0-x86_64.iso), that we will update to the current experimental branch in order to include the latest influxes and iperf3 packages.

### Creating the initial image

Using your favourite virtualisation platform (libvirt-based platforms like [QEMU](http://wiki.qemu.org/Index.html) recommended) create a QCOW2 hard disk image file of (at least) 5GBytes for the client and server VMs and 10GBytes for the capture and control VM.

Boot a virtual machine with the CDROM and disk images you just downloaded and created. 512M memory should suffice.

Follow the instructions and create the initial image with the `setup-alpine` command. Reboot the VM, booting from the hard disk.

### Customising the image

Edit the `/etc/apk/repositories` file, uncommenting all lines referring to the repositories of the `3.7`distribution and the line referring to the `main` repository of the `edge` distribution. The file should look like this:

```
#/media/cdrom/apks
http://uk.alpinelinux.org/alpine/v3.7/main
http://uk.alpinelinux.org/alpine/v3.7/community
http://uk.alpinelinux.org/alpine/edge/main
http://uk.alpinelinux.org/alpine/edge/community
#http://uk.alpinelinux.org/alpine/edge/testing
```

Update the VM with the following commands:

```
apk update
apk upgrade
```

Modify the root user to use bash as login shell. Log out and login again as root.
Modify `/etc/ssh/sshd_config` to allow root login: search `PermitRootLogin` in the file and change the line to read:

```PermitRootLogin yes```

Reboot the VM once again.

Install the following packages with the `apk add` command:

 - `bash`
 - `git`
 - `musl-dev`
 - `make`

Finally install  `iperf3` and `go` from the *testing* repository. (This will install the required versions for these packages.)

```
apk add go=1.10.1-r0
apk add iperf3=3.5-r1
```



### Installing the trafic sources

Before installing the traffic sources from GitHub, create the following `.profile` for the root user:
```
PATH=$(go env GOPATH)/bin:$PATH
export PATH
```

Then, prepare the infrastructure for Go:
```
mkdir -p $(go env GOPATH)/{bin,src}
```

Then install the trafic repository following the guide included in `README.md`.

You may safely ignore messages related to cgo:
```
# github.com/adnanh/webhook/test
loadinternal: cannot find runtime/cgo
trafic:~/go/src/github.com/adnanh/webhook# webhook
```

Finally, create the `$HOME/share` directory and link the scripts and flows directories as follows:

```
mkdir ${HOME}/share
cd ${HOME}/share
ln -sf ${HOME}/go/src/github.com/mami-project/trafic/docker/etc/scripts
ln -sf ${HOME}/go/src/github.com/mami-project/trafic/docker/etc/flows
```

Shutdown the machine. Keep the QCOW2 file you just generated in a safe place. This is the base image used for all Virtual Network Function Components.

### Specific install for the InfluxDB VM

The InfluxDB VM does not need the `trafic` sources and therefore, it will not need the `go` distribution either. The packages will suffice:

 - `bash`
 - `influxdb`

Make sure that the `influxdb` service is started and added to the default level in the OpenRC infrastructure:

```
rc-service -I influxdb start
rc-update add influxdb default
```

### Specific install for the tshark VM

The `tshark` VM is created as a regular `trafic` VM. One everything is installed, include the `tshark` package.

## Integration with NFV

Currently, Alpine Linux has a very limited [cloud-init](https://cloud-init.io) support that is not enough to automatically configure VM instances
at boot time. Therefore the integration has to be done manually on a case by case basis.
Alternative Linux distributions and OSes with better cloud-init support are being evaluated.

### Configuration requirements of the measurement scenario

The scenario is composed of four VMs:
- `iperf-client`,`iperf-server` and `tshark` have two Ethernet interfaces, one for control and one for measuring.
- all four are connected to a common *control* LAN.
- the measurement interface of `iperf-client` is recommended to be a passthrough interface and is connected to the client side of the test network
- the measurement interface of `iperf-server` is recommended to be a virtualisation-friendly interface and is connected to the server side of the test network
-  the measurement interface of `tshark` receives the traffic present in the test interface of `iperf-client` (through a port mirror, for example).

```
 control
    |
    |     +--------------+
    |     |              |                 |
    +-----+ iperf-server +-----------------+
    |     |              |                 |
    |     +--------------+                 |
    |                                Network under
    |                                     test
    |     +--------------+                 |
    |     |              |     mirror      |
    +-----+ iperf-client +--------+--------+
    |     |              |        |        |
    |     +--------------+        |
    |                             |
    |     +--------------+        |
    |     |              |        |
    +-----+    tshark    |<-------+
    |     |              |
    |     +--------------+
    |
    |     +--------------+
    |     |              |
    +-----+   influxdb   |
    |     |              |
    |     +--------------+
```

### Setting up the communications

External access is *required* for `tshark` (ssh) and `influxdb` (http/https). `iperf-client` and `iperf-server` can be accessed from ` tshark` to start the client and server scripts.

All VMs need to resolve the names of all VMs in the scenario. Add entries for `iperf-client`, `iperf-server`,` influxdb` and `tshark` in the `/etc/hosts`. Following best current practices use 127.0.1.1 for the local name and 127.0.0.1 for localhost.

Example:

```

127.0.0.1 localhost

10.7.1.40 iperf-client
10.7.1.41 iperf-server
127.0.1.1 influxdb
10.7.1.42 tshark
```

Configure `iperf-client` to communicate with `iperf-server` using the *Network under test*.

### Note: regarding go compilation warnings

If you experience problems compiling further down, you may need to add the following to your `.profile`:

```
GOOS=linux
export GOOS
GOARCH=amd64
export GOARCH
CGO_ENABLED=0
export CGO_ENABLED
```
