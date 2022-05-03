# Cursecli

[![Current Tag](https://img.shields.io/github/v/tag/webhippie/cursecli?sort=semver)](https://github.com/webhippie/cursecli) [![Build Status](https://github.com/webhippie/cursecli/actions/workflows/general.yml/badge.svg)](https://github.com/webhippie/cursecli/actions) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Docker Size](https://img.shields.io/docker/image-size/webhippie/cursecli/latest)](https://hub.docker.com/r/webhippie/cursecli) [![Docker Pulls](https://img.shields.io/docker/pulls/webhippie/cursecli)](https://hub.docker.com/r/webhippie/cursecli) [![Go Reference](https://pkg.go.dev/badge/github.com/webhippie/cursecli.svg)](https://pkg.go.dev/github.com/webhippie/cursecli) [![Go Report Card](https://goreportcard.com/badge/github.com/webhippie/cursecli)](https://goreportcard.com/report/github.com/webhippie/cursecli) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/47d8475af4a64c49857835dd68781565)](https://www.codacy.com/gh/webhippie/cursecli/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/cursecli&amp;utm_campaign=Badge_Grade)

A commandline client to interact with Curseforge. For now it's mostly used to
fetch mods defined within modpack manifests.

## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our Docker images published on [Docker Hub][dockerhub] or [Quay][quay].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.17, at least that's the version we are using.

```console
git clone https://github.com/webhippie/cursecli.git
cd cursecli

make generate build

./bin/cursecli -h
```

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2022 Thomas Boerger <thomas@webhippie.de>
```

[releases]: https://github.com/webhippie/cursecli/releases
[dockerhub]: https://hub.docker.com/r/webhippie/cursecli/tags/
[quay]: https://quay.io/repository/webhippie/cursecli?tab=tags
[docs]: https://webhippie.github.io/cursecli/#getting-started
[golang]: http://golang.org/doc/install.html
