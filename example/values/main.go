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
		"A simple demo of values flag.",
		os.Args,
	)

	nameList := args.ValuesString(
		"[-n | --name]",
		"The name string for the values.",
	)

	byteList := args.ValuesUint8(
		"[-b | --byte]",
		"The byte (0-255) for the values.",
	)

	args.Done() // All arguments should have consumed.

	if args.HasErr() {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage())
	} else {
		fmt.Printf("%d Name(s) Found: %v.\n", len(nameList), nameList)
		fmt.Printf("%d Byte(s) Found: %v.\n", len(byteList), byteList)
	}
}
