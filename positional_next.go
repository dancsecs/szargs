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

// NextString extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextString(name, desc string) string {
	args.addUsage(name, desc)
	result, newArgs, err := next(name, args.args)
	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextFloat64 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextFloat64(name, desc string) float64 {
	var (
		arg    string
		result float64
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseFloat64(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextFloat32 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextFloat32(name, desc string) float32 {
	var (
		arg    string
		result float32
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseFloat32(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextInt64 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextInt64(name, desc string) int64 {
	var (
		arg    string
		result int64
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseInt64(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextInt32 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextInt32(name, desc string) int32 {
	var (
		arg    string
		result int32
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseInt32(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextInt16 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextInt16(name, desc string) int16 {
	var (
		arg    string
		result int16
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseInt16(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextInt8 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextInt8(name, desc string) int8 {
	var (
		arg    string
		result int8
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseInt8(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextInt extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextInt(name, desc string) int {
	var (
		arg    string
		result int
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseInt(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextUint64 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextUint64(name, desc string) uint64 {
	var (
		arg    string
		result uint64
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseUint64(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextUint32 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextUint32(name, desc string) uint32 {
	var (
		arg    string
		result uint32
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseUint32(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextUint16 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextUint16(name, desc string) uint16 {
	var (
		arg    string
		result uint16
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseUint16(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextUint8 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextUint8(name, desc string) uint8 {
	var (
		arg    string
		result uint8
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseUint8(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}

// NextUint extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func (args *Args) NextUint(name, desc string) uint {
	var (
		arg    string
		result uint
		err    error
	)

	args.addUsage(name, desc)

	arg, newArgs, err := next(name, args.args)

	if err == nil {
		result, err = parseUint(name, arg)
	}

	args.args = newArgs
	args.PushErr(err)

	return result
}
