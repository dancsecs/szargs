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

// Values scans the args looking for all instances of the specified flag.  If
// it finds it then the next arg as the value absorbing both the flag the
// value from the argument list.  If there is no next arg an error is returned
// and the original arg array is returned.
func (a Arg) Values(args []string) ([]string, []string, error) {
	values := []string(nil)
	cleanedArgs := make([]string, 0, len(args))

	for i, mi := 0, len(args); i < mi; i++ {
		if a.argIs(args[i]) {
			if (i + 1) >= mi {
				return nil, args,
					fmt.Errorf(
						"%w: '%s value'",
						ErrMissing,
						a,
					)
			}

			i++
			values = append(values, args[i])
		} else {
			cleanedArgs = append(cleanedArgs, args[i])
		}
	}

	return values, cleanedArgs, nil
}

// ValuesString scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesString(args []string) ([]string, []string, error) {
	return a.Values(args)
}

// ValuesFloat64 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesFloat64(
	args []string,
) ([]float64, []string, error) {
	var (
		matches []string
		result  []float64
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem float64
			argErr  error
		)

		result = make([]float64, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseFloat64(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesFloat32 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesFloat32(args []string) ([]float32, []string, error) {
	var (
		matches []string
		result  []float32
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem float32
			argErr  error
		)

		result = make([]float32, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseFloat32(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesInt64 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesInt64(args []string) ([]int64, []string, error) {
	var (
		matches []string
		result  []int64
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem int64
			argErr  error
		)

		result = make([]int64, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseInt64(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesInt32 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesInt32(args []string) ([]int32, []string, error) {
	var (
		matches []string
		result  []int32
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem int32
			argErr  error
		)

		result = make([]int32, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseInt32(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesInt16 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesInt16(args []string) ([]int16, []string, error) {
	var (
		matches []string
		result  []int16
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem int16
			argErr  error
		)

		result = make([]int16, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseInt16(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesInt8 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesInt8(args []string) ([]int8, []string, error) {
	var (
		matches []string
		result  []int8
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem int8
			argErr  error
		)

		result = make([]int8, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseInt8(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesInt scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesInt(args []string) ([]int, []string, error) {
	var (
		matches []string
		result  []int
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem int
			argErr  error
		)

		result = make([]int, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseInt(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesUint64 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesUint64(args []string) ([]uint64, []string, error) {
	var (
		matches []string
		result  []uint64
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem uint64
			argErr  error
		)

		result = make([]uint64, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseUint64(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesUint32 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesUint32(args []string) ([]uint32, []string, error) {
	var (
		matches []string
		result  []uint32
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem uint32
			argErr  error
		)

		result = make([]uint32, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseUint32(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesUint16 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesUint16(args []string) ([]uint16, []string, error) {
	var (
		matches []string
		result  []uint16
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem uint16
			argErr  error
		)

		result = make([]uint16, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseUint16(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesUint8 scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesUint8(args []string) ([]uint8, []string, error) {
	var (
		matches []string
		result  []uint8
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem uint8
			argErr  error
		)

		result = make([]uint8, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseUint8(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}

// ValuesUint scans the args looking for all instances of the specified flag
// returning all found in a typed slice.
func (a Arg) ValuesUint(args []string) ([]uint, []string, error) {
	var (
		matches []string
		result  []uint
		err     error
	)

	matches, args, err = a.Values(args)

	if err == nil { //nolint:nestif // Ok.
		var (
			argItem uint
			argErr  error
		)

		result = make([]uint, len(matches))

		for i, arg := range matches {
			argItem, argErr = parseUint(string(a), arg)
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

	if err == nil {
		return result, args, nil
	}

	return nil, args, err
}
