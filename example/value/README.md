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

# Example Flagged Value


## Overview

This example demos the setting of a flag's value.  If found the matching flag
and its argument are removed as encountered.  An ambiguous error will be
generated if the flag appears more than once.  If a specialized value is named
then it must pass rules regarding its syntax and range.

It demonstrates the following szargs functions:

<!--- gotomd::Bgn::dcln::./../../New Args.ValueString Args.ValueUint8 Args.Done Args.HasErr Args.Err Args.Usage -->
```go
// New creates a new Args object based in the arguments passed.  The first
// element of the arguments must be the program name.
func New(programDesc string, args []string) *Args

// ValueString scans for a specific flagged argument and captures its
// following value as a string. The flag and its value are removed from the
// argument list.
// 
// If the flag appears more than once or lacks a following value, an error is
// registered.
// 
// Returns the string value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueString(flag, desc string) (string, bool)

// ValueUint8 scans for a specific flagged argument and parses its value as an
// unsigned 8 bit integer. The flag and its value are removed from the
// argument list.
// 
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for a uint8, an error is
// registered.
// 
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueUint8(flag, desc string) (uint8, bool)

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
<!--- gotomd::End::dcln::./../../New Args.ValueString Args.ValueUint8 Args.Done Args.HasErr Args.Err Args.Usage -->

## Contents

- [Source (main.go)](#source)
    - [Example: **PASS**: No Arguments](#pass-no-arguments)
    - [Example: **PASS**: Single Long Form](#pass-single-long-form)
    - [Example: **PASS**: Single Short Form](#pass-single-short-form)
    - [Example: **FAIL**: Ambiguous Argument](#fail-ambiguous-argument)
    - [Example: **FAIL**: Extra Unknown Argument](#fail-extra-unknown-argument)

## Source

The source used for this example.  It simply defines a new szargs.Arg object,
then defines two flagged arguments.  Finally the Done function is called to
insure that no more arguments exist in the list.

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
```
<!--- gotomd::End::file::./main.go -->

[Top of Page](#example-flagged-value) --
[Szargs Contents](../../README.md#contents)

### **PASS**: No Arguments

<!--- gotomd::Bgn::run::./. -->
---
```bash
go run .
```

<pre>
Name Not Found.
Byte Not Found.
</pre>
---
<!--- gotomd::End::run::./. -->

[Top of Page](#example-flagged-value) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Single Long Form

<!--- gotomd::Bgn::run::./. --name theName --byte 23 -->
---
```bash
go run . --name theName --byte 23
```

<pre>
Name Found: theName.
Byte Found: 23.
</pre>
---
<!--- gotomd::End::run::./. --name theName --byte 23 -->

[Top of Page](#example-flagged-value) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Single Short Form

<!--- gotomd::Bgn::run::./. -n anotherName -b 42 -->
---
```bash
go run . -n anotherName -b 42
```

<pre>
Name Found: anotherName.
Byte Found: 42.
</pre>
---
<!--- gotomd::End::run::./. -n anotherName -b 42 -->

[Top of Page](#example-flagged-value) --
[Szargs Contents](../../README.md#contents)


### **FAIL**: Ambiguous Argument

The error is because a true flag appeared more than once.  A second error is
generated by the ```args.Done()``` statement causing both the error and a
usage statement to be returned.

<!--- gotomd::Bgn::run::./. --name first -n second --byte 1 -b 2 -->
---
```bash
go run . --name first -n second --byte 1 -b 2
```

<pre>
Error: ambiguous argument: '[-n | --name]' for 'second' already set to: 'first': ambiguous argument: '[-b | --byte]' for '2' already set to: '1'

value
A simple demo of value flags.

Usage: value [-n | --name] [-b | --byte]

[-n | --name]
The name string value.

[-b | --byte]
The byte (0-255) value.
</pre>
---
<!--- gotomd::End::run::./. --name first -n second --byte 1 -b 2 -->

[Top of Page](#example-flagged-value) --
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

value
A simple demo of value flags.

Usage: value [-n | --name] [-b | --byte]

[-n | --name]
The name string value.

[-b | --byte]
The byte (0-255) value.
</pre>
---
<!--- gotomd::End::run::./. extraUnknownArgument -->

[Top of Page](#example-flagged-value) --
[Szargs Contents](../../README.md#contents)
