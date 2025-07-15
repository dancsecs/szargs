// Package main implements a simple example of using szargs.
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
		"[-c || --count ...]", // Short and long forms both repeatable.
		"How many times?)",
	)

	args.Done() // All arguments should have consumed.

	_, _ = fmt.Fprintf(os.Stdout, "How many: %d\n", howMany)

	if args.HasError() {
		_, _ = fmt.Fprintln(os.Stderr, "Error: "+args.Err().Error())
	}
}
