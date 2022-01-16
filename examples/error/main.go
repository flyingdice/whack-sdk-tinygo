package main

import (
	"errors"
	"github.com/flyingdice/whack-tinygo-guest-sdk/whack"
)

type errorApp struct{}

func (a *errorApp) Main(req *whack.Request) *whack.Response {
	return whack.Error(errors.New("oh no! an error occurred in your wasm!"))
}

func main() {
	whack.Register(&errorApp{})
}
