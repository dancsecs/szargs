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

	value, found, cleanArgs, err := Flag(flag).value(args)
	if err != nil {
		return "", args, srcErr, err
	}

	if !found {
		value = def
		srcErr = nil

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

// SettingString scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingFloat64 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingFloat32 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingInt64 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingInt32 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingInt16 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingInt8 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingInt scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingUint64 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingUint32 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingUint16 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingUint8 scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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

// SettingUint scans the args looking for arg.  If it is not found then it
// looks for an environment variable and if this does not exist then it will
// return the specified default.
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
