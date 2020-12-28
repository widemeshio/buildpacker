#!/usr/bin/env bash
go build ./cmd/buildpacker && ./buildpacker --path scripts/test_node build "docker.io/bithavoc/image1" --builder heroku/buildpacks:18 --buildpack "https://github.com/heroku/heroku-buildpack-nodejs.git" --buildpack https://github.com/weibeld/heroku-buildpack-run.git
