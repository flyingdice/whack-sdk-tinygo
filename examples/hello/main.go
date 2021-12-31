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
