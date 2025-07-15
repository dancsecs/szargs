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

// ValueString scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueString(flag, desc string) (string, bool) {
	args.addUsage(flag, desc)
	result, found, newArgs, err := argFlag(flag).value(args.args)
	args.args = newArgs
	args.PushErr(err)

	return result, found
}

// ValueFloat64 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueFloat64(flag, desc string) (float64, bool) {
	var (
		arg    string
		found  bool
		result float64
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueFloat32 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueFloat32(flag, desc string) (float32, bool) {
	var (
		arg    string
		found  bool
		result float32
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueInt64 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueInt64(flag, desc string) (int64, bool) {
	var (
		arg    string
		found  bool
		result int64
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueInt32 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueInt32(flag, desc string) (int32, bool) {
	var (
		arg    string
		found  bool
		result int32
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueInt16 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueInt16(flag, desc string) (int16, bool) {
	var (
		arg    string
		found  bool
		result int16
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueInt8 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueInt8(flag, desc string) (int8, bool) {
	var (
		arg    string
		found  bool
		result int8
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueInt scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueInt(flag, desc string) (int, bool) {
	var (
		arg    string
		found  bool
		result int
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueUint64 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueUint64(flag, desc string) (uint64, bool) {
	var (
		arg    string
		found  bool
		result uint64
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueUint32 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueUint32(flag, desc string) (uint32, bool) {
	var (
		arg    string
		found  bool
		result uint32
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueUint16 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueUint16(flag, desc string) (uint16, bool) {
	var (
		arg    string
		found  bool
		result uint16
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueUint8 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueUint8(flag, desc string) (uint8, bool) {
	var (
		arg    string
		found  bool
		result uint8
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueUint scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueUint(flag, desc string) (uint, bool) {
	var (
		arg    string
		found  bool
		result uint
		err    error
	)

	args.addUsage(flag, desc)

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

// ValueOption scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (args *Args) ValueOption(
	flag string, validOptions []string, desc string,
) (string, bool) {
	var (
		arg    string
		found  bool
		result string
		err    error
	)

	args.addUsage(flag, desc)

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
