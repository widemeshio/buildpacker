#!/bin/env bash
go build ./cmd/pack-shimmer && ./pack-shimmer build "docker.io/bithavoc/image1" --builder heroku/buildpacks:18
