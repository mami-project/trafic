.DEFAULT_GOAL := baseimg

ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

# 3.7 is latest at the time of writing
ALPINE_LINUX_VERSION := 3.7

DOCKER_BUILD_FLAGS += --rm
DOCKER_BUILD_FLAGS += --tag alpine_$(ALPINE_LINUX_VERSION)-trafic

DOCKER_BASE_IMG_DIR := $(ROOT)

baseimg: Dockerfile ; docker build $(DOCKER_BUILD_FLAGS) $(DOCKER_BASE_IMG_DIR)

.PHONY: baseimg
