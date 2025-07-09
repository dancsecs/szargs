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
	"os"
)

const defaultStandIn = "~~--default--~~"

// Value is a convenience function that selects a values from a default that
// can be overridden by an optional environment variable or an optional
// argument.  If the argument is found its value is chosen and removed from
// the returned argument list.  An error is returned if the argument is missing
// or ambiguous.
func Value(
	defaultValue, envOverride, argOverride string,
	args []string,
) (string, []string, error) {
	value, found, cleanArgs, err := Arg(argOverride).Value(args)
	if err != nil {
		return "", args, err
	}

	if !found {
		value = defaultValue

		if envOverride != "" {
			envValue, ok := os.LookupEnv(envOverride)
			if ok {
				value = envValue
			}
		}
	}

	return value, cleanArgs, nil
}

// SettingString scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingString(
	_ /*name*/, def, env, arg string, args []string,
) (string, []string, error) {
	return Value(def, env, arg, args)
}

// SettingFloat64 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingFloat64(
	name string, def float64, env, arg string, args []string,
) (float64, []string, error) {
	var (
		value  string
		result float64
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseFloat64(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingFloat32 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingFloat32(
	name string, def float32, env, arg string, args []string,
) (float32, []string, error) {
	var (
		value  string
		result float32
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseFloat32(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingInt64 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingInt64(
	name string, def int64, env, arg string, args []string,
) (int64, []string, error) {
	var (
		value  string
		result int64
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseInt64(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingInt32 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingInt32(
	name string, def int32, env, arg string, args []string,
) (int32, []string, error) {
	var (
		value  string
		result int32
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseInt32(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingInt16 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingInt16(
	name string, def int16, env, arg string, args []string,
) (int16, []string, error) {
	var (
		value  string
		result int16
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseInt16(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingInt8 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingInt8(
	name string, def int8, env, arg string, args []string,
) (int8, []string, error) {
	var (
		value  string
		result int8
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseInt8(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingInt scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingInt(
	name string, def int, env, arg string, args []string,
) (int, []string, error) {
	var (
		value  string
		result int
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseInt(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingUint64 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingUint64(
	name string, def uint64, env, arg string, args []string,
) (uint64, []string, error) {
	var (
		value  string
		result uint64
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseUint64(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingUint32 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingUint32(
	name string, def uint32, env, arg string, args []string,
) (uint32, []string, error) {
	var (
		value  string
		result uint32
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseUint32(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingUint16 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingUint16(
	name string, def uint16, env, arg string, args []string,
) (uint16, []string, error) {
	var (
		value  string
		result uint16
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseUint16(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingUint8 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingUint8(
	name string, def uint8, env, arg string, args []string,
) (uint8, []string, error) {
	var (
		value  string
		result uint8
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseUint8(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}

// SettingUint scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
func SettingUint(
	name string, def uint, env, arg string, args []string,
) (uint, []string, error) {
	var (
		value  string
		result uint
		err    error
	)

	value, args, err = Value(defaultStandIn, env, arg, args)

	if err == nil {
		if value == defaultStandIn {
			result = def
		} else {
			result, err = parseUint(name, value)
		}
	}

	if err == nil {
		return result, args, nil
	}

	return result, args, err
}
