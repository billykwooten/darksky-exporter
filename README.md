# Dark Sky Prometheus Exporter

Dark Sky Exporter is a Prometheus exporter for weather metrics consumed from the [Dark Sky API](https://darksky.net/dev).

## Requirements

* Linux / MacOS / Windows
* [docker](https://www.docker.com)

## Development

* [go](https://golang.org/dl)
* [godep](https://github.com/tools/godep)
* [goreleaser](https://github.com/goreleaser/goreleaser)

## Configuration

Dark Sky Exporter can be controlled by both ENV or CLI flags as described below.

| Environment        	       | CLI (`--flag`)              | Default                 	    | Description                                                                                                      |
|----------------------------|-----------------------------|---------------------------- |------------------------------------------------------------------------------------------------------------------|
| `LISTEN_ADDRESS`           | `listen-address`            | `:9091`                     | The PORT to listen on |
| `APIKEY`                   | `apikey`                    | `<REQUIRED>`                | Your Dark Sky API Key |
| `CITY`                     | `city`                      | `New York, NY`              | City/Location in which to gather weather metrics |
| `INTERVAL`                 | `interval`                  | `2m`                        | Interval to poll the Dark Sky API |

## Usage

```
# Export weather metrics from Seattle using binary
./darksky-exporter-<os>-<arch> --city "Seattle, WA" --apikey mi4o2n54i0510n4510

# Export weather metrics from Seattle using docker
docker run -d --restart on-failure --name=darksky-exporter -p 9091:9091 darksky-exporter:<TAG> --city "Seattle, WA" --apikey mi4o2n54i0510n4510
```

## Building from Source

To build from source run `make build`, other options are below.

```
# make help

 Choose a command run in darksky-exporter:

  build                                                            Build local binaries and docker image.
  install-goreleaser-linux                                         Install goreleaser on your system for Linux systems.
  install-goreleaser-darwin                                        Install goreleaser on your system for macOS (Darwin).
  github-release                                                   Publish a release to github.
  clean                                                            Clean directory.

```
### Deploying to github

Github deployment utilized goreleaser to push to github.

First change the version in version/version.go to the correct tag release. Then run the following:

```
export GITHUB_TOKEN="<YOUR_TOKEN>"
make github-release
```
