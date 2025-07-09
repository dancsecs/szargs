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

import (
	"fmt"
)

// Next is a helper method to consume the next arg returning an error
// if there are no more args.  Otherwise it extracts the first argument
// returning the modified list.
func Next(name string, args []string) (string, []string, error) {
	if len(args) < 1 {
		return "", args, fmt.Errorf(
			"%w: %s", ErrMissing, name,
		)
	}

	return args[0], args[1:], nil
}

// NextString extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextString(name string, args []string) (string, []string, error) {
	return Next(name, args)
}

// NextFloat64 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextFloat64(name string, args []string) (float64, []string, error) {
	var (
		arg    string
		result float64
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseFloat64(name, arg)
	}

	return result, args, err
}

// NextFloat32 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextFloat32(name string, args []string) (float32, []string, error) {
	var (
		arg    string
		result float32
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseFloat32(name, arg)
	}

	return result, args, err
}

// NextInt64 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextInt64(name string, args []string) (int64, []string, error) {
	var (
		arg    string
		result int64
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseInt64(name, arg)
	}

	return result, args, err
}

// NextInt32 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextInt32(name string, args []string) (int32, []string, error) {
	var (
		arg    string
		result int32
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseInt32(name, arg)
	}

	return result, args, err
}

// NextInt16 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextInt16(name string, args []string) (int16, []string, error) {
	var (
		arg    string
		result int16
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseInt16(name, arg)
	}

	return result, args, err
}

// NextInt8 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextInt8(name string, args []string) (int8, []string, error) {
	var (
		arg    string
		result int8
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseInt8(name, arg)
	}

	return result, args, err
}

// NextInt extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextInt(name string, args []string) (int, []string, error) {
	var (
		arg    string
		result int
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseInt(name, arg)
	}

	return result, args, err
}

// NextUint64 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextUint64(name string, args []string) (uint64, []string, error) {
	var (
		arg    string
		result uint64
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseUint64(name, arg)
	}

	return result, args, err
}

// NextUint32 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextUint32(name string, args []string) (uint32, []string, error) {
	var (
		arg    string
		result uint32
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseUint32(name, arg)
	}

	return result, args, err
}

// NextUint16 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextUint16(name string, args []string) (uint16, []string, error) {
	var (
		arg    string
		result uint16
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseUint16(name, arg)
	}

	return result, args, err
}

// NextUint8 extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextUint8(name string, args []string) (uint8, []string, error) {
	var (
		arg    string
		result uint8
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseUint8(name, arg)
	}

	return result, args, err
}

// NextUint extracts the next positional argument returning an error if none
// are present, otherwise it parses and returns the named data type.
func NextUint(name string, args []string) (uint, []string, error) {
	var (
		arg    string
		result uint
		err    error
	)

	arg, args, err = Next(name, args)

	if err == nil {
		result, err = parseUint(name, arg)
	}

	return result, args, err
}
