#!/bin/sh
set -x

CI_COMMIT_TAG=$(git describe --always --tags)

docker build -t linclaus/prometheus-operator:$CI_COMMIT_TAG -f build/package/Dockerfile .