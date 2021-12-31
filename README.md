# Whack TinyGo Guest SDK

[Whack](https://github.com/flyingdice/whack) guest SDK implementation for [TinyGo](https://tinygo.org/).

## Getting Started

```
go get "github.com/flyingdice/whack-tinygo-guest-sdk"
```

Follow the simple snippet below or the [examples](examples) code for a more in-depth look.

```
package main

import (
	"github.com/flyingdice/whack-tinygo-guest-sdk/whack"
)

type helloApp struct{}

func (a *helloApp) Main(req *whack.Request) *whack.Response {
	return whack.Success([]byte("Hello World"))
}

func main() {
	whack.Register(&helloApp{})
}
```

## Usage

Run `make help` to see list of available commands.

```
❯ make help
modules                        Tidy up and vendor go modules.
help                           Print Makefile usage.
```

## License

[AGPL v3](LICENSE)
