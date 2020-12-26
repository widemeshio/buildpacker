#!/usr/bin/env bash
go build ./cmd/pack-shimmer && ./pack-shimmer --path scripts/test_go build "docker.io/bithavoc/image1" --builder heroku/buildpacks:18 --buildpack "https://github.com/heroku/heroku-buildpack-go.git#v149" --buildpack https://github.com/weibeld/heroku-buildpack-run.git --buildpack "heroku/procfile"
