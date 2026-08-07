# Cursecli

[![General Workflow](https://github.com/webhippie/cursecli/actions/workflows/general.yml/badge.svg)](https://github.com/webhippie/cursecli/actions/workflows/general.yml) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/47d8475af4a64c49857835dd68781565)](https://app.codacy.com/gh/webhippie/cursecli/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Go Reference](https://pkg.go.dev/badge/github.com/webhippie/cursecli.svg)](https://pkg.go.dev/github.com/webhippie/cursecli) [![Go Report Card](https://goreportcard.com/badge/github.com/webhippie/cursecli)](https://goreportcard.com/report/github.com/webhippie/cursecli) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/webhippie/cursecli) [![Hosted By: Cloudsmith](https://img.shields.io/badge/OSS%20hosting%20by-cloudsmith-blue?logo=cloudsmith&style=flat-square)](https://cloudsmith.com)

A commandline client to interact with Curseforge. For now it's mostly used to
fetch mods defined within modpack manifests.

## Install

You can download prebuilt binaries from the [GitHub releases][releases] or from
our [download site][downloads]. Besides that we also prepared repositories for
DEB and RPM packages which can be found at [Cloudsmith][pkgrepo]. If you prefer
to use containers you could use our images published on [GHCR][ghcr],
[Docker Hub][dockerhub] or [Quay][quay]. If you need further guidance how to
install this take a look at our [documentation][docs].

Package repository hosting is graciously provided by [Cloudsmith][cloudsmith].
Cloudsmith is the only fully hosted, cloud-native, universal package management
solution, that enables your organization to create, store and share packages in
any format, to any place, with total confidence.

## Prerequisites

We use [mise][mise] to manage all required tools and their versions. Install it
by following the [official installation instructions][mise-install], then run
the following commands inside the repository to activate mise and install all
tools defined in `mise.toml`:

```console
mise trust
mise install
```

## Build

Since all required commands ar part of our [go-task][gotask] taskfile the
commands you got to execute are quite simple:

```console
git clone https://github.com/webhippie/cursecli.git
cd cursecli

task build
./bin/cursecli -h
```

## Development

To start developing on this project you have to execute only a few commands in
multiple terminal tabs or windows:

```console
task watch
```

After that you can simply execute the tool via `bin/cursecli -h`. Generally it
supports hot reloading which means the binary gets automatically recompiled on
code changes.

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
[downloads]: https://dl.webhippie.de/#cursecli/
[ghcr]: https://github.com/webhippie/cursecli/pkgs/container/cursecli
[dockerhub]: https://hub.docker.com/r/webhippie/cursecli/tags/
[quay]: https://quay.io/repository/webhippie/cursecli?tab=tags
[docs]: https://webhippie.github.io/cursecli/#getting-started
[pkgrepo]: https://cloudsmith.io/~webhippie/repos/general/groups/
[cloudsmith]: https://cloudsmith.com/
[gotask]: https://taskfile.dev/installation/
[mise]: https://mise.jdx.dev/
[mise-install]: https://mise.jdx.dev/getting-started.html
[commits]: https://www.conventionalcommits.org/en/v1.0.0/
[semver]: https://semver.org/
