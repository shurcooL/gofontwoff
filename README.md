gofontwoff
==========

[![Go Reference](https://pkg.go.dev/badge/github.com/shurcooL/gofontwoff.svg)](https://pkg.go.dev/github.com/shurcooL/gofontwoff)

Package gofontwoff provides the Go font family in Web Open Font Format.

It's a Go package that statically embeds Go font family WOFF data,
exposing it via an [`http.FileSystem`](https://go.dev/pkg/net/http#FileSystem).

These fonts were created by the Bigelow & Holmes foundry specifically
for the Go project. See https://go.dev/blog/go-fonts for details.

Installation
------------

```sh
go get github.com/shurcooL/gofontwoff
```

Alternatives
------------

- [gofontweb](https://artyom.dev/gofontweb) - Embeddable Go fonts in woff2 format, exposed via the [`fs.FS`](https://go.dev/pkg/io/fs#FS) interface.

License
-------

-	[BSD-style License](LICENSE)
