# Development Guide

## Prerequisites
Install the necessary tools for development by following their respective installation instructions.

- [Go](https://go.dev/doc/install)
- [Mage](https://magefile.org/)

## Build
After making changes to the Go source code, build the project with the following command:

```shell
$ mage build
$ ./trivy -h
```

## Lint
You must pass the linter checks:

```shell
$ mage lint
```

Additionally, you need to have run `go mod tidy`, so execute the following command as well:

```shell
$ mage tidy
```

## Unit tests
Your PR must pass all the unit tests. You can test it as below.

```
$ mage test:unit
```

## Integration tests
Your PR must pass all the integration tests. You can test it as below.

```
$ mage test:integration
```

## Documentation
If you update CLI flags, you need to generate the CLI references.
The test will fail if they are not up-to-date.

```shell
$ mage docs:generate
```

You can build the documents as below and view it at http://localhost:8000.

```
$ mage docs:serve
```
