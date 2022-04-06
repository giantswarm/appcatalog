[![GoDoc](https://godoc.org/github.com/giantswarm/appcatalog?status.svg)](http://godoc.org/github.com/giantswarm/appcatalog)
[![CircleCI](https://circleci.com/gh/giantswarm/appcatalog.svg?style=shield)](https://circleci.com/gh/giantswarm/appcatalog)

# appcatalog

This repo contains 2 components that support the Giant Swarm App Platform.

- A Helm chart used by opsctl ensure appcatalogs.
- A Go library we use in integration tests to get the correct version of
managed apps.

## Getting Project

Clone the git repository: https://github.com/giantswarm/appcatalog.git

### How to build

Build it using the standard `go build` command.

```
go build github.com/giantswarm/appcatalog
```

## Contact

- Mailing list: [giantswarm](https://groups.google.com/forum/!forum/giantswarm)
- IRC: #[giantswarm](irc://irc.freenode.org:6667/#giantswarm) on freenode.org
- Bugs: [issues](https://github.com/giantswarm/appcatalog/issues)

## Contributing & Reporting Bugs

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches, the
contribution workflow as well as reporting bugs.

## License

appcatalog is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for
details.
