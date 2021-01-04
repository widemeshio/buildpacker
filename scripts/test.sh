#!/usr/bin/env bash

go build ./cmd/buildpacker && ./buildpacker --path scripts/test_go \
    build "docker.io/bithavoc/image1" \
    --id-file test-ids.json \
    --trust-builder \
    --builder "heroku/buildpacks:18" \
    --buildpack "https://github.com/heroku/heroku-buildpack-go.git#v149" \
    --buildpack "https://github.com/weibeld/heroku-buildpack-run.git" \
    --buildpack "https://buildpack-registry.s3.amazonaws.com/buildpacks/aissaoui-ahmed/vim.tgz" \
    --buildpack "heroku-community/awscli" \
    --buildpack "heroku/procfile"
