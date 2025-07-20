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

# Example Flagged Values


## Overview

This example demos the collecting of a flag's values into a typed slice.  If
found the matching flags and arguments are removed as encountered. If a
specialized value is named then it must pass rules regarding its syntax and
range.

It demonstrates the following szargs functions:

<!--- gotomd::Bgn::dcln::./../../New Args.ValuesString Args.ValuesUint8 Args.Done Args.HasErr Args.Err Args.Usage -->
```go
// New creates a new Args object based in the arguments passed.  The first
// element of the arguments must be the program name.
func New(programDesc string, args []string) *Args

// ValuesString scans for repeated instances of the specified flag and
// captures the following values as a slice of strings. The flags and values
// are removed from the argument list.
// 
// If any instance of the flag lacks a following value, an error is
// registered.
// 
// Returns a slice of the captured string values.
func (args *Args) ValuesString(flag, desc string) []string

// ValuesUint8 scans for repeated instances of the specified flag and parses
// the following values as unsigned 8 bit integers. The flags and values are
// removed from the argument list.
// 
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for a uint8, an error is registered.
// 
// Returns a slice of the parsed uint8 values.
func (args *Args) ValuesUint8(flag, desc string) []uint8

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
<!--- gotomd::End::dcln::./../../New Args.ValuesString Args.ValuesUint8 Args.Done Args.HasErr Args.Err Args.Usage -->

## Contents

- [Source (main.go)](#source)
    - [Example: **PASS**: No Arguments](#pass-no-arguments)
    - [Example: **PASS**: Single Long Form](#pass-single-long-form)
    - [Example: **PASS**: Single Short Form](#pass-single-short-form)
    - [Example: **PASS**: Multiple Mixed Forms](#pass-multiple-mixed-forms)
    - [Example: **FAIL**: Extra Unknown Argument](#fail-extra-unknown-argument)

## Source

The source used for this example.  It simply defines a new szargs.Arg object,
then defines two flagged argument collections.  Finally the Done function is
called to insure that no more arguments exist in the list.

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
```
<!--- gotomd::End::file::./main.go -->

[Top of Page](#example-flagged-values) --
[Szargs Contents](../../README.md#contents)

### **PASS**: No Arguments

<!--- gotomd::Bgn::run::./. -->
---
```bash
go run .
```

<pre>
0 Name(s) Found: [].
0 Byte(s) Found: [].
</pre>
---
<!--- gotomd::End::run::./. -->

[Top of Page](#example-flagged-values) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Single Long Form

<!--- gotomd::Bgn::run::./. --name theName --byte 23 -->
---
```bash
go run . --name theName --byte 23
```

<pre>
1 Name(s) Found: [theName].
1 Byte(s) Found: [23].
</pre>
---
<!--- gotomd::End::run::./. --name theName --byte 23 -->

[Top of Page](#example-flagged-values) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Single Short Form

<!--- gotomd::Bgn::run::./. -n anotherName -b 42 -->
---
```bash
go run . -n anotherName -b 42
```

<pre>
1 Name(s) Found: [anotherName].
1 Byte(s) Found: [42].
</pre>
---
<!--- gotomd::End::run::./. -n anotherName -b 42 -->

[Top of Page](#example-flagged-values) --
[Szargs Contents](../../README.md#contents)


### **PASS**: Multiple Mixed Forms

<!--- gotomd::Bgn::run::./. -n aName --name anotherName -b 42 --byte 23 -->
---
```bash
go run . -n aName --name anotherName -b 42 --byte 23
```

<pre>
2 Name(s) Found: [aName anotherName].
2 Byte(s) Found: [42 23].
</pre>
---
<!--- gotomd::End::run::./. -n aName --name anotherName -b 42 --byte 23 -->

[Top of Page](#example-flagged-values) --
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

values
A simple demo of values flag.

Usage: values [-n | --name] [-b | --byte]

[-n | --name]
The name string for the values.

[-b | --byte]
The byte (0-255) for the values.
</pre>
---
<!--- gotomd::End::run::./. extraUnknownArgument -->

[Top of Page](#example-flagged-values) --
[Szargs Contents](../../README.md#contents)
