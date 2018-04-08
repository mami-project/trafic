FROM golang:alpine

RUN apk add --no-cache \
      bash \
      bc \
      bridge \
      ca-certificates \
      coreutils \
      drill \
      git \
      iftop \
      iperf3 \
      iproute2 \
      iproute2-doc \
      iputils \
      man \
      tcpdump \
      tshark \
      util-linux

RUN update-ca-certificates

RUN go get -u github.com/mami-project/trafic/schedule \
              github.com/spf13/viper \
              github.com/spf13/cobra

# Compile and install the trafic scheduler
RUN cd /go/src/github.com/mami-project/trafic/schedule && go install


# Manually fix traffic control distributions.  They are actually installed
# but not in the location tc expects to find them.
RUN ln -s /usr/lib/tc /lib/tc

# netemd is launched in the entrypoint script after network interfaces have
# been configured
COPY entrypoint.bash /root/

# The entrypoint for this container is defined in docker-compose.yml
