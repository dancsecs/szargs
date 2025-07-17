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
		"A simple utility to demo identifying a boolean flag.",
		os.Args,
	)

	isTrue := args.Is(
		"[-t | --true]", // A short and long form.
		"A single flag indicating true.",
	)

	args.Done() // All arguments should have consumed.

	if args.HasError() {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage())
	} else {
		if isTrue {
			fmt.Println("Is true.")
		} else {
			fmt.Println("Is NOT true.")
		}
	}
}
