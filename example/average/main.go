// Package main implements a simple example of using a flag to provide a list
// of floating point numbers and returning the sum or average of the list.
package main

import (
	"fmt"
	"os"

	"github.com/dancsecs/szargs"
)

const (
	verboseFlag = "[-v | --verbose ...]"
	verboseDesc = "The verbose level."

	numberFlag = "[-n | --number float64 ...]"
	numberDesc = "The numbers to act on."

	operationName = "[operation]"
	operationDesc = "The operation (add or average) defaulting to add."
)

// Example function being tested.
func main() {
	args := szargs.New(
		"A simple utility to add or average a number list.",
		os.Args,
	)

	// How chatty should I be.  Set the global value.
	verbose = args.Count(verboseFlag, verboseDesc)

	// Gather all of the number to operate on.
	numbers := args.ValuesFloat64(numberFlag, numberDesc)

	// Set to default if not present.
	if !args.HasNext() {
		args.PushArg("add")
	}

	operation := args.NextOption(
		operationName,
		[]string{
			"add",
			"average",
		},
		operationDesc,
	)

	args.Done()

	// Report parsing errors or process the arguments.
	if args.Err() != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage())
	} else {
		process(numbers, operation)
	}
}
