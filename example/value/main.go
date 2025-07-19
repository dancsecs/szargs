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
		"A simple demo of value flags.",
		os.Args,
	)

	name, nameFound := args.ValueString(
		"[-n | --name]",
		"The name string value.",
	)

	num, numFound := args.ValueUint8(
		"[-b | --byte]",
		"The byte (0-255) value.",
	)

	args.Done() // All arguments should have consumed.

	if args.HasErr() { //nolint:nestif // Ok for the demo.
		fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage())
	} else {
		if nameFound {
			fmt.Printf("Name Found: %s.\n", name)
		} else {
			fmt.Println("Name Not Found.")
		}

		if numFound {
			fmt.Printf("Byte Found: %d.\n", num)
		} else {
			fmt.Println("Byte Not Found.")
		}
	}
}
