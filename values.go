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

// ValuesString scans for repeated instances of the specified flag and
// captures the following values as a slice of strings. The flags and values
// are removed from the argument list.
//
// If any instance of the flag lacks a following value, an error is
// registered.
//
// Returns a slice of the captured string values.
func (args *Args) ValuesString(flag, desc string) []string {
	args.RegisterUsage(flag, desc)
	result, newArgs, err := argFlag(flag).values(args.args)
	args.args = newArgs
	args.PushErr(err)

	return result
}

// ValuesFloat64 scans for repeated instances of the specified flag and parses
// the following values as 64 bit floating point numbers. The flags and values
// are removed from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for a float64, an error is registered.
//
// Returns a slice of the parsed float64 values.
func (args *Args) ValuesFloat64(flag, desc string) []float64 {
	var (
		matches     []string
		cleanedArgs []string
		result      []float64
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem float64
			argErr  error
		)

		result = make([]float64, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseFloat64(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesFloat32 scans for repeated instances of the specified flag and parses
// the following values as 32 bit floating point numbers. The flags and values
// are removed from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for a float32, an error is registered.
//
// Returns a slice of the parsed float32 values.
func (args *Args) ValuesFloat32(flag, desc string) []float32 {
	var (
		matches     []string
		cleanedArgs []string
		result      []float32
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem float32
			argErr  error
		)

		result = make([]float32, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseFloat32(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesInt64 scans for repeated instances of the specified flag and parses
// the following values as signed 64 bit integers. The flags and values are
// removed from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for an int64, an error is registered.
//
// Returns a slice of the parsed int64 values.
func (args *Args) ValuesInt64(flag, desc string) []int64 {
	var (
		matches     []string
		cleanedArgs []string
		result      []int64
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem int64
			argErr  error
		)

		result = make([]int64, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseInt64(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesInt32 scans for repeated instances of the specified flag and parses
// the following values as signed 32 bit integers. The flags and values are
// removed from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for an int32, an error is registered.
//
// Returns a slice of the parsed int32 values.
func (args *Args) ValuesInt32(flag, desc string) []int32 {
	var (
		matches     []string
		cleanedArgs []string
		result      []int32
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem int32
			argErr  error
		)

		result = make([]int32, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseInt32(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesInt16 scans for repeated instances of the specified flag and parses
// the following values as signed 16 bit integers. The flags and values are
// removed from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for an int16, an error is registered.
//
// Returns a slice of the parsed int16 values.
func (args *Args) ValuesInt16(flag, desc string) []int16 {
	var (
		matches     []string
		cleanedArgs []string
		result      []int16
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem int16
			argErr  error
		)

		result = make([]int16, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseInt16(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesInt8 scans for repeated instances of the specified flag and parses
// the following values as signed 8 bit integers. The flags and values are
// removed from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for an int8, an error is registered.
//
// Returns a slice of the parsed int8 values.
func (args *Args) ValuesInt8(flag, desc string) []int8 {
	var (
		matches     []string
		cleanedArgs []string
		result      []int8
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem int8
			argErr  error
		)

		result = make([]int8, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseInt8(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesInt scans for repeated instances of the specified flag and parses
// the following values as signed integers. The flags and values are removed
// from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for an int, an error is registered.
//
// Returns a slice of the parsed int values.
func (args *Args) ValuesInt(flag, desc string) []int {
	var (
		matches     []string
		cleanedArgs []string
		result      []int
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem int
			argErr  error
		)

		result = make([]int, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseInt(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesUint64 scans for repeated instances of the specified flag and parses
// the following values as unsigned 64 bit integers. The flags and values are
// removed from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for a uint64, an error is registered.
//
// Returns a slice of the parsed uint64 values.
func (args *Args) ValuesUint64(flag, desc string) []uint64 {
	var (
		matches     []string
		cleanedArgs []string
		result      []uint64
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem uint64
			argErr  error
		)

		result = make([]uint64, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseUint64(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesUint32 scans for repeated instances of the specified flag and parses
// the following values as unsigned 32 bit integers. The flags and values are
// removed from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for a uint32, an error is registered.
//
// Returns a slice of the parsed uint32 values.
func (args *Args) ValuesUint32(flag, desc string) []uint32 {
	var (
		matches     []string
		cleanedArgs []string
		result      []uint32
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem uint32
			argErr  error
		)

		result = make([]uint32, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseUint32(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesUint16 scans for repeated instances of the specified flag and parses
// the following values as unsigned 16 bit integers. The flags and values are
// removed from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for a uint16, an error is registered.
//
// Returns a slice of the parsed uint16 values.
func (args *Args) ValuesUint16(flag, desc string) []uint16 {
	var (
		matches     []string
		cleanedArgs []string
		result      []uint16
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem uint16
			argErr  error
		)

		result = make([]uint16, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseUint16(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesUint8 scans for repeated instances of the specified flag and parses
// the following values as unsigned 8 bit integers. The flags and values are
// removed from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for a uint8, an error is registered.
//
// Returns a slice of the parsed uint8 values.
func (args *Args) ValuesUint8(flag, desc string) []uint8 {
	var (
		matches     []string
		cleanedArgs []string
		result      []uint8
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem uint8
			argErr  error
		)

		result = make([]uint8, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseUint8(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesUint scans for repeated instances of the specified flag and parses
// the following values as unsigned integers. The flags and values are removed
// from the argument list.
//
// If any flag lacks a following value, or if a value has invalid syntax or is
// out of range for a uint, an error is registered.
//
// Returns a slice of the parsed uint values.
func (args *Args) ValuesUint(flag, desc string) []uint {
	var (
		matches     []string
		cleanedArgs []string
		result      []uint
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem uint
			argErr  error
		)

		result = make([]uint, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseUint(flag, arg)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}

// ValuesOption scans for repeated instances of the specified flag and
// captures the following values. Each value must appear in the provided list
// of validOptions. The flags and values are removed from the argument list.
//
// If any flag lacks a following value, or if a value is not found in
// validOptions, an error is registered.
//
// Returns a slice of the captured values.
func (args *Args) ValuesOption(
	flag string, validOptions []string, desc string,
) []string {
	var (
		matches     []string
		cleanedArgs []string
		result      []string
		err         error
	)

	args.RegisterUsage(flag, desc)

	matches, cleanedArgs, err = argFlag(flag).values(args.Args())

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem string
			argErr  error
		)

		result = make([]string, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseOption(flag, arg, validOptions)
			if argErr != nil {
				if err == nil {
					err = argErr
				} else {
					err = fmt.Errorf("%w: %w", err, argErr)
				}
			} else {
				result[i] = argItem
			}
		}
	}

	args.args = cleanedArgs
	args.PushErr(err)

	if err == nil {
		return result
	}

	return nil
}
