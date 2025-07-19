// Package main implements a simple example of using szargs.
//
//nolint:forbidigo // OK to print to os.Stdout.
package main

import (
	"fmt"
	"os"

	"github.com/dancsecs/szargs"
)

func main() {
	args := szargs.New(
		"A simple demo of repeating a string.",
		os.Args,
	)

	mandatory := args.NextString(
		"message",
		"What to repeat.",
	)

	if !args.HasNext() {
		args.PushArg(("3")) // Default value dor optional integer.
	}

	optional := args.NextUint(
		"[times]",
		"The number of times to repeat.  Defaults to 3.",
	)

	args.Done() // All arguments should have consumed.

	if args.HasErr() {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage())
	} else {
		for range optional {
			fmt.Println(mandatory)
		}
	}
}
