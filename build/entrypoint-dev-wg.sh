#!/bin/sh

export WG_CONFIG_PATH=/project/tmp/config/config.docker.yaml

export GOROOT=/usr/local/go
export GOPATH=/go
export PATH=/usr/local/go/bin:$PATH
export GOCACHE=/root/.cache/go-build

cd /project

/usr/bin/air -build.exclude_dir=frontend
