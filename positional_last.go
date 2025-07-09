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
func Last(name string, args []string) (string, error) {
	value, cleanedArgs, err := Next(name, args)

	if err == nil {
		err = Done(cleanedArgs)
	}

	if err != nil {
		return "", err
	}

	return value, nil
}

// LastString extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastString(name string, args []string) (string, error) {
	return Last(name, args)
}

// LastFloat64 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastFloat64(name string, args []string) (float64, error) {
	var (
		arg    string
		result float64
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseFloat64(name, arg)
	}

	return result, err
}

// LastFloat32 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastFloat32(name string, args []string) (float32, error) {
	var (
		arg    string
		result float32
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseFloat32(name, arg)
	}

	return result, err
}

// LastInt64 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastInt64(name string, args []string) (int64, error) {
	var (
		arg    string
		result int64
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseInt64(name, arg)
	}

	return result, err
}

// LastInt32 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastInt32(name string, args []string) (int32, error) {
	var (
		arg    string
		result int32
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseInt32(name, arg)
	}

	return result, err
}

// LastInt16 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastInt16(name string, args []string) (int16, error) {
	var (
		arg    string
		result int16
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseInt16(name, arg)
	}

	return result, err
}

// LastInt8 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastInt8(name string, args []string) (int8, error) {
	var (
		arg    string
		result int8
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseInt8(name, arg)
	}

	return result, err
}

// LastInt extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastInt(name string, args []string) (int, error) {
	var (
		arg    string
		result int
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseInt(name, arg)
	}

	return result, err
}

// LastUint64 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastUint64(name string, args []string) (uint64, error) {
	var (
		arg    string
		result uint64
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseUint64(name, arg)
	}

	return result, err
}

// LastUint32 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastUint32(name string, args []string) (uint32, error) {
	var (
		arg    string
		result uint32
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseUint32(name, arg)
	}

	return result, err
}

// LastUint16 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastUint16(name string, args []string) (uint16, error) {
	var (
		arg    string
		result uint16
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseUint16(name, arg)
	}

	return result, err
}

// LastUint8 extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastUint8(name string, args []string) (uint8, error) {
	var (
		arg    string
		result uint8
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseUint8(name, arg)
	}

	return result, err
}

// LastUint extracts the next positional argument returning an error if none
// are present or it is not the last, otherwise it parses and returns the
// named data type.
func LastUint(name string, args []string) (uint, error) {
	var (
		arg    string
		result uint
		err    error
	)

	arg, err = Last(name, args)

	if err == nil {
		result, err = parseUint(name, arg)
	}

	return result, err
}
