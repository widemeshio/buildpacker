#!/usr/bin/env bash
go build ./cmd/pack-shimmer && ./pack-shimmer --path scripts/test_go build "docker.io/bithavoc/image1" \
    --builder heroku/buildpacks:18 \
    --buildpack "https://github.com/heroku/heroku-buildpack-go.git#v149" \
    --buildpack "https://buildpack-registry.s3.amazonaws.com/buildpacks/deadmanssnitch/field-agent.tgz" \
    --buildpack "heroku/procfile"
