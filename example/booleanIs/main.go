// Package main implements a simple example of using szargs.
package main

import (
	"fmt"
	"os"

	"github.com/dancsecs/szargs"
)

func main() {
	args := szargs.New(
		"A simple utility to demo identifying a boolean flags.",
		os.Args,
	)

	isTrue := args.Is(
		"[-t | --true]", // A short and long form.
		"A single flag indicating true.",
	)

	args.Done() // All arguments should have consumed.

	if isTrue {
		_, _ = fmt.Fprintln(os.Stdout, "Is true.")
	} else {
		_, _ = fmt.Fprintln(os.Stdout, "Is NOT true.")
	}

	if args.HasError() {
		_, _ = fmt.Fprintln(os.Stderr, "Error: "+args.Err().Error())
	}
}
