<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

<!---
   Szerszam argument library: szargs.
   Copyright (C) 2024  Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
-->

# Example Setting


## Overview

This example demos the use of the setting function where a variable has a
default value that can be overridden by an environment variable which can be
overridden by a command line argument.

It demonstrates the following szargs functions:

It demonstrates the following szargs functions:

<!--- gotomd::Bgn::dcln::./../../New Args.NextString Args.HasNext Args.PushArg Args.NextUint Args.Done Args.HasErr Args.Err Args.Usage -->
```go
// New creates a new Args object based in the arguments passed.  The first
// element of the arguments must be the program name.
func New(programDesc string, args []string) *Args

// NextString removes and returns the next argument from the argument list.
// 
// If no arguments remain, an error is registered.
// 
// Returns the next argument value as a string.
func (args *Args) NextString(name, desc string) string

// HasNext returns true if any arguments remain unabsorbed.
func (args *Args) HasNext() bool

// PushArg places the supplied argument to the end of the internal args list.
func (args *Args) PushArg(arg string)

// NextUint removes and returns the next argument from the argument list,
// parsing it as an unsigned integer.
// 
// If no arguments remain, or if the value has invalid syntax or is out of
// range for a uint, an error is registered.
// 
// Returns the next argument value parsed as a uint.
func (args *Args) NextUint(name, desc string) uint

// Done registers an error if there are any remaining arguments.
func (args *Args) Done()

// HasErr returns true if any errors have been encountered or registered.
func (args *Args) HasErr() bool

// Err returns any errors encountered or registered while parsing the
// arguments.
func (args *Args) Err() error

// Usage returns a usage message based on the parsed arguments.
func (args *Args) Usage() string
```
<!--- gotomd::End::dcln::./../../New Args.NextString Args.HasNext Args.PushArg Args.NextUint Args.Done Args.HasErr Args.Err Args.Usage -->

## Contents

- [Source (main.go)](#source)
    - [Example: **PASS**: Use Default](#pass-use-default)
    - [Example: **PASS**: Use Env](#pass-use-env)
    - [Example: **PASS**: Use Flag](#pass-use-flag)
    - [Example: **FAIL**: Extra Unknown Argument](#fail-extra-unknown-argument)

## Source

The source used for this example.  It simply defines a new szargs.Arg object,
then defines a setting option. Finally the Done function is called to insure
that no more arguments exist in the list.

<!--- gotomd::Bgn::file::./main.go -->
```bash
cat ./main.go
```

```go
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
```
<!--- gotomd::End::file::./main.go -->

[Top of Page](#example-setting) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Use Default

<!--- gotomd::Bgn::run::./.  -->
---
```bash
go run .
```

<pre>
Using 'c' for temperatures.
</pre>
---
<!--- gotomd::End::run::./.  -->

[Top of Page](#example-setting) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Use Env

---
```bash
go run ./mainSetEnv
```

<pre>
Using 'f' for temperatures.
</pre>

[Top of Page](#example-setting) --
[Szargs Contents](../../README.md#contents)


### **PASS**: Use Flag

<!--- gotomd::Bgn::run::./. -t f -->
---
```bash
go run . -t f
```

<pre>
Using 'f' for temperatures.
</pre>
---
<!--- gotomd::End::run::./. -t f -->

[Top of Page](#example-setting) --
[Szargs Contents](../../README.md#contents)


### **FAIL**: Extra Unknown Argument

The error is generated by the ```args.Done()``` statement causing both the
error and a usage statement to be returned.

<!--- gotomd::Bgn::run::./. extraUnknownArgument -->
---
```bash
go run . extraUnknownArgument
```

<pre>
Error: unexpected argument: [extraUnknownArgument]

setting
A simple demo of a setting.

Usage: setting [-t | --temp {c,f}]

[-t | --temp {c,f}]
Temperature measurement to use (celsius or fahrenheit).
</pre>
---
<!--- gotomd::End::run::./. extraUnknownArgument -->

[Top of Page](#example-setting) --
[Szargs Contents](../../README.md#contents)
