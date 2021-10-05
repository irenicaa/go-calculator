# go-calculator

[![GoDoc](https://godoc.org/github.com/irenicaa/go-calculator?status.svg)](https://godoc.org/github.com/irenicaa/go-calculator)
[![Go Report Card](https://goreportcard.com/badge/github.com/irenicaa/go-calculator)](https://goreportcard.com/report/github.com/irenicaa/go-calculator)
[![Build Status](https://app.travis-ci.com/irenicaa/go-calculator.svg?branch=master)](https://app.travis-ci.com/irenicaa/go-calculator)
[![codecov](https://codecov.io/gh/irenicaa/go-calculator/branch/master/graph/badge.svg)](https://codecov.io/gh/irenicaa/go-calculator)

The simplified clone of the Unix bc tool.

## Installation

```
$ go get github.com/irenicaa/go-calculator/...
```

## Usage

```
$ go-calculator -h | -help | --help
```

Stdin: code (see [docs](docs/) for details).

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit.

## Docs

[Docs](docs/)

## Output Example

```
// https://en.wikipedia.org/wiki/Gauss-Legendre_algorithm

// initial values
a0 = 1
1
b0 = 1/sqrt(2)
0.7071067811865475
t0 = 1/4
0.25
p0 = 1
1

// first iteration
a1 = (a0 + b0) / 2
0.8535533905932737
b1 = sqrt(a0 * b0)
0.8408964152537145
t1 = t0 - p0 * (a0 - a1)^2
0.22855339059327376
p1 = 2*p0
2

// second iteration
a2 = (a1 + b1) / 2
0.8472249029234942
b2 = sqrt(a1 * b1)
0.8472012667468914
t2 = t1 - p1 * (a1 - a2)^2
0.22847329108090064
p2 = 2*p1
4

pi = (a2 + b2)^2 / (4*t2)
3.141592646213543
```

## License

The MIT License (MIT)

Copyright &copy; 2020-2021 irenica
