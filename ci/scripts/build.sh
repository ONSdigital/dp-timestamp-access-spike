#!/bin/bash -eux

pushd dp-timestamp-access-spike
  make build
  cp build/dp-timestamp-access-spike Dockerfile.concourse ../build
popd
