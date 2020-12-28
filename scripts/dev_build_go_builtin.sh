#!/usr/bin/env bash
go build ./cmd/buildpacker && ./buildpacker --path scripts/test_go build "docker.io/bithavoc/image1" --builder heroku/buildpacks:18 --buildpack "heroku/go" --buildpack https://github.com/weibeld/heroku-buildpack-run.git --buildpack "heroku/procfile"
