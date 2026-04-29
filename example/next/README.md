<!---             *****  AUTO GENERATED:  DO NOT MODIFY  ***** -->
<!---            MODIFY TEMPLATE: 'example/next/.README.gtm.md' -->
<!---               See: 'https://github.com/dancsecs/gotomd' -->

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

# Example Next Positional Argument

## Overview

This example demos the parsing of the next argument.  If found the argument is
removed and parsed as encountered.  A missing argument will generate an error.

It demonstrates the following szargs functions:

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

// Usage returns a usage messages representing the Args object.  It is
// formatted to the lineWidth provided.  A zero uses the defaultLineWidth
// while a negative value caused an effort to determine if writing to a
// terminal and if so using its width otherwise defaulting.
func (args *Args) Usage(lineWidth int) string
```

## Contents

- [Source (main.go)](#source)
    - [Example: **PASS**: Count Defaulted](#pass-count-defaulted)
    - [Example: **PASS**: Count Provided](#pass-count-provided)
    - [Example: **FAIL**: No Arguments](#fail-no-arguments)
    - [Example: **FAIL**: Extra Unknown Argument](#fail-extra-unknown-argument)

## Source

The source used for this example.  It simply defines a new szargs.Arg object,
then defines two positional arguments (one mandatory and one optional).
Finally the Done function is called to insure that no more arguments exist in
the list.

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
        "A simple demo of repeating a string.",
        os.Args,
    )

    mandatory := args.NextString(
        "message",
        "What to repeat.",
    )

    if !args.HasNext() {
        args.PushArg(("3")) // Default value dor optional integer.
    }

    optional := args.NextUint(
        "[times]",
        "The number of times to repeat.  Defaults to 3.",
    )

    args.Done() // All arguments should have consumed.

    if args.HasErr() {
        fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage(0))
    } else {
        for range optional {
            fmt.Println(mandatory)
        }
    }
}
```

[Top of Page](#example-next-positional-argument) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Count Defaulted

---
```bash
go run . StringToRepeat
```

```
StringToRepeat
StringToRepeat
StringToRepeat
```
---

[Top of Page](#example-next-positional-argument) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Count Provided

---
```bash
go run . StringToRepeat 5
```

```
StringToRepeat
StringToRepeat
StringToRepeat
StringToRepeat
StringToRepeat
```
---

[Top of Page](#example-next-positional-argument) --
[Szargs Contents](../../README.md#contents)

### **FAIL**: No Arguments

---
```bash
go run .
```

```
Error: missing argument: message

usage: next message [times]

A simple demo of repeating a string.

    message
        What to repeat.

    [times]
        The number of times to repeat.  Defaults to 3.
```
---

[Top of Page](#example-next-positional-argument) --
[Szargs Contents](../../README.md#contents)

### **FAIL**: Extra Unknown Argument

The error is generated by the ```args.Done()``` statement causing both the
error and a usage statement to be returned.

---
```bash
go run . stringToRepeat 5 extraUnknownArgument
```

```
Error: unexpected argument: [extraUnknownArgument]

usage: next message [times]

A simple demo of repeating a string.

    message
        What to repeat.

    [times]
        The number of times to repeat.  Defaults to 3.
```
---

[Top of Page](#example-next-positional-argument) --
[Szargs Contents](../../README.md#contents)
