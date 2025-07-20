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
	"errors"
	"os"
	"slices"
	"strings"
)

const defaultStandIn = "~~--default--~~"

// Setting is a convenience function that selects a values from a default that
// can be overridden by an optional environment variable or an optional
// argument.  If the argument is found its value is chosen and removed from
// the returned argument list.  An error is returned if the argument is missing
// or ambiguous.
func setting(
	flag, env, def string, args []string,
) (string, []string, error, error) {
	srcErr := ErrInvalidFlag

	value, found, cleanArgs, err := argFlag(flag).value(args)
	if err != nil {
		return "", args, srcErr, err
	}

	if !found {
		value = def
		srcErr = ErrInvalidDefault

		if env != "" {
			envValue, ok := os.LookupEnv(env)
			if ok {
				value = envValue
				srcErr = ErrInvalidEnv
			}
		}
	}

	return value, cleanArgs, srcErr, nil
}

// SettingString returns a configuration value based on a default,
// optionally overridden by an environment variable, and further overridden
// by a flagged command-line argument.
//
// Returns the final selected string value.
func (args *Args) SettingString(
	flag, env, def, desc string,
) string {
	args.addUsage(flag, desc)
	result, cleanedArgs, srcErr, err := setting(flag, env, def, args.Args())
	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingFloat64 returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as a 64 bit floating point
// number.
//
// If the final value has invalid syntax or is out of range for a float64, an
// error is registered.
//
// Returns the final parsed float64 value.
func (args *Args) SettingFloat64(
	flag, env string, def float64, desc string,
) float64 {
	var (
		value       string
		cleanedArgs []string
		result      float64
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseFloat64(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingFloat32 returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as a 32 bit floating point
// number.
//
// If the final value has invalid syntax or is out of range for a float32, an
// error is registered.
//
// Returns the final parsed float32 value.
func (args *Args) SettingFloat32(
	flag, env string, def float32, desc string,
) float32 {
	var (
		value       string
		cleanedArgs []string
		result      float32
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseFloat32(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingInt64 returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as a signed 64 bit integer.
//
// If the final value has invalid syntax or is out of range for an int64, an
// error is registered.
//
// Returns the final parsed int64 value.
func (args *Args) SettingInt64(
	flag, env string, def int64, desc string,
) int64 {
	var (
		value       string
		cleanedArgs []string
		result      int64
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseInt64(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingInt32 returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as a signed 32 bit integer.
//
// If the final value has invalid syntax or is out of range for an int32, an
// error is registered.
//
// Returns the final parsed int32 value.
func (args *Args) SettingInt32(
	flag, env string, def int32, desc string,
) int32 {
	var (
		value       string
		cleanedArgs []string
		result      int32
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseInt32(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingInt16 returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as a signed 16 bit integer.
//
// If the final value has invalid syntax or is out of range for an int16, an
// error is registered.
//
// Returns the final parsed int16 value.
func (args *Args) SettingInt16(
	flag, env string, def int16, desc string,
) int16 {
	var (
		value       string
		cleanedArgs []string
		result      int16
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseInt16(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingInt8 returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as a signed 8 bit integer.
//
// If the final value has invalid syntax or is out of range for an int8, an
// error is registered.
//
// Returns the final parsed int8 value.
func (args *Args) SettingInt8(
	flag, env string, def int8, desc string,
) int8 {
	var (
		value       string
		cleanedArgs []string
		result      int8
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseInt8(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingInt returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as a signed integer.
//
// If the final value has invalid syntax or is out of range for an int, an
// error is registered.
//
// Returns the final parsed int value.
func (args *Args) SettingInt(
	flag, env string, def int, desc string,
) int {
	var (
		value       string
		cleanedArgs []string
		result      int
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseInt(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingUint64 returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as an unsigned 64 bit integer.
//
// If the final value has invalid syntax or is out of range for a uint64, an
// error is registered.
//
// Returns the final parsed uint64 value.
func (args *Args) SettingUint64(
	flag, env string, def uint64, desc string,
) uint64 {
	var (
		value       string
		cleanedArgs []string
		result      uint64
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseUint64(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingUint32 returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as an unsigned 32 bit integer.
//
// If the final value has invalid syntax or is out of range for a uint32, an
// error is registered.
//
// Returns the final parsed uint32 value.
func (args *Args) SettingUint32(
	flag, env string, def uint32, desc string,
) uint32 {
	var (
		value       string
		cleanedArgs []string
		result      uint32
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseUint32(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingUint16 returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as an unsigned 16 bit integer.
//
// If the final value has invalid syntax or is out of range for a uint16, an
// error is registered.
//
// Returns the final parsed uint16 value.
func (args *Args) SettingUint16(
	flag, env string, def uint16, desc string,
) uint16 {
	var (
		value       string
		cleanedArgs []string
		result      uint16
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseUint16(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingUint8 returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as an unsigned 8 bit integer.
//
// If the final value has invalid syntax or is out of range for a uint8, an
// error is registered.
//
// Returns the final parsed uint8 value.
func (args *Args) SettingUint8(
	flag, env string, def uint8, desc string,
) uint8 {
	var (
		value       string
		cleanedArgs []string
		result      uint8
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseUint8(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingUint returns a configuration value based on a default, optionally
// overridden by an environment variable, and further overridden by a flagged
// command-line argument. The value is parsed as an unsigned integer.
//
// If the final value has invalid syntax or is out of range for a uint, an
// error is registered.
//
// Returns the final parsed uint value.
func (args *Args) SettingUint(
	flag, env string, def uint, desc string,
) uint {
	var (
		value       string
		cleanedArgs []string
		result      uint
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			result = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}

			result, err = parseUint(parseName, value)
		}
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingOption returns a configuration value based on a default,
// optionally overridden by an environment variable, and further overridden
// by a flagged command-line argument.
//
// If the final value is not found in the list of validOptions,
// an error is registered.
//
// Returns the final selected value.
func (args *Args) SettingOption(
	flag, env string, def string, validOptions []string, desc string,
) string {
	var (
		value       string
		cleanedArgs []string
		result      string
		parseName   string
		srcErr      error
		err         error
	)

	args.addUsage(flag, desc)

	value, cleanedArgs, srcErr, err = setting(
		flag, env, defaultStandIn, args.Args(),
	)

	if err == nil { //nolint:nestif // Ok.
		if value == defaultStandIn {
			parseName = "default"
			value = def
		} else {
			if errors.Is(srcErr, ErrInvalidEnv) {
				parseName = env
			} else {
				parseName = flag
			}
		}

		result, err = parseOption(parseName, value, validOptions)
	}

	args.args = cleanedArgs

	if err != nil {
		args.PushErr(srcErr)
		args.PushErr(err)
	}

	return result
}

// SettingIs returns true if a specified environment variable is set to a
// truthy value, or if a corresponding boolean command-line flag is present.
//
// Unlike other Setting methods, there is no default.
//
// The environment variable is considered true if it is set to one of: "",
// "T", "Y", "TRUE", "YES", "ON" or "1" (case-insensitive). Any other value is
// considered false.
//
// The command-line flag override takes no value—its presence alone indicates
// true.
//
// Returns the resulting boolean value.
func (args *Args) SettingIs(flag, env string, desc string) bool {
	var (
		value  string
		result bool
	)

	args.addUsage(flag, desc)

	result = args.Is(flag, desc)

	if !args.HasErr() && !result && env != "" {
		envValue, ok := os.LookupEnv(env)
		if ok {
			value = strings.ToLower(envValue)
			result = slices.Contains(
				[]string{
					"",
					"t",
					"true",
					"y",
					"yes",
					"on",
					"1",
				},
				value,
			)
		}
	}

	return result
}
