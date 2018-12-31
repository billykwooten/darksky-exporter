.DEFAULT_GOAL := help

# Project name is the same as the binary name in .goreleaser.yml
PROJECTNAME := darksky-exporter

GORELEASER_VERSION := 0.95.2
TAG := $(shell cat version/version.go | grep "Version" | head -1 | sed 's/\"//g' | cut -d' ' -f3 )

## build: Build local binaries and docker image.
build:
	@echo "=> Building with goreleaser ..."
	git tag -a v$(TAG)
	goreleaser release --skip-publish
.PHONY: build

## build-image: Build just docker image.
build-image:
	@echo "=> Building docker image ..."
	docker build -f Dockerfile -t "$(PROJECTNAME):v$(TAG)" .
	@echo "=> Cleanup ..."
	rm -rf $(PROJECTNAME)
.PHONY: build-image

test:
	@echo "=> Running Go Test via Goveralls ..."
	mkdir _test
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls
	overalls -covermode=atomic -project=github.com/billykwooten/$(PROJECTNAME) -- -race -v
	mv overalls.coverprofile _test/$(PROJECTNAME).cover
	go tool cover -func=_test/$(PROJECTNAME).cover
	rm -rf _test
.PHONY: test

## install-goreleaser-linux: Install goreleaser on your system for Linux systems.
install-goreleaser-linux:
	wget -nv -P /tmp/ https://github.com/goreleaser/goreleaser/releases/download/v$(GORELEASER_VERSION)/goreleaser_Linux_x86_64.tar.gz
	tar -C ~/bin -xzf /tmp/goreleaser_Linux_x86_64.tar.gz goreleaser
	rm -r /tmp/goreleaser_Linux_x86_64.tar.gz
.PHONY: install-goreleaser-linux

## install-goreleaser-darwin: Install goreleaser on your system for macOS (Darwin).
install-goreleaser-darwin:
	wget -nv -P /tmp/ https://github.com/goreleaser/goreleaser/releases/download/v$(GORELEASER_VERSION)/goreleaser_Darwin_x86_64.tar.gz
	tar -C /usr/local/bin -xzf /tmp/goreleaser_Darwin_x86_64.tar.gz goreleaser
	rm -r /tmp/goreleaser_Darwin_x86_64.tar.gz
.PHONY: install-goreleaser-darwin

## github-release: Publish a release to github.
github-release:
	@echo "=> Running Publish Release to Github ..."
	git tag -a v$(TAG)
	git push origin v$(TAG)
	goreleaser
.PHONY: github-release

## clean: Clean directory.
clean:
	@echo "=> Cleaning directory ..."
	rm -rf _dist/
	rm -rf _test/
.PHONY: clean

all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
.PHONY: help
