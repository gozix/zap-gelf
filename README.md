# GoZix Zap Gelf

[documentation-img]: https://img.shields.io/badge/godoc-reference-blue.svg?color=24B898&style=for-the-badge&logo=go&logoColor=ffffff
[documentation-url]: https://pkg.go.dev/github.com/gozix/zap-gelf/v2
[license-img]: https://img.shields.io/github/license/gozix/zap-gelf.svg?style=for-the-badge
[license-url]: https://github.com/gozix/zap-gelf/blob/master/LICENSE
[release-img]: https://img.shields.io/github/tag/gozix/zap-gelf.svg?label=release&color=24B898&logo=github&style=for-the-badge
[release-url]: https://github.com/gozix/zap-gelf/releases/latest
[build-status-img]: https://img.shields.io/github/actions/workflow/status/gozix/zap-gelf/go.yml?logo=github&style=for-the-badge
[build-status-url]: https://github.com/gozix/zap-gelf/actions
[go-report-img]: https://img.shields.io/badge/go%20report-A%2B-green?style=for-the-badge
[go-report-url]: https://goreportcard.com/report/github.com/gozix/zap-gelf
[code-coverage-img]: https://img.shields.io/codecov/c/github/gozix/zap-gelf.svg?style=for-the-badge&logo=codecov
[code-coverage-url]: https://codecov.io/gh/gozix/zap-gelf

[![License][license-img]][license-url]
[![Documentation][documentation-img]][documentation-url]

[![Release][release-img]][release-url]
[![Build Status][build-status-img]][build-status-url]
[![Go Report Card][go-report-img]][go-report-url]
[![Code Coverage][code-coverage-img]][code-coverage-url]

The bundle provide a Zap GELF core to GoZix application.

## Installation

```shell
go get github.com/gozix/zap-gelf/v2
```

## Dependencies

* [viper](https://github.com/gozix/viper)
* [zap](https://github.com/gozix/zap-gelf)

## Configuration example

```json
{
  "zap": {
    "cores": {
      "gelf": {
        "type": "gelf",
        "addr": "127.0.0.1:12001",
        "level": "debug"
      }
    }
  }
}
```

## Documentation

You can find documentation on [pkg.go.dev][documentation-url] and read source code if needed.

## Questions

If you have any questions, feel free to create an issue.