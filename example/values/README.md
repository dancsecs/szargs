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

# Example Flag Values


## Overview

This example demos the collecting of a flag's values into a typed slice.  If
found the matching flags and arguments are removed as encountered. If a
specialized value is named then it must pass rules regarding its syntax and
range.

<!--- gotomd::Bgn::dcln::./../../Args.ValuesString -->
```go
// ValuesString scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesString(flag, desc string) []string
```
<!--- gotomd::End::dcln::./../../Args.ValuesString -->

and its specific versions:

<!--- gotomd::Bgn::dcls::./../../Args.ValuesFloat64 Args.ValuesFloat32 Args.ValuesInt64 Args.ValuesInt32 Args.ValuesInt16 Args.ValuesInt8 Args.ValuesInt Args.ValuesUint64 Args.ValuesUint32 Args.ValuesUint16 Args.ValuesUint8 Args.ValuesUint Args.ValuesOption -->
```go
func (args *Args) ValuesFloat64(flag, desc string) []float64
func (args *Args) ValuesFloat32(flag, desc string) []float32
func (args *Args) ValuesInt64(flag, desc string) []int64
func (args *Args) ValuesInt32(flag, desc string) []int32
func (args *Args) ValuesInt16(flag, desc string) []int16
func (args *Args) ValuesInt8(flag, desc string) []int8
func (args *Args) ValuesInt(flag, desc string) []int
func (args *Args) ValuesUint64(flag, desc string) []uint64
func (args *Args) ValuesUint32(flag, desc string) []uint32
func (args *Args) ValuesUint16(flag, desc string) []uint16
func (args *Args) ValuesUint8(flag, desc string) []uint8
func (args *Args) ValuesUint(flag, desc string) []uint
func (args *Args) ValuesOption(flag string, validOptions []string, desc string) []string
```
<!--- gotomd::End::dcls::./../../Args.ValuesFloat64 Args.ValuesFloat32 Args.ValuesInt64 Args.ValuesInt32 Args.ValuesInt16 Args.ValuesInt8 Args.ValuesInt Args.ValuesUint64 Args.ValuesUint32 Args.ValuesUint16 Args.ValuesUint8 Args.ValuesUint Args.ValuesOption -->



---

## Contents

- [Source (main.go)](#source)
    - [Example: PASS: No Arguments](#pass-no-arguments)
    - [Example: PASS: Single Long Form](#pass-single-long-form)
    - [Example: PASS: Single Short Form](#pass-single-short-form)
    - [Example: PASS: Multiple Mixed Forms](#pass-multiple-mixed-forms)
    - [Example: FAIL: Extra Unknown Argument](#fail-extra-unknown-argument)

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

    if args.HasError() {
        fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage())
    } else {
        fmt.Printf("%d Name(s) Found: %v.\n", len(nameList), nameList)
        fmt.Printf("%d Byte(s) Found: %v.\n", len(byteList), byteList)
    }
}
```
<!--- gotomd::End::file::./main.go -->

[Top of Page](#example-flag-values) --
[Szargs Contents](../../README.md#contents)

### PASS: No Arguments

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

[Top of Page](#example-flag-values) --
[Szargs Contents](../../README.md#contents)

### PASS: Single Long Form

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

[Top of Page](#example-flag-values) --
[Szargs Contents](../../README.md#contents)

### PASS: Single Short Form

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

[Top of Page](#example-flag-values) --
[Szargs Contents](../../README.md#contents)


### PASS: Multiple Mixed Forms

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

[Top of Page](#example-flag-values) --
[Szargs Contents](../../README.md#contents)


### FAIL: Extra Unknown Argument

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

[Top of Page](#example-flag-values) --
[Szargs Contents](../../README.md#contents)
