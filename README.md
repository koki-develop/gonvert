# gat

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/koki-develop/gonvert)](https://github.com/koki-develop/gonvert/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/koki-develop/gonvert/ci.yml?logo=github)](https://github.com/koki-develop/gonvert/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/koki-develop/gonvert)](https://goreportcard.com/report/github.com/koki-develop/gonvert)
[![LICENSE](https://img.shields.io/github/license/koki-develop/gonvert)](./LICENSE)

Convert between JSON, YAML

- [Installation](#installation)
  - [Homebrew](#homebrew)
  - [`go install`](#go-install)
  - [Releases](#releases)
- [Usage](#usage)
  - [JSON to YAML](#json-to-yaml)
  - [YAML to JSON](#yaml-to-json)
- [LICENSE](#license)

## Installation

### Homebrew

```console
$ brew install koki-develop/tap/gonvert
```

### `go install`

```console
$ go install github.com/koki-develop/gonvert@latest
```

### Releases

Download the binary from the [releases page](https://github.com/koki-develop/gonvert/releases/latest).

## Usage

```console
$ gonvert --help
Convert between JSON, YAML

Usage:
  gonvert [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  json2yaml   Convert JSON to YAML
  yaml2json   Convert YAML to JSON

Flags:
  -h, --help   help for gonvert

Use "gonvert [command] --help" for more information about a command.
```

### JSON to YAML

```console
$ gonvert json2yaml --help
Convert JSON to YAML.

Usage:
  gonvert json2yaml [flags]

Aliases:
  json2yaml, jsontoyaml, json2yml, jsontoyml, j2y, jtoy, jy

Flags:
  -h, --help   help for json2yaml
```

### YAML to JSON

```console
$ gonvert yaml2json --help
Convert YAML to JSON.

Usage:
  gonvert yaml2json [flags]

Aliases:
  yaml2json, yamltojson, yml2json, ymltojson, y2j, ytoj, yj

Flags:
  -h, --help   help for yaml2json
```

## LICENSE

[MIT](./LICENSE)
