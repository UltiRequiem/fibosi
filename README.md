# Fibonacci

[![GitMoji](https://img.shields.io/badge/Gitmoji-%F0%9F%8E%A8%20-FFDD67.svg)](https://gitmoji.dev)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/UltiRequiem/fibonacci)](https://pkg.go.dev/github.com/UltiRequiem/fibonacci)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Lines Of Code](https://img.shields.io/tokei/lines/github.com/UltiRequiem/fibonnaci?color=blue&label=Total%20Lines)
![CodeQL](https://github.com/UltiRequiem/fibonnaci/workflows/CodeQL/badge.svg)
![Build](https://github.com/UltiRequiem/fibonnaci/workflows/Build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/fibonnaci)](https://goreportcard.com/report/github.com/UltiRequiem/fibonnaci)

## PKG

Install the package by:

```bash
go get github.com/UltiRequiem/fibonacci/pkg
```

This packages exposes two functions, [`Fibonacci`](https://github.com/UltiRequiem/fibonacci/blob/main/pkg/root.go#L8) and [`FibonacciSequence`](https://github.com/UltiRequiem/fibonacci/blob/main/pkg/root.go#L34).

Usage Example:

```go
package main

import (
	"fmt"

	fina "github.com/UltiRequiem/fibonacci/pkg"
)

func main() {
	fiboNum, _ := fina.Fibonacci(9)
	fmt.Println(fiboNum) // 34

	fiboSequence, _ := fina.FibonacciSequence(9)
	fmt.Println(fiboSequence) // [0 1 1 2 3 5 8 13 21]
}
```

For more examples, see [internal](./internal/) directory.

Check for more detailed info on [pkg.go.dev](https://pkg.go.dev/github.com/UltiRequiem/fibonnaci/pkg).

## API

If you want to try the API, there are a few options:

- Install as a CLI Program...

```
go install github.com/UltiRequiem/fibonacci@latest
```

If you have correctly configured your [GOPATH](https://golang.org/doc/gopath_code#GOPATH), then:

```bash
fibonacci -p 8080
```

- Using Binary from Releases

You can get the build of the project in [releases](https://github.com/UltiRequiem/fibonacci/releases/latest).

- From Source

`git clone` this project or [download it in zip format](https://github.com/UltiRequiem/fibonacci/archive/refs/heads/main.zip).

```bash
git clone https://github.com/UltiRequiem/fibonacci
```

Then run [`main.go`](./main.go) file.

#### License

This project is licensed under the [MIT license](./LICENSE.md).
