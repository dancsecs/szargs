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

// Value scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.
func (a Arg) Value(
	args []string,
) (string, bool, []string, error) {
	found := false
	value := ""
	cleanedArgs := make([]string, 0, len(args))
	err := error(nil)

	pushErr := func(newErr error) {
		if err == nil {
			err = newErr
		} else {
			err = fmt.Errorf("%w: %w", err, newErr)
		}
	}

	for i, mi := 0, len(args); i < mi; i++ {
		if a.argIs(args[i]) { //nolint:nestif // Ok.
			if (i + 1) >= mi {
				pushErr(
					fmt.Errorf(
						"%w: '%s value'",
						ErrMissing,
						a,
					),
				)
			} else {
				i++
				if found {
					pushErr(
						fmt.Errorf(
							"%w: '%s %s' already set to: '%s'",
							ErrAmbiguous,
							a,
							args[i],
							value,
						),
					)
				} else {
					value = args[i]
					found = true
				}
			}
		} else {
			cleanedArgs = append(cleanedArgs, args[i])
		}
	}

	if err == nil {
		return value, found, cleanedArgs, nil
	}

	return "", false, cleanedArgs, err
}

// ValueString scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueString(
	args []string,
) (string, bool, []string, error) {
	return a.Value(args)
}

// ValueFloat64 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueFloat64(
	args []string,
) (float64, bool, []string, error) {
	var (
		arg    string
		found  bool
		result float64
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseFloat64(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueFloat32 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueFloat32(
	args []string,
) (float32, bool, []string, error) {
	var (
		arg    string
		found  bool
		result float32
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseFloat32(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueInt64 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueInt64(
	args []string,
) (int64, bool, []string, error) {
	var (
		arg    string
		found  bool
		result int64
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseInt64(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueInt32 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueInt32(
	args []string,
) (int32, bool, []string, error) {
	var (
		arg    string
		found  bool
		result int32
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseInt32(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueInt16 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueInt16(
	args []string,
) (int16, bool, []string, error) {
	var (
		arg    string
		found  bool
		result int16
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseInt16(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueInt8 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueInt8(
	args []string,
) (int8, bool, []string, error) {
	var (
		arg    string
		found  bool
		result int8
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseInt8(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueInt scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueInt(
	args []string,
) (int, bool, []string, error) {
	var (
		arg    string
		found  bool
		result int
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseInt(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueUint64 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueUint64(
	args []string,
) (uint64, bool, []string, error) {
	var (
		arg    string
		found  bool
		result uint64
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseUint64(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueUint32 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueUint32(
	args []string,
) (uint32, bool, []string, error) {
	var (
		arg    string
		found  bool
		result uint32
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseUint32(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueUint16 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueUint16(
	args []string,
) (uint16, bool, []string, error) {
	var (
		arg    string
		found  bool
		result uint16
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseUint16(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueUint8 scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueUint8(
	args []string,
) (uint8, bool, []string, error) {
	var (
		arg    string
		found  bool
		result uint8
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseUint8(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}

// ValueUint scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) ValueUint(
	args []string,
) (uint, bool, []string, error) {
	var (
		arg    string
		found  bool
		result uint
		err    error
	)

	arg, found, args, err = a.Value(args)

	if err == nil && found {
		result, err = parseUint(string(a), arg)
		if err != nil {
			found = false
		}
	}

	return result, found, args, err
}
