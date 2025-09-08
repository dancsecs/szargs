/*
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
*/

package szargs

import "fmt"

// Next is a helper method to consume the next arg returning an error
// if there are no more args.  Otherwise it extracts the first argument
// returning the modified list.
func next(name string, args []string) (string, []string, error) {
	if len(args) < 1 {
		return "", args, fmt.Errorf(
			"%w: %s", ErrMissing, name,
		)
	}

	return args[0], args[1:], nil
}

// NextString removes and returns the next argument from the argument list.
//
// If no arguments remain, an error is registered.
//
// Returns the next argument value as a string.
func (args *Args) NextString(name, desc string) string {
	args.RegisterUsage(name, desc)
	result, newArgs, err := next(name, args.args)
	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextFloat64 removes and returns the next argument from the argument list,
// parsing it as a 64 bit floating point number.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for a float64, an error is registered.
//
// Returns the next argument value parsed as a float64.
func (args *Args) NextFloat64(name, desc string) float64 {
	var (
		arg    string
		result float64
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseFloat64(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextFloat32 removes and returns the next argument from the argument list,
// parsing it as a 32 bit floating point number.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for a float32, an error is registered.
//
// Returns the next argument value parsed as a float32.
func (args *Args) NextFloat32(name, desc string) float32 {
	var (
		arg    string
		result float32
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseFloat32(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextInt64 removes and returns the next argument from the argument list,
// parsing it as a signed 64 bit integer.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for an int64, an error is registered.
//
// Returns the next argument value parsed as an int64.
func (args *Args) NextInt64(name, desc string) int64 {
	var (
		arg    string
		result int64
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseInt64(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextInt32 removes and returns the next argument from the argument list,
// parsing it as a signed 32 bit integer.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for an int32, an error is registered.
//
// Returns the next argument value parsed as an int32.
func (args *Args) NextInt32(name, desc string) int32 {
	var (
		arg    string
		result int32
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseInt32(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextInt16 removes and returns the next argument from the argument list,
// parsing it as a signed 16 bit integer.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for an int16, an error is registered.
//
// Returns the next argument value parsed as an int16.
func (args *Args) NextInt16(name, desc string) int16 {
	var (
		arg    string
		result int16
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseInt16(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextInt8 removes and returns the next argument from the argument list,
// parsing it as a signed 8 bit integer.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for an int8, an error is registered.
//
// Returns the next argument value parsed as an int8.
func (args *Args) NextInt8(name, desc string) int8 {
	var (
		arg    string
		result int8
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseInt8(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextInt removes and returns the next argument from the argument list,
// parsing it as n signed integer.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for an int, an error is registered.
//
// Returns the next argument value parsed as an int.
func (args *Args) NextInt(name, desc string) int {
	var (
		arg    string
		result int
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseInt(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextUint64 removes and returns the next argument from the argument list,
// parsing it as an unsigned 64 bit integer.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for a uint64, an error is registered.
//
// Returns the next argument value parsed as a uint64.
func (args *Args) NextUint64(name, desc string) uint64 {
	var (
		arg    string
		result uint64
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseUint64(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextUint32 removes and returns the next argument from the argument list,
// parsing it as an unsigned 32 bit integer.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for a uint32, an error is registered.
//
// Returns the next argument value parsed as a uint32.
func (args *Args) NextUint32(name, desc string) uint32 {
	var (
		arg    string
		result uint32
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseUint32(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextUint16 removes and returns the next argument from the argument list,
// parsing it as an unsigned 16 bit integer.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for a uint16, an error is registered.
//
// Returns the next argument value parsed as a uint16.
func (args *Args) NextUint16(name, desc string) uint16 {
	var (
		arg    string
		result uint16
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseUint16(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextUint8 removes and returns the next argument from the argument list,
// parsing it as an unsigned 8 bit integer.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for a uint8, an error is registered.
//
// Returns the next argument value parsed as a uint8.
func (args *Args) NextUint8(name, desc string) uint8 {
	var (
		arg    string
		result uint8
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseUint8(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextUint removes and returns the next argument from the argument list,
// parsing it as an unsigned integer.
//
// If no arguments remain, or if the value has invalid syntax or is out of
// range for a uint, an error is registered.
//
// Returns the next argument value parsed as a uint.
func (args *Args) NextUint(name, desc string) uint {
	var (
		arg    string
		result uint
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseUint(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextOption removes and returns the next argument from the argument list.
// The value must match one of the entries in validOptions.
//
// If no arguments remain, or if the value is not found in validOptions,
// an error is registered.
//
// Returns the next argument value.
func (args *Args) NextOption(
	name string, validOptions []string, desc string,
) string {
	var (
		arg    string
		result string
		err    error
	)

	args.RegisterUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseOption(name, arg, validOptions)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}
