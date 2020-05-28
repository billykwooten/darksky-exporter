# Due to Apple purchasing Dark Sky, this exporter will continue to work until the end of 2021. Apple will most likely absorb Dark Sky into it's ecosystem at which point this exporter will stop working. See more at: https://blog.darksky.net/

# Dark Sky Prometheus Exporter

![Docker Pulls](https://img.shields.io/docker/pulls/billykwooten/darksky-exporter.svg)
![Docker Automated](https://img.shields.io/docker/cloud/automated/billykwooten/darksky-exporter.svg)
![Docker Build](https://img.shields.io/docker/cloud/build/billykwooten/darksky-exporter.svg)

Dark Sky Exporter is a prometheus exporter for weather metrics consumed from the [Dark Sky API](https://darksky.net/dev).

## Requirements

* Linux / MacOS / Windows
* [docker](https://www.docker.com)

## Development

* [go](https://golang.org/dl)
* [godep](https://github.com/tools/godep)
* [goreleaser](https://github.com/goreleaser/goreleaser)

## Configuration

Dark sky exporter can be controlled by both ENV or CLI flags as described below.

| Environment        	       | CLI (`--flag`)              | Default                 	    | Description                                                                                                      |
|----------------------------|-----------------------------|---------------------------- |------------------------------------------------------------------------------------------------------------------|
| `LISTEN_ADDRESS`           | `listen-address`            | `:9091`                     | The port to listen on |
| `APIKEY`                   | `apikey`                    | `<REQUIRED>`                | Your dark sky API key |
| `CITY`                     | `city`                      | `New York, NY`              | City/Location in which to gather weather metrics |
| `INTERVAL`                 | `interval`                  | `2m`                        | Interval to poll the dark sky API |

## Usage

```
# Export weather metrics from Seattle using binary
./darksky-exporter --city "Seattle, WA" --apikey mi4o2n54i0510n4510

# Export weather metrics from Seattle using docker
docker run -d --restart on-failure --name=darksky-exporter -p 9091:9091 billykwooten/darksky-exporter --city "Seattle, WA" --apikey mi4o2n54i0510n4510
```

## Building from Source

Prerequisites:

* [Go compiler](https://golang.org/dl/)


To build from source run `make build`, other options are below.

```
# make help

 Choose a command run in darksky-exporter:

  build                       Build local binaries and docker image. Requires `go` to be installed.
  build-image                 Build just docker image.
  install-goreleaser-linux    Install goreleaser on your system for Linux systems.
  install-goreleaser-darwin   Install goreleaser on your system for macOS (Darwin).
  github-release              Publish a release to github.
  clean                       Clean directory.

```

### Deploying to github

Github deployment utilized goreleaser to push to github.

First change the version in version/version.go to the correct tag release. Then run the following:

```
export GITHUB_TOKEN="<YOUR_TOKEN>"
export DOCKER_LOGIN='<USERNAME>'
export DOCKER_PASSWORD='<PASSWORD>'
make github-release
```
