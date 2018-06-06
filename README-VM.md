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

Edit the `/etc/apk/repositories` file, commenting the lines referring to the `3.7` repositories and uncommenting the lines referring to the `edge` repositories. Update the VM with the following commands:
```
apk update
apk upgrade
```
Reboot the VM once again.

Install the following packages with the `apk add` command:

 - `bash`
 - `iperf3`
 - `git`
 - `go`,`musl-dev`
 - `make`
 - `influxdb`
 - `tshark`

Modify the root user to use bash as login shell. Log out and login again as root.

### Installing the trafic sources

Before installing the traffic sources from GitHub, create the following `.profile` for the root user:
```
PATH=$(go env GOPATH)/bin:$PATH
export PATH
GOOS=linux
export GOOS
GOARCH=amd64
export GOARCH
CGO_ENABLED=0
export CGO_ENABLED
```
Then, prepare the infrastructure for Go:
```
mkdir -p $(go env GOPATH)/{bin,src}
```

Then install the trafic repository following the guide included in `README.md`.

 Finally, create the `$HOME/share` directory and link the scripts and flows directories as follows:
 ```
 mkdir ${HOME}/share
 cd ${HOME}/share
 ln -sf ${HOME}/go/src/github.com/mami-project/trafic/docker/etc/flows
 ln -sf ${HOME}/go/src/github.com/mami-project/trafic/docker/etc/scripts
 ```

Shutdown the machine. Keep the QCOW2 file you just generated in a safe place. This is the base image used for all Virtual Network Function Components.

## Integration with NFV

Currently, Alpine Linux has a very limited [cloud-init](https://cloud-init.io) support that is not enough to automatically configure VM instances
at boot time. Therefore the integration has to be done manually on a case by case basis.
Alternative Linux distributions and OSes with better cloud-init support are being evaluated.


