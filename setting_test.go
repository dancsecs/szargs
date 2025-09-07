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

//nolint:dupl // Ok.
package szargs_test

import (
	"testing"

	"github.com/dancsecs/szargs"
	"github.com/dancsecs/sztestlog"
)

const (
	tstEnv     = "SZARGS_TESTING_ENVIRONMENT_VARIABLE"
	tstArgFlag = "[-t value]"
	tstArg     = "-t"
)

/*
 ***************************************************************************
 *
 *  Test string setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingString_MissingArgError(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		tstArg,
	})

	result := args.SettingString(
		tstArgFlag, tstEnv, "def", "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrMissing,
			"'"+tstArgFlag+"'",
		),
	)
	chk.Str(result, "")
	chk.StrSlice(args.Args(), []string{tstArg})
}

func TestSzargs_SettingString_Default(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.SettingString(
		tstArgFlag, tstEnv, "def", "testName",
	)

	chk.NoErr(args.Err())
	chk.Str(result, "def")
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_SettingString_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg1",
	})

	chk.SetEnv(tstEnv, "envValue")

	result := args.SettingString(
		tstArgFlag, tstEnv, "def", "testName",
	)

	chk.NoErr(args.Err())
	chk.Str(result, "envValue")
	chk.StrSlice(
		args.Args(),
		[]string{
			"309",
			"anotherArg1",
		},
	)
}

func TestSzargs_SettingString_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		tstArg,
		"testValue1",
		"anotherArg1",
	})

	chk.SetEnv(tstEnv, "envValue")

	result := args.SettingString(
		tstArgFlag, tstEnv, "def", "testName",
	)

	chk.NoErr(args.Err())
	chk.Str(result, "testValue1")
	chk.StrSlice(
		args.Args(),
		[]string{
			"309",
			"anotherArg1",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test float64 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingFloat64_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingFloat64(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Float64(result, 0, 0)      // No tolerance.
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingFloat64_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingFloat64(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Float64(result, 0, 0)      // No tolerance.
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingFloat64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingFloat64(
		tstArgFlag, tstEnv, 222.222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Float64(result, 222.222, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "333.333")
	result = args.SettingFloat64(
		tstArgFlag, tstEnv, 222.222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Float64(result, 333.333, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"444.444",
	})

	result = args.SettingFloat64(
		tstArgFlag, tstEnv, 222.222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Float64(result, 444.444, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test float32 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingFloat32_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingFloat32(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidFloat32,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Float32(result, 0, 0)      // No tolerance.
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingFloat32_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingFloat32(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidFloat32,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Float32(result, 0, 0)      // No tolerance.
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingFloat32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingFloat32(
		tstArgFlag, tstEnv, 222.222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Float32(result, 222.222, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "333.333")
	result = args.SettingFloat32(
		tstArgFlag, tstEnv, 222.222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Float32(result, 333.333, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"444.444",
	})
	result = args.SettingFloat32(
		tstArgFlag, tstEnv, 222.222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Float32(result, 444.444, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test int64 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingInt64_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingInt64(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidInt64,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Int64(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingInt64_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingInt64(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidInt64,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Int64(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingInt64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingInt64(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int64(result, 222)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result = args.SettingInt64(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int64(result, 333)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"444",
	})
	result = args.SettingInt64(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int64(result, 444)
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test int32 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingInt32_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingInt32(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidInt32,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Int32(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingInt32_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingInt32(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidInt32,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Int32(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingInt32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingInt32(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int32(result, 222)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result = args.SettingInt32(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int32(result, 333)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"444",
	})
	result = args.SettingInt32(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int32(result, 444)
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test int16 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingInt16_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingInt16(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidInt16,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Int16(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingInt16_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingInt16(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidInt16,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Int16(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingInt16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingInt16(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int16(result, 222)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result = args.SettingInt16(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int16(result, 333)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"444",
	})
	result = args.SettingInt16(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int16(result, 444)
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test int8 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingInt8_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingInt8(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidInt8,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Int8(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingInt8_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingInt8(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidInt8,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Int8(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingInt8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingInt8(
		tstArgFlag, tstEnv, 22, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int8(result, 22)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "33")
	result = args.SettingInt8(
		tstArgFlag, tstEnv, 22, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int8(result, 33)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"44",
	})
	result = args.SettingInt8(
		tstArgFlag, tstEnv, 22, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int8(result, 44)
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test int setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingInt_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingInt(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidInt,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Int(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingInt_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingInt(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidInt,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Int(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingInt_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingInt(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int(result, 222)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result = args.SettingInt(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int(result, 333)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"444",
	})
	result = args.SettingInt(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Int(result, 444)
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test uint64 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingUint64_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingUint64(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidUint64,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Uint64(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingUint64_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingUint64(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidUint64,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Uint64(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingUint64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingUint64(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint64(result, 222)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result = args.SettingUint64(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint64(result, 333)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"444",
	})
	result = args.SettingUint64(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint64(result, 444)
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test uint32 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingUint32_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingUint32(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidUint32,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Uint32(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingUint32_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingUint32(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidUint32,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Uint32(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingUint32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingUint32(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint32(result, 222)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result = args.SettingUint32(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint32(result, 333)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"444",
	})
	result = args.SettingUint32(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint32(result, 444)
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test uint16 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingUint16_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingUint16(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidUint16,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Uint16(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingUint16_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingUint16(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidUint16,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Uint16(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingUint16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingUint16(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint16(result, 222)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result = args.SettingUint16(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint16(result, 333)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"444",
	})
	result = args.SettingUint16(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint16(result, 444)
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test uint8 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingUint8_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingUint8(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidUint8,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Uint8(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingUint8_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingUint8(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidUint8,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Uint8(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingUint8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingUint8(
		tstArgFlag, tstEnv, 22, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint8(result, 22)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "33")
	result = args.SettingUint8(
		tstArgFlag, tstEnv, 22, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint8(result, 33)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"44",
	})
	result = args.SettingUint8(
		tstArgFlag, tstEnv, 22, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint8(result, 44)
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test uint setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingUint_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber1",
	})

	result := args.SettingUint(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidUint,
			szargs.ErrSyntax,
			tstArgFlag,
			"'notANumber1'",
		),
	)
	chk.Uint(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingUint_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notANumber")

	result := args.SettingUint(
		tstArgFlag, tstEnv, 123, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidUint,
			szargs.ErrSyntax,
			tstEnv,
			"'notANumber'",
		),
	)
	chk.Uint(result, 0)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingUint_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingUint(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint(result, 222)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result = args.SettingUint(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint(result, 333)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"444",
	})
	result = args.SettingUint(
		tstArgFlag, tstEnv, 222, "testName",
	)

	chk.NoErr(args.Err())
	chk.Uint(result, 444)
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test option setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingOption_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notAnOption",
	})

	result := args.SettingOption(
		tstArgFlag, tstEnv, "abc", []string{"abc", "def"}, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFlag,
			szargs.ErrInvalidOption,
			"'notAnOption' "+
				"([-t value] must be one of [abc def])",
		),
	)
	chk.Str(result, "")
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingOption_Invalid_Default(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.SettingOption(
		tstArgFlag, tstEnv, "notAnOption", []string{"abc", "def"}, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidDefault,
			szargs.ErrInvalidOption,
			"'notAnOption' "+
				"(default must be one of [abc def])",
		),
	)
	chk.Str(result, "")
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingOption_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	chk.SetEnv(tstEnv, "notAnOption")

	result := args.SettingOption(
		tstArgFlag, tstEnv, "abc", []string{"abc", "def"}, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidEnv,
			szargs.ErrInvalidOption,
			"'notAnOption' ("+
				tstEnv+" must be one of [abc def])",
		),
	)
	chk.Str(result, "")
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingOption_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingOption(
		tstArgFlag, tstEnv, "abc", []string{"abc", "def", "ghi"}, "testName",
	)

	chk.NoErr(args.Err())
	chk.Str(result, "abc")
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "def")
	result = args.SettingOption(
		tstArgFlag, tstEnv, "abc", []string{"abc", "def", "ghi"}, "testName",
	)

	chk.NoErr(args.Err())
	chk.Str(result, "def")
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
		"ghi",
	})
	result = args.SettingOption(
		tstArgFlag, tstEnv, "abc", []string{"abc", "def", "ghi"}, "testName",
	)

	chk.NoErr(args.Err())
	chk.Str(result, "ghi")
	chk.StrSlice(args.Args(), nil)
}

/*
 ***************************************************************************
 *
 *  Test is setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingIs_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"-t",
	})

	result := args.SettingIs(
		"[-t]", tstEnv, "testName",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrAmbiguous,
			"'[-t]' found 2 times",
		),
	)
	chk.False(result)
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_SettingIs_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	// Default.
	result := args.SettingIs(
		"[-t]", tstEnv, "testName",
	)

	chk.NoErr(args.Err())
	chk.False(result)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "TRUE")
	result = args.SettingIs(
		"[-t]", tstEnv, "testName",
	)

	chk.NoErr(args.Err())
	chk.True(result)
	chk.StrSlice(args.Args(), nil)

	// Environment.
	chk.SetEnv(tstEnv, "False")
	result = args.SettingIs(
		"[-t]", tstEnv, "testName",
	)

	chk.NoErr(args.Err())
	chk.False(result)
	chk.StrSlice(args.Args(), nil)

	// Argument
	args = szargs.New("program description", []string{
		"programName",
		"-t",
	})
	result = args.SettingIs(
		"[-t]", tstEnv, "testName",
	)

	chk.NoErr(args.Err())
	chk.True(result)
	chk.StrSlice(args.Args(), nil)
}
