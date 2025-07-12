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

// Last is a helper method to consume the next arg returning an error
// if there are no more args or any extra args.  Otherwise it extracts the
// argument.
func (args *Args) last(name, desc string) string {
	args.addUsage(name, desc)

	value, cleanedArgs, err := next(name, args.Args())

	args.args = cleanedArgs
	args.PushErr(err)

	args.Done()

	if args.Err() == nil {
		return value
	}

	return ""
}

// LastString extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastString(name, desc string) string {
	args.addUsage(name, desc)

	return args.last(name, desc)
}

// LastFloat64 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastFloat64(name, desc string) float64 {
	var (
		arg    string
		result float64
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseFloat64(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastFloat32 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastFloat32(name, desc string) float32 {
	var (
		arg    string
		result float32
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseFloat32(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastInt64 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastInt64(name, desc string) int64 {
	var (
		arg    string
		result int64
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseInt64(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastInt32 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastInt32(name, desc string) int32 {
	var (
		arg    string
		result int32
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseInt32(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastInt16 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastInt16(name, desc string) int16 {
	var (
		arg    string
		result int16
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseInt16(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastInt8 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastInt8(name, desc string) int8 {
	var (
		arg    string
		result int8
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseInt8(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastInt extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastInt(name, desc string) int {
	var (
		arg    string
		result int
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseInt(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastUint64 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastUint64(name, desc string) uint64 {
	var (
		arg    string
		result uint64
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseUint64(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastUint32 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastUint32(name, desc string) uint32 {
	var (
		arg    string
		result uint32
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseUint32(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastUint16 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastUint16(name, desc string) uint16 {
	var (
		arg    string
		result uint16
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseUint16(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastUint8 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastUint8(name, desc string) uint8 {
	var (
		arg    string
		result uint8
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseUint8(name, arg)
		args.PushErr(err)
	}

	return result
}

// LastUint extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func (args *Args) LastUint(name, desc string) uint {
	var (
		arg    string
		result uint
		err    error
	)

	arg = args.last(name, desc)

	if args.Err() == nil {
		result, err = parseUint(name, arg)
		args.PushErr(err)
	}

	return result
}
