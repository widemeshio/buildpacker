#!/bin/env bash
go build ./cmd/pack-shimmer && ./pack-shimmer --path scripts/test build "docker.io/bithavoc/image1" --builder heroku/buildpacks:18 --buildpack https://github.com/weibeld/heroku-buildpack-run.git
