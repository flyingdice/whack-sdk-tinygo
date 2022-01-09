package main

import (
	"github.com/flyingdice/whack-tinygo-guest-sdk/whack"
)

type echoApp struct{}

func (a *echoApp) Main(req *whack.Request) *whack.Response {
	return whack.Success(req.Bytes)
}

func main() {
	whack.Register(&echoApp{})
}
