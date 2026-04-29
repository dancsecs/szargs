<!---             *****  AUTO GENERATED:  DO NOT MODIFY  ***** -->
<!---           MODIFY TEMPLATE: 'example/average/.README.gtm.md' -->
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

# Example Average

## Overview

This larger example demos several aspects of using the szargs library.  It
performs one of two operations (specified in an optional positional argument)
on a list of float64 values passed as flagged arguments.  Further it uses a
boolean count to set a verbose level to increase the chattiness of the
program.

It demonstrates the following szargs functions:

```go
// New creates a new Args object based in the arguments passed.  The first
// element of the arguments must be the program name.
func New(programDesc string, args []string) *Args

// Count returns the number of times the flag appears.
func (args *Args) Count(flag, desc string) int

// ValuesFloat64 scans for repeated instances of the specified flag and parses
// the following values as 64 bit floating point numbers. The flags and values
// are removed from the argument list.
// 
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for a float64, an error is registered.
// 
// Returns a slice of the parsed float64 values.
func (args *Args) ValuesFloat64(flag, desc string) []float64

// HasNext returns true if any arguments remain unabsorbed.
func (args *Args) HasNext() bool

// PushArg places the supplied argument to the end of the internal args list.
func (args *Args) PushArg(arg string)

// NextOption removes and returns the next argument from the argument list.
// The value must match one of the entries in validOptions.
// 
// If no arguments remain, or if the value is not found in validOptions,
// an error is registered.
// 
// Returns the next argument value.
func (args *Args) NextOption(name string, validOptions []string, desc string) string

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
    - [Example: **PASS**: No Arguments](#pass-no-arguments)
    - [Example: **PASS**: Just Operation Add](#pass-just-operation-add)
    - [Example: **PASS**: Just operation Average](#pass-just-operation-average)
    - [Example: **PASS**: One Number Defaulting To Add](#pass-one-number-defaulting-to-add)
    - [Example: **PASS**: Two Numbers Defaulting To Add](#pass-two-numbers-defaulting-to-add)
    - [Example: **PASS**: One Number Average](#pass-one-number-average)
    - [Example: **PASS**: Two Number Average](#pass-two-number-average)
    - [Example: **PASS**: Three Number Add](#pass-three-number-average)
    - [Example: **PASS**: Three Number Average](#pass-three-number-average)
    - [Example: **FAIL**: Extra Unknown Argument](#fail-extra-unknown-argument)
    - [Example: **FAIL**: Invalid Operation](#fail-invalid-operation)
    - [Example: **FAIL**: Invalid Number](#fail-invalid-number)

## Source

The sources used for this example are broken into three files as follows:

This defines the verbose processing.

```bash
cat ./verbose.go
```

```go
package main

import (
    "fmt"
)

// Supported verbose levels.
const (
    vAll = iota
    vLv1
    vLv2
    vLv3
)

var verbose int //nolint:goCheckNoGlobals // Ok.

func sayPrintf(minLevel int, msgFormat string, msgArgs ...any) {
    if verbose >= minLevel {
        fmt.Printf(msgFormat, msgArgs...) //nolint:forbidigo // Ok.
    }
}
```

and this defines all of the processing to be performed if and only if the
arguments are parsed without any errors.

```bash
cat ./process.go
```

```go
package main

func process(numbers []float64, operation string) {
    sayPrintf(vLv1, "Verbose set to %d\n", verbose)
    sayPrintf(vLv2, "Read in %d numbers\n", len(numbers))
    sayPrintf(vLv2, "Operation: %s\n", operation)

    sum := float64(0)

    for i, n := range numbers {
        sayPrintf(vLv3, "Number (%d): %f\n", i, n)
        sum += n
    }

    if operation == "average" {
        avg := float64(0)
        if len(numbers) > 0 {
            avg = sum / float64(len(numbers))
        }

        sayPrintf(vAll, "Avg: %f\n", avg)
    } else {
        sayPrintf(vAll, "Sum: %f\n", sum)
    }
}
```

finally the mainline where the arguments are parsed is defined as follows:

```bash
cat ./main.go
```

```go
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
    if args.HasErr() {
        fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage(0))
    } else {
        process(numbers, operation)
    }
}
```

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **PASS**: No Arguments

An empty list sums and averages to zero for simplicity. No error is reported.

---
```bash
go run .
```

```
Sum: 0.000000
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Just Operation Add

---
```bash
go run . add
```

```
Sum: 0.000000
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Just Operation Average

---
```bash
go run . average
```

```
Avg: 0.000000
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **PASS**: One Number Defaulting To Add

---
```bash
go run . -n 1000
```

```
Sum: 1000.000000
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Two Numbers Defaulting To Add

---
```bash
go run . --number 1 -n 2
```

```
Sum: 3.000000
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **PASS**: One Number Average

---
```bash
go run . --number 512 average
```

```
Avg: 512.000000
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Two Number Average

---
```bash
go run . -n 100 -n 200 average
```

```
Avg: 150.000000
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Three Number Add

---
```bash
go run . --number 23 --number 56 -n 22 add
```

```
Sum: 101.000000
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **PASS**: Three Number Average

---
```bash
go run . -n 23 -n 56 -n 22 average
```

```
Avg: 33.666667
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **FAIL**: Extra Unknown Argument

---
```bash
go run . add extraUnknownArgument
```

```
Error: unexpected argument: [extraUnknownArgument]

usage: average [-v | --verbose ...] [-n | --number float64 ...] [operation]

A simple utility to add or average a number list.

    [-v | --verbose ...]
        The verbose level.

    [-n | --number float64 ...]
        The numbers to act on.

    [operation]
        The operation (add or average) defaulting to add.
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **FAIL**: Invalid operation

---
```bash
go run . invalidOperation
```

```
Error: invalid option: 'invalidOperation' ([operation] must be one of [add average])

usage: average [-v | --verbose ...] [-n | --number float64 ...] [operation]

A simple utility to add or average a number list.

    [-v | --verbose ...]
        The verbose level.

    [-n | --number float64 ...]
        The numbers to act on.

    [operation]
        The operation (add or average) defaulting to add.
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)

### **FAIL**: Invalid Number

---
```bash
go run . -n invalidNumber
```

```
Error: invalid float64: syntax: [-n | --number float64 ...]: 'invalidNumber'

usage: average [-v | --verbose ...] [-n | --number float64 ...] [operation]

A simple utility to add or average a number list.

    [-v | --verbose ...]
        The verbose level.

    [-n | --number float64 ...]
        The numbers to act on.

    [operation]
        The operation (add or average) defaulting to add.
```
---

[Top of Page](#example-average) --
[Szargs Contents](../../README.md#contents)
