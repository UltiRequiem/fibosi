# Fibonacci

[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/fibonacci)](https://goreportcard.com/report/github.com/UltiRequiem/fibonnaci)

High Performance Fibonacci Abstraction Layer and an API.

## PKG

Install the package by:

```bash
go get github.com/UltiRequiem/fibonacci/pkg
```

This packages exposes two functions,
[`Fibonacci`](https://github.com/UltiRequiem/fibonacci/blob/main/pkg/root.go#L8)
and
[`FibonacciSequence`](https://github.com/UltiRequiem/fibonacci/blob/main/pkg/root.go#L34).

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

Check for more detailed info on
[pkg.go.dev](https://pkg.go.dev/github.com/UltiRequiem/fibonacci/pkg).

## API

API Demo: https://z91bf2703-za30e6095-gtw.qovery.io

### Endpoints

- Root Path

Try `curl http://localhost:3000/890`

- Sequence Path

Try `curl http://localhost:3000/sequence/890`

### Run The API

- Install as a CLI Program...

```sh
go install github.com/UltiRequiem/fibonacci@latest
```

If you have correctly configured your
[GOPATH](https://golang.org/doc/gopath_code#GOPATH), then:

```bash
fibonacci -p 8080
```

- Using Binary from Releases

You can get the build of the project in
[releases](https://github.com/UltiRequiem/fibonacci/releases/latest).

- From Source

`git clone` this project or
[download it in zip format](https://github.com/UltiRequiem/fibonacci/archive/refs/heads/main.zip).

```bash
git clone https://github.com/UltiRequiem/fibonacci
```

Then run [`main.go`](./main.go) file.

## Support

Open an Issue, I will check it a soon as possible 👀

If you want to hurry me up a bit
[send me a tweet](https://twitter.com/UltiRequiem) 😆

Consider [supporting me on Patreon](https://patreon.com/UltiRequiem) if you like
my work 🙏

Don't forget to start the repo ⭐

## Authors

[Eliaz Bobadilla](https://ultirequiem.com) - Creator and Maintainer 💪

See also the full list of
[contributors](https://github.com/UltiRequiem/fibonacci/contributors) who
participated in this project ✨

## Licence

Licensed under the MIT License 📄
