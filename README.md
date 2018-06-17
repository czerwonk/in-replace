# in-replace
[![Build Status](https://travis-ci.org/czerwonk/in-replace.svg)](https://travis-ci.org/czerwonk/in-replace)
[![Go Report Card](https://goreportcard.com/badge/github.com/czerwonk/in-replace)](https://goreportcard.com/report/github.com/czerwonk/in-replace)

simple multi platform inplace replacer

## Installation
```bash
go get github.com/czerwonk/in-replace
```

## Configuration
to configure in-replace a yaml file of the following form is needed:

```yaml
files:
  - path: "test.txt"
    replacements:
      - regex: "(t)est"
        replacement: "T"
        group: 1
```

group is the capture group and can be omited if the whole match is meant to be replaced.

## License
(c) Daniel Czerwonk, 2018. Licensed under [MIT](LICENSE) license.
