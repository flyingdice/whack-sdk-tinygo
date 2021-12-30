package whack

import (
	"fmt"
	"os"
)

// TODO (ahawker)
// Create logger iface bound to the request that allows 3rd party code
// to log via host function import.

func stdout(a ...interface{}) {
	_, _ = fmt.Fprint(os.Stdout, a...)
}

func stderr(a ...interface{}) {
	_, _ = fmt.Fprint(os.Stderr, a...)
}
