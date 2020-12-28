# Widemesh Buildpacker
> Cloud-Native Buildpacks + Heroku Shims

[![Tests status](https://github.com/widemeshcloud/buildpacker/workflows/test/badge.svg)](https://github.com/widemeshcloud/buildpacker/actions?query=workflow%3Atest)

## What is Buildpacker?

Legacy buildpacks initially intended for Heroku do not work out of the box with [buildpacks.io](https://buildpacks.io/) `pack build` command.

Buildpacker makes all the old Heroku buildpacks work out of the box with the newest `pack build` command,
all you have to do is call `buildpacker build` instead and buildpacker invokes `pack build` for you.

## Features

* Git references with versioning `https://github.com/heroku/heroku-buildpack-go.git#v149`
* Tar.gz references `https://buildpack-registry.s3.amazonaws.com/buildpacks/aissaoui-ahmed/vim.tgz`
* Heroku Registry references such as `heroku-community/awscli`
* Automatic permission fixes for `bin/*`
* Automatic application of [`cnb-shim`](https://github.com/heroku/cnb-shim)

## Requirements

* [`pack`](https://buildpacks.io/)
* Docker
* Git

## License

See `LICENSE`
 
## Hacking this project

* git clone it
* install requirements in `$PATH`
* make your changes
* Test your changes with `scripts/test.sh`

Contributions Welcome.
