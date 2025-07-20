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

# Package szargs


## Overview

<!--- gotomd::Bgn::doc::./package -->
```go
package szargs
```

Package szargs provides a minimal and consistent interface for retrieving
settings from command-line arguments ([]string) and environment variables.

It supports three types of arguments:
  - Flagged: Identified by a single dash (e.g., "-v") for short flags, or a
    double dash (e.g., "--dir") for long-form flags. Flags may be standalone
    booleans or followed by a value.
  - Positional: Identified by their order in the argument list after all
    flagged arguments have been processed. Positional arguments can be
    retrieved in two forms: Next (in order) or Last (ensuring no trailing
    arguments remain).
  - Settings: A composite configuration mechanism that combines a default
    value, an environment variable, and a flagged argument—allowing each to
    override the previous in precedence: default < env < flag.

The package includes built-in parsers for standard Go data types.

Usage centers around the Args type, created using:

    szargs.New(programDesc string, programArgs []string)

The `programArgs` slice must include the program name as the first element;
this is ignored during argument parsing.

After retrieving all relevant arguments, the `Args.Done()` method must be
called to report an error if any unprocessed arguments remain.

This utility reflects a preference for simplicity and clarity in tooling. If
it helps your project flow a little more smoothly, it's done its job.

NOTE: Documentation reviewed and polished with the assistance of ChatGPT from
OpenAI.
<!--- gotomd::End::doc::./package -->


---

## Contents

- [Usage](#usage)
  - [Example: Average](example/average/README.md#example-average)
    - Description: Calculates the average of input values.

- [Boolean Flags](#boolean-flags)
  - [Example: Boolean Is](example/booleanIs/README.md#example-boolean-is)
    - Description: Demonstrates use of single boolean flag.
  - [Example: Boolean Count](example/booleanCount/README.md#example-boolean-count)
    - Description: Demonstrates use of counting multiple boolean flags.

- [Value Flagged Arguments](#value-flagged-arguments)
  - [Example: Flagged Value](example/value/README.md#example-flagged-value)
    - Description: Demonstrates use of typed flagged arguments.

- [Value Flagged Slices](#value-flagged-slices)
  - [Example: Flagged Values](example/values/README.md#example-flagged-values)
    - Description: Demonstrates use of multiple typed flagged arguments.

- [Positional Arguments](#positional-arguments)
  - [Example: Next](example/next/README.md#example-positional-arguments)
    - Description: Demonstrates use of positional arguments.

- [Settings](#settings)
  - [Example: Settings](example/setting/README.md#example-settings)
    - Description: Demonstrates use of settings with environment overrides.

- [Version](#version)

## Usage

Generally the flow of argument extraction proceeds as follows:

```go
func main() {
  // Create the args object with a program description (that will be used in
  // the Usage message) and the system arguments.  These will be copied
  // leaving the original os.Args untouched.
  args := szargs.New(
        "A simple demo of values flag.",
        os.Args,
    )

  // Flagged arguments are then extracted using various methods defined on the
  // args object.

  verbose := args.Count("[-v | --verbose ...]","The verbose level.")
  lines := args.Value("[-n | --num numOfLines]","Number of lines to display.")

  // Positional arguments are then extracted.
  file := args.Next("filename","The file to display the lines.")

  // All expected arguments have now been extracted.
  args.Done()

  if args.HasErr() {
    //Report any errors and optionally providing a Usage message.
    fmt.Fprintf(os.Stderr, "Error: %v\n\n%s\n", args.Err(), args.Usage())
  } else {
    // Process with the arguments.
  }
}
```

General functions operating on the state of the szargs object can be divided
into three categories as follows:

Relating to errors:

<!--- gotomd::Bgn::dcln::./Args.HasErr Args.PushErr Args.Err -->
```go
// HasErr returns true if any errors have been encountered or registered.
func (args *Args) HasErr() bool

// PushErr registers the provided error if not nil to the Args error stack.
func (args *Args) PushErr(err error)

// Err returns any errors encountered or registered while parsing the
// arguments.
func (args *Args) Err() error
```
<!--- gotomd::End::dcln::./Args.HasErr Args.PushErr Args.Err -->

Relating to the raw argument list:

<!--- gotomd::Bgn::dcln::./Args.HasNext Args.PushArg Args.Args -->
```go
// HasNext returns true if any arguments remain unabsorbed.
func (args *Args) HasNext() bool

// PushArg places the supplied argument to the end of the internal args list.
func (args *Args) PushArg(arg string)

// Args returns a copy of the current argument list.
func (args *Args) Args() []string
```
<!--- gotomd::End::dcln::./Args.HasNext Args.PushArg Args.Args -->

And general reporting and processing:

<!--- gotomd::Bgn::dcln::./Args.Usage Args.Done -->
```go
// Usage returns a usage message based on the parsed arguments.
func (args *Args) Usage() string

// Done registers an error if there are any remaining arguments.
func (args *Args) Done()
```
<!--- gotomd::End::dcln::./Args.Usage Args.Done -->

A working example can be found in the example directory as described here:
- [Example: Average](example/average/README.md#example-average)

[Contents](#contents)

## Boolean Flags

Boolean flags are defined by their presence only.  If they are present then
they are true and/or counted.  If not present then they are considered false.
There are two methods that operate with boolean flags as follows:

<!--- gotomd::Bgn::dcln::./Args.Is Args.Count -->
```go
// Is returns true if the flag is present one and only one time.
func (args *Args) Is(flag, desc string) bool

// Count returns the number of times the flag appears.
func (args *Args) Count(flag, desc string) int
```
<!--- gotomd::End::dcln::./Args.Is Args.Count -->

- [Example: Boolean Is](example/booleanIs/README.md#example-boolean-is)
- [Example: Boolean Count](example/booleanCount/README.md#example-boolean-count)

[Contents](#contents)

## Value Flagged Arguments

A flagged argument has two components: the flag followed by the value. It may
only appear once in the argument list.  The basic string functions are:

<!--- gotomd::Bgn::dcln::./Args.ValueString Args.ValueOption -->
```go
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

// ValueOption scans for a specific flagged argument (e.g., "--mode value")
// and captures its associated value. The flag and its value are removed from
// the argument list.
// 
// If the flag appears more than once, or if it lacks a following value, an
// error is registered. If the value is not found in the provided list of
// validOptions, an error is also registered.
// 
// Returns the value and a boolean indicating whether the flag was found.
func (args *Args) ValueOption(flag string, validOptions []string, desc string) (string, bool)
```
<!--- gotomd::End::dcln::./Args.ValueString Args.ValueOption -->

with numeric versions for basic go data types

<!--- gotomd::Bgn::dcls::./Args.ValueFloat64 Args.ValueFloat32 Args.ValueInt64 Args.ValueInt32 Args.ValueInt16 Args.ValueInt8 Args.ValueInt Args.ValueUint64 Args.ValueUint32 Args.ValueUint16 Args.ValueUint8 Args.ValueUint -->
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
```
<!--- gotomd::End::dcls::./Args.ValueFloat64 Args.ValueFloat32 Args.ValueInt64 Args.ValueInt32 Args.ValueInt16 Args.ValueInt8 Args.ValueInt Args.ValueUint64 Args.ValueUint32 Args.ValueUint16 Args.ValueUint8 Args.ValueUint -->

[Contents](#contents)


## Value Flagged Slices

A flagged argument has two components: the flag followed by the value.
Multiple instances may be provided with all the values collected and returned
in a slice. The basic string functions are:

<!--- gotomd::Bgn::dcln::./Args.ValuesString Args.ValuesOption -->
```go
// ValuesString scans for repeated instances of the specified flag and
// captures the following values as a slice of strings. The flags and values
// are removed from the argument list.
// 
// If any instance of the flag lacks a following value, an error is
// registered.
// 
// Returns a slice of the captured string values.
func (args *Args) ValuesString(flag, desc string) []string

// ValuesOption scans for repeated instances of the specified flag and
// captures the following values. Each value must appear in the provided list
// of validOptions. The flags and values are removed from the argument list.
// 
// If any flag lacks a following value, or if a value is not found in
// validOptions, an error is registered.
// 
// Returns a slice of the captured values.
func (args *Args) ValuesOption(flag string, validOptions []string, desc string) []string
```
<!--- gotomd::End::dcln::./Args.ValuesString Args.ValuesOption -->

with numeric versions for basic go data types

<!--- gotomd::Bgn::dcls::./Args.ValuesFloat64 Args.ValuesFloat32 Args.ValuesInt64 Args.ValuesInt32 Args.ValuesInt16 Args.ValuesInt8 Args.ValuesInt Args.ValuesUint64 Args.ValuesUint32 Args.ValuesUint16 Args.ValuesUint8 Args.ValuesUint -->
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
```
<!--- gotomd::End::dcls::./Args.ValuesFloat64 Args.ValuesFloat32 Args.ValuesInt64 Args.ValuesInt32 Args.ValuesInt16 Args.ValuesInt8 Args.ValuesInt Args.ValuesUint64 Args.ValuesUint32 Args.ValuesUint16 Args.ValuesUint8 Args.ValuesUint -->

[Contents](#contents)

## Positional Arguments

A positional argument depends on its location in the argument list.  Since
flagged arguments are not automatically distinguished from positional ones, it
is recommended to extract all flagged arguments first—before retrieving
positional ones.  The basic string functions are:

<!--- gotomd::Bgn::dcln::./Args.NextString Args.NextOption -->
```go
// NextString removes and returns the next argument from the argument list.
// 
// If no arguments remain, an error is registered.
// 
// Returns the next argument value as a string.
func (args *Args) NextString(name, desc string) string

// NextOption removes and returns the next argument from the argument list.
// The value must match one of the entries in validOptions.
// 
// If no arguments remain, or if the value is not found in validOptions,
// an error is registered.
// 
// Returns the next argument value.
func (args *Args) NextOption(name string, validOptions []string, desc string) string
```
<!--- gotomd::End::dcln::./Args.NextString Args.NextOption -->

with numeric versions for basic go data types

<!--- gotomd::Bgn::dcls::./Args.NextFloat64 Args.NextFloat32 Args.NextInt64 Args.NextInt32 Args.NextInt16 Args.NextInt8 Args.NextInt Args.NextUint64 Args.NextUint32 Args.NextUint16 Args.NextUint8 Args.NextUint -->
```go
func (args *Args) NextFloat64(name, desc string) float64
func (args *Args) NextFloat32(name, desc string) float32
func (args *Args) NextInt64(name, desc string) int64
func (args *Args) NextInt32(name, desc string) int32
func (args *Args) NextInt16(name, desc string) int16
func (args *Args) NextInt8(name, desc string) int8
func (args *Args) NextInt(name, desc string) int
func (args *Args) NextUint64(name, desc string) uint64
func (args *Args) NextUint32(name, desc string) uint32
func (args *Args) NextUint16(name, desc string) uint16
func (args *Args) NextUint8(name, desc string) uint8
func (args *Args) NextUint(name, desc string) uint
```
<!--- gotomd::End::dcls::./Args.NextFloat64 Args.NextFloat32 Args.NextInt64 Args.NextInt32 Args.NextInt16 Args.NextInt8 Args.NextInt Args.NextUint64 Args.NextUint32 Args.NextUint16 Args.NextUint8 Args.NextUint -->


[Contents](#contents)

## Argument Settings

A setting implements an argument that has a default that can be overridden by
an system environment variable which can be overridden by a flagged value. The
basic string functions are:

<!--- gotomd::Bgn::dcln::./Args.SettingString Args.SettingOption  Args.SettingIs -->
```go
// SettingString returns a configuration value based on a default,
// optionally overridden by an environment variable, and further overridden
// by a flagged command-line argument.
// 
// Returns the final selected string value.
func (args *Args) SettingString(flag, env, def, desc string) string

// SettingOption returns a configuration value based on a default,
// optionally overridden by an environment variable, and further overridden
// by a flagged command-line argument.
// 
// If the final value is not found in the list of validOptions,
// an error is registered.
// 
// Returns the final selected value.
func (args *Args) SettingOption(flag, env string, def string, validOptions []string, desc string) string

// SettingIs returns true if a specified environment variable is set to a
// truthy value, or if a corresponding boolean command-line flag is present.
// 
// Unlike other Setting methods, there is no default.
// 
// The environment variable is considered true if it is set to one of: "",
// "T", "Y", "TRUE", "YES", "ON" or "1" (case-insensitive). Any other value is
// considered false.
// 
// The command-line flag override takes no value—its presence alone indicates
// true.
// 
// Returns the resulting boolean value.
func (args *Args) SettingIs(flag, env string, desc string) bool
```
<!--- gotomd::End::dcln::./Args.SettingString Args.SettingOption  Args.SettingIs -->

with numeric versions for basic go data types

<!--- gotomd::Bgn::dcls::./Args.SettingFloat64 Args.SettingFloat32 Args.SettingInt64 Args.SettingInt32 Args.SettingInt16 Args.SettingInt8 Args.SettingInt Args.SettingUint64 Args.SettingUint32 Args.SettingUint16 Args.SettingUint8 Args.SettingUint -->
```go
func (args *Args) SettingFloat64(flag, env string, def float64, desc string) float64
func (args *Args) SettingFloat32(flag, env string, def float32, desc string) float32
func (args *Args) SettingInt64(flag, env string, def int64, desc string) int64
func (args *Args) SettingInt32(flag, env string, def int32, desc string) int32
func (args *Args) SettingInt16(flag, env string, def int16, desc string) int16
func (args *Args) SettingInt8(flag, env string, def int8, desc string) int8
func (args *Args) SettingInt(flag, env string, def int, desc string) int
func (args *Args) SettingUint64(flag, env string, def uint64, desc string) uint64
func (args *Args) SettingUint32(flag, env string, def uint32, desc string) uint32
func (args *Args) SettingUint16(flag, env string, def uint16, desc string) uint16
func (args *Args) SettingUint8(flag, env string, def uint8, desc string) uint8
func (args *Args) SettingUint(flag, env string, def uint, desc string) uint
```
<!--- gotomd::End::dcls::./Args.SettingFloat64 Args.SettingFloat32 Args.SettingInt64 Args.SettingInt32 Args.SettingInt16 Args.SettingInt8 Args.SettingInt Args.SettingUint64 Args.SettingUint32 Args.SettingUint16 Args.SettingUint8 Args.SettingUint -->


[Contents](#contents)

## Version

- Current: v0.0.6
- Go Version: 1.23+

[Contents](#contents)
