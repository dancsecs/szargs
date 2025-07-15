// Package main implements a simple example of using szargs.
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

	if nameFound {
		_, _ = fmt.Fprintf(os.Stdout, "Name Found: %s.\n", name)
	} else {
		_, _ = fmt.Fprintln(os.Stdout, "Name Not Found.")
	}

	if numFound {
		_, _ = fmt.Fprintf(os.Stdout, "Byte Found: %d.\n", num)
	} else {
		_, _ = fmt.Fprintln(os.Stdout, "Byte Not Found.")
	}

	if args.HasError() {
		_, _ = fmt.Fprintln(os.Stderr, "Error: "+args.Err().Error())
	}
}
