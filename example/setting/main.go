// Package main implements a simple example of using szargs.
//
//nolint:forbidigo // OK to print to os.Stdout.
package main

import (
	"fmt"
	"os"

	"github.com/dancsecs/szargs"
)

func mainSetEnv() {
	_ = os.Setenv("ENV_SZARGS_EXAMPLE_SETTING_TEMP", "f")

	main()
}

func main() {
	args := szargs.New(
		"A simple demo of a setting.",
		os.Args,
	)

	temp := args.SettingOption(
		"[-t | --temp {c,f}]",
		"ENV_SZARGS_EXAMPLE_SETTING_TEMP",
		"c",
		[]string{
			"c",
			"f",
		},
		"Temperature measurement to use (celsius or fahrenheit).",
	)

	args.Done() // All arguments should have consumed.

	if args.HasErr() {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage())
	} else {
		fmt.Printf("Using '%s' for temperatures.", temp)
	}
}
