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

// ValuesString scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesString(flag, desc string) []string {
	args.addUsage(flag, desc)
	result, newArgs, err := Flag(flag).values(args.args)
	args.args = newArgs
	args.PushErr(err)

	return result
}

// ValuesFloat64 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesFloat64(flag, desc string) []float64 {
	var (
		matches     []string
		cleanedArgs []string
		result      []float64
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesFloat32 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesFloat32(flag, desc string) []float32 {
	var (
		matches     []string
		cleanedArgs []string
		result      []float32
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesInt64 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesInt64(flag, desc string) []int64 {
	var (
		matches     []string
		cleanedArgs []string
		result      []int64
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesInt32 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesInt32(flag, desc string) []int32 {
	var (
		matches     []string
		cleanedArgs []string
		result      []int32
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesInt16 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesInt16(flag, desc string) []int16 {
	var (
		matches     []string
		cleanedArgs []string
		result      []int16
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesInt8 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesInt8(flag, desc string) []int8 {
	var (
		matches     []string
		cleanedArgs []string
		result      []int8
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesInt scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesInt(flag, desc string) []int {
	var (
		matches     []string
		cleanedArgs []string
		result      []int
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesUint64 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesUint64(flag, desc string) []uint64 {
	var (
		matches     []string
		cleanedArgs []string
		result      []uint64
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesUint32 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesUint32(flag, desc string) []uint32 {
	var (
		matches     []string
		cleanedArgs []string
		result      []uint32
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesUint16 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesUint16(flag, desc string) []uint16 {
	var (
		matches     []string
		cleanedArgs []string
		result      []uint16
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesUint8 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesUint8(flag, desc string) []uint8 {
	var (
		matches     []string
		cleanedArgs []string
		result      []uint8
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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

// ValuesUint scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (args *Args) ValuesUint(flag, desc string) []uint {
	var (
		matches     []string
		cleanedArgs []string
		result      []uint
		err         error
	)

	args.addUsage(flag, desc)

	matches, cleanedArgs, err = Flag(flag).values(args.Args())

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
