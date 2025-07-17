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

# Example Flag Value


## Overview

This example demos the setting of a flag's value.  If found the matching flag
and its argument are removed as encountered.  An ambiguous error will be
generated if the flag appears more than once.  If a specialized value is named
then it must pass rules regarding its syntax and range.

<!--- gotomd::Bgn::dcln::./../../Args.ValueString -->
```go
// ValueString scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueString(flag, desc string) (string, bool)
```
<!--- gotomd::End::dcln::./../../Args.ValueString -->

and its specific versions:

<!--- gotomd::Bgn::dcls::./../../Args.ValueFloat64 Args.ValueFloat32 Args.ValueInt64 Args.ValueInt32 Args.ValueInt16 Args.ValueInt8 Args.ValueInt Args.ValueUint64 Args.ValueUint32 Args.ValueUint16 Args.ValueUint8 Args.ValueUint Args.ValueOption -->
```go
func (args *Args) ValueFloat64(flag, desc string) (float64, bool)
func (args *Args) ValueFloat32(flag, desc string) (float32, bool)
func (args *Args) ValueInt64(flag, desc string) (int64, bool)
func (args *Args) ValueInt32(flag, desc string) (int32, bool)
func (args *Args) ValueInt16(flag, desc string) (int16, bool)
func (args *Args) ValueInt8(flag, desc string) (int8, bool)
func (args *Args) ValueInt(flag, desc string) (int, bool)
func (args *Args) ValueUint64(flag, desc string) (uint64, bool)
func (args *Args) ValueUint32(flag, desc string) (uint32, bool)
func (args *Args) ValueUint16(flag, desc string) (uint16, bool)
func (args *Args) ValueUint8(flag, desc string) (uint8, bool)
func (args *Args) ValueUint(flag, desc string) (uint, bool)
func (args *Args) ValueOption(flag string, validOptions []string, desc string) (string, bool)
```
<!--- gotomd::End::dcls::./../../Args.ValueFloat64 Args.ValueFloat32 Args.ValueInt64 Args.ValueInt32 Args.ValueInt16 Args.ValueInt8 Args.ValueInt Args.ValueUint64 Args.ValueUint32 Args.ValueUint16 Args.ValueUint8 Args.ValueUint Args.ValueOption -->



---

## Contents

- [Source (main.go)](#source)
    - [Example: Successful No Arguments](#successful-no-arguments)
    - [Example: Successful Single Long Form](#successful-single-long-form)
    - [Example: Successful Single Short Form](#successful-single-short-form)
    - [Example: Failure Ambiguous Argument](#failure-ambiguous-argument)
    - [Example: Failure Extra Unknown Argument](#failure-extra-unknown-argument)

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

    if args.HasError() { //nolint:nestif // Ok for the demo.
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

[Top of Page](#example-flag-value) --
[Szargs Contents](../../README.md#contents)

### Successful No Arguments

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

[Top of Page](#example-flag-value) --
[Szargs Contents](../../README.md#contents)

### Successful Single Long Form

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

[Top of Page](#example-flag-value) --
[Szargs Contents](../../README.md#contents)

### Successful Single Short Form

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

[Top of Page](#example-flag-value) --
[Szargs Contents](../../README.md#contents)


### Failure Ambiguous Argument

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

[Top of Page](#example-flag-value) --
[Szargs Contents](../../README.md#contents)

### Failure Extra Unknown Argument

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

[Top of Page](#example-flag-value) --
[Szargs Contents](../../README.md#contents)
