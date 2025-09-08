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

// ValueString scans for a specific flagged argument and captures its
// following value as a string. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once or lacks a following value, an error is
// registered.
//
// Returns the string value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueString(flag, desc string) (string, bool) {
	args.RegisterUsage(flag, desc)
	result, found, newArgs, err := argFlag(flag).value(args.args)
	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueFloat64 scans for a specific flagged argument and parses its value as
// a 64 bit floating point number. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for a float64, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueFloat64(flag, desc string) (float64, bool) {
	var (
		arg    string
		found  bool
		result float64
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseFloat64(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueFloat32 scans for a specific flagged argument and parses its value as
// a 32 bit floating point number. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for a float32, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueFloat32(flag, desc string) (float32, bool) {
	var (
		arg    string
		found  bool
		result float32
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseFloat32(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueInt64 scans for a specific flagged argument and parses its value as a
// signed 64 bit integer. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for an int64, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueInt64(flag, desc string) (int64, bool) {
	var (
		arg    string
		found  bool
		result int64
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseInt64(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueInt32 scans for a specific flagged argument and parses its value as a
// signed 32 bit integer. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for an int32, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueInt32(flag, desc string) (int32, bool) {
	var (
		arg    string
		found  bool
		result int32
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseInt32(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueInt16 scans for a specific flagged argument and parses its value as a
// signed 16 bit integer. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for an int16, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueInt16(flag, desc string) (int16, bool) {
	var (
		arg    string
		found  bool
		result int16
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseInt16(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueInt8 scans for a specific flagged argument and parses its value as a
// signed 8 bit integer. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for an int8, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueInt8(flag, desc string) (int8, bool) {
	var (
		arg    string
		found  bool
		result int8
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseInt8(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueInt scans for a specific flagged argument and parses its value as a
// signed integer. The flag and its value are removed from the argument
// list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for an int, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueInt(flag, desc string) (int, bool) {
	var (
		arg    string
		found  bool
		result int
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseInt(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueUint64 scans for a specific flagged argument and parses its value as
// an unsigned 64 bit integer. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for a uint64, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueUint64(flag, desc string) (uint64, bool) {
	var (
		arg    string
		found  bool
		result uint64
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseUint64(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueUint32 scans for a specific flagged argument and parses its value as
// an unsigned 32 bit integer. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for a uint32, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueUint32(flag, desc string) (uint32, bool) {
	var (
		arg    string
		found  bool
		result uint32
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseUint32(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueUint16 scans for a specific flagged argument and parses its value as
// an unsigned 16 bit integer. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for a uint16, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueUint16(flag, desc string) (uint16, bool) {
	var (
		arg    string
		found  bool
		result uint16
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseUint16(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueUint8 scans for a specific flagged argument and parses its value as an
// unsigned 8 bit integer. The flag and its value are removed from the
// argument list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for a uint8, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueUint8(flag, desc string) (uint8, bool) {
	var (
		arg    string
		found  bool
		result uint8
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseUint8(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueUint scans for a specific flagged argument and parses its value as an
// unsigned integer. The flag and its value are removed from the argument
// list.
//
// If the flag appears more than once, lacks a following value, or if the
// value has invalid syntax or is out of range for a uint, an error is
// registered.
//
// Returns the parsed value and a boolean indicating whether the flag was
// found.
func (args *Args) ValueUint(flag, desc string) (uint, bool) {
	var (
		arg    string
		found  bool
		result uint
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseUint(flag, arg)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueOption scans for a specific flagged argument (e.g., "--mode value")
// and captures its associated value. The flag and its value are removed from
// the argument list.
//
// If the flag appears more than once, or if it lacks a following value, an
// error is registered. If the value is not found in the provided list of
// validOptions, an error is also registered.
//
// Returns the value and a boolean indicating whether the flag was found.
func (args *Args) ValueOption(
	flag string, validOptions []string, desc string,
) (string, bool) {
	var (
		arg    string
		found  bool
		result string
		err    error
	)

	args.RegisterUsage(flag, desc)

	arg, found, newArgs, err := argFlag(flag).value(args.args)

	if err == nil && found {
		result, err = parseOption(flag, arg, validOptions)
		if err != nil {
			found = false
		}
	}

	args.args = newArgs
	args.PushErr(err)

	return result, found
}
