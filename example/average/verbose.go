package main

import (
	"fmt"
)

// Supported verbose levels.
const (
	vAll = iota
	vLv1
	vLv2
	vLv3
)

var verbose int //nolint:goCheckNoGlobals // Ok.

func sayPrintf(minLevel int, msgFormat string, msgArgs ...any) {
	if verbose >= minLevel {
		fmt.Printf(msgFormat, msgArgs...) //nolint:forbidigo // Ok.
	}
}
