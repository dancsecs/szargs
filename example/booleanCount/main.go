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
		"A simple utility to demo counting boolean flags.",
		os.Args,
	)

	howMany := args.Count(
		"[-c | --count ...]", // Short and long forms both repeatable.
		"How many times?",
	)

	args.Done() // All arguments should have consumed.

	if args.HasErr() {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage())
	} else {
		fmt.Printf("How many: %d\n", howMany)
	}
}
