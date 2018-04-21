# negotiate [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

Negotiate is a simple HTTP content type negotiation package. Since its core
dependency doesn't tag releases, it checks in `vendor/` and doesn't leak any
vendored types.

## Installation

```
go get -u github.com/akshayjshah/vendor
```

## Current Status

Stable. Once https://github.com/golang/go/issues/19307 is resolved, I'll stop
maintaining this package.

[doc-img]: https://godoc.org/github.com/akshayjshah/negotiate?status.svg
[doc]: https://godoc.org/github.com/akshayjshah/negotiate
[ci-img]: https://travis-ci.org/akshayjshah/negotiate.svg?branch=master
[ci]: https://travis-ci.org/akshayjshah/negotiate
[cov-img]: https://codecov.io/gh/akshayjshah/negotiate/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/akshayjshah/negotiate
