#!/bin/bash -eux

export cwd=$(pwd)

pushd $cwd/dp-timestamp-access-spike
  make audit
popd