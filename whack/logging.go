package whack

import (
	"fmt"
	"os"
)

// TODO (ahawker)
// Create logger iface bound to the request that allows 3rd party code
// to log via host function import.

func stdout(s string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, s, a...)
}

func stderr(s string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, s, a...)
}
