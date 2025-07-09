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

package szargs_test

import (
	"testing"

	"github.com/dancsecs/szargs"
	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztestlog"
)

const (
	tstEnv = "SZARGS_TESTING_ENVIRONMENT_VARIABLE"
	tstArg = "-t"
)

func TestSzargs_ValueNoArgNoEnv(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Value("def", tstEnv, tstArg, nil)

	chk.Str(value, "def")
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	value, args, err = szargs.Value("def", tstEnv, tstArg,
		[]string{"arg1", "arg2"},
	)

	chk.Str(value, "def")
	chk.StrSlice(args,
		[]string{"arg1", "arg2"},
	)
	chk.NoErr(err)
}

func TestSzargs_ValueArgAmbiguous(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Value("def", tstEnv, tstArg,
		[]string{tstArg, "first", "arg1", "arg2", tstArg, "second"},
	)

	chk.Str(value, "")
	chk.StrSlice(args,
		[]string{tstArg, "first", "arg1", "arg2", tstArg, "second"},
	)
	chk.Err(
		err,
		szargs.ErrAmbiguous.Error()+
			": '-t second' already set to: 'first'",
	)
}

func TestSzargs_ValueArgMissing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Value("def", tstEnv, tstArg,
		[]string{"arg1", "arg2", tstArg},
	)

	chk.Str(value, "")
	chk.StrSlice(args,
		[]string{"arg1", "arg2", tstArg},
	)
	chk.Err(
		err,
		szargs.ErrMissing.Error()+
			": '-t value'",
	)
}

func TestSzargs_ValueEnv(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	chk.SetEnv(tstEnv, "env")

	value, args, err := szargs.Value("def", tstEnv, tstArg,
		[]string{"arg1", "arg2"},
	)

	chk.Str(value, "env")
	chk.StrSlice(args,
		[]string{"arg1", "arg2"},
	)
	chk.NoErr(err)
}

func TestSzargs_ValueEnvAndArg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	chk.SetEnv(tstEnv, "env")

	value, args, err := szargs.Value("def", tstEnv, tstArg,
		[]string{"arg1", tstArg, "arg", "arg2"},
	)

	chk.Str(value, "arg")
	chk.StrSlice(args,
		[]string{"arg1", "arg2"},
	)
	chk.NoErr(err)
}

/*
 ***************************************************************************
 *
 *  Test string setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingString_Default(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.SettingString(
		"testName", "def", tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Str(result, "def")
	chk.StrSlice(args, nil)
}

func TestSzargs_SettingString_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"anotherArg1",
	}

	chk.SetEnv(tstEnv, "envValue")

	result, args, err := szargs.SettingString(
		"testName", "def", tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Str(result, "envValue")
	chk.StrSlice(
		args,
		[]string{
			"309",
			"anotherArg1",
		},
	)
}

func TestSzargs_SettingString_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		tstArg,
		"testValue1",
		"anotherArg1",
	}

	chk.SetEnv(tstEnv, "envValue")

	result, args, err := szargs.SettingString(
		"testName", "def", tstEnv, tstArg, args)

	chk.NoErr(err)
	chk.Str(result, "testValue1")
	chk.StrSlice(
		args,
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
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingFloat64(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Float64(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)   // Argument extracted.
}

func TestSzargs_SettingFloat64_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingFloat64(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Float64(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)   // Argument extracted.
}

func TestSzargs_SettingFloat64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingFloat64(
		"testName", 222.222, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Float64(result, 222.222, 0) // No tolerance.
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "333.333")
	result, args, err = szargs.SettingFloat64(
		"testName", 222.222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Float64(result, 333.333, 0) // No tolerance.
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"444.444",
	}
	result, args, err = szargs.SettingFloat64(
		"testName", 222.222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Float64(result, 444.444, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test float32 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingFloat32_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingFloat32(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Float32(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)   // Argument extracted.
}

func TestSzargs_SettingFloat32_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingFloat32(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Float32(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)   // Argument extracted.
}

func TestSzargs_SettingFloat32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingFloat32(
		"testName", 222.222, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Float32(result, 222.222, 0) // No tolerance.
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "333.333")
	result, args, err = szargs.SettingFloat32(
		"testName", 222.222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Float32(result, 333.333, 0) // No tolerance.
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"444.444",
	}
	result, args, err = szargs.SettingFloat32(
		"testName", 222.222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Float32(result, 444.444, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test int64 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingInt64_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingInt64(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Int64(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingInt64_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingInt64(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Int64(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingInt64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingInt64(
		"testName", 222, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Int64(result, 222)
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result, args, err = szargs.SettingInt64(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Int64(result, 333)
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"444",
	}
	result, args, err = szargs.SettingInt64(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Int64(result, 444)
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test int32 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingInt32_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingInt32(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Int32(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingInt32_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingInt32(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Int32(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingInt32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingInt32(
		"testName", 222, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Int32(result, 222)
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result, args, err = szargs.SettingInt32(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Int32(result, 333)
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"444",
	}
	result, args, err = szargs.SettingInt32(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Int32(result, 444)
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test int16 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingInt16_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingInt16(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Int16(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingInt16_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingInt16(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Int16(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingInt16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingInt16(
		"testName", 222, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Int16(result, 222)
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result, args, err = szargs.SettingInt16(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Int16(result, 333)
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"444",
	}
	result, args, err = szargs.SettingInt16(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Int16(result, 444)
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test int8 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingInt8_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingInt8(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Int8(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingInt8_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingInt8(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Int8(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingInt8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingInt8(
		"testName", 22, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Int8(result, 22)
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "33")
	result, args, err = szargs.SettingInt8(
		"testName", 22, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Int8(result, 33)
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"44",
	}
	result, args, err = szargs.SettingInt8(
		"testName", 22, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Int8(result, 44)
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test int setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingInt_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingInt(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Int(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingInt_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingInt(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Int(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingInt_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingInt(
		"testName", 222, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Int(result, 222)
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result, args, err = szargs.SettingInt(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Int(result, 333)
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"444",
	}
	result, args, err = szargs.SettingInt(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Int(result, 444)
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test uint64 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingUint64_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingUint64(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Uint64(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingUint64_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingUint64(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Uint64(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingUint64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingUint64(
		"testName", 222, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Uint64(result, 222)
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result, args, err = szargs.SettingUint64(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Uint64(result, 333)
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"444",
	}
	result, args, err = szargs.SettingUint64(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Uint64(result, 444)
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test uint32 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingUint32_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingUint32(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Uint32(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingUint32_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingUint32(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Uint32(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingUint32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingUint32(
		"testName", 222, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Uint32(result, 222)
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result, args, err = szargs.SettingUint32(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Uint32(result, 333)
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"444",
	}
	result, args, err = szargs.SettingUint32(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Uint32(result, 444)
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test uint16 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingUint16_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingUint16(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Uint16(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingUint16_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingUint16(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Uint16(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingUint16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingUint16(
		"testName", 222, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Uint16(result, 222)
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result, args, err = szargs.SettingUint16(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Uint16(result, 333)
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"444",
	}
	result, args, err = szargs.SettingUint16(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Uint16(result, 444)
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test uint8 setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingUint8_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingUint8(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Uint8(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingUint8_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingUint8(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Uint8(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingUint8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingUint8(
		"testName", 22, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Uint8(result, 22)
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "33")
	result, args, err = szargs.SettingUint8(
		"testName", 22, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Uint8(result, 33)
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"44",
	}
	result, args, err = szargs.SettingUint8(
		"testName", 22, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Uint8(result, 44)
	chk.StrSlice(args, nil)
}

/*
 ***************************************************************************
 *
 *  Test uint setting.
 *
 ***************************************************************************
 */

func TestSzargs_SettingUint_Invalid_Arg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
	}

	result, args, err := szargs.SettingUint(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrSyntax,
			"testName",
			"'notANumber1'",
		),
	)
	chk.Uint(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingUint_Invalid_Env(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{}

	chk.SetEnv(tstEnv, "notANumber")

	result, args, err := szargs.SettingUint(
		"testName", 123, tstEnv, tstArg, args,
	)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrSyntax,
			"testName",
			"'notANumber'",
		),
	)
	chk.Uint(result, 0)
	chk.StrSlice(args, nil) // Argument extracted.
}

func TestSzargs_SettingUint_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	// Default.
	result, args, err := szargs.SettingUint(
		"testName", 222, tstEnv, tstArg, nil,
	)

	chk.NoErr(err)
	chk.Uint(result, 222)
	chk.StrSlice(args, nil)

	// Environment.
	chk.SetEnv(tstEnv, "333")
	result, args, err = szargs.SettingUint(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Uint(result, 333)
	chk.StrSlice(args, nil)

	// Argument
	args = []string{
		"-t",
		"444",
	}
	result, args, err = szargs.SettingUint(
		"testName", 222, tstEnv, tstArg, args,
	)

	chk.NoErr(err)
	chk.Uint(result, 444)
	chk.StrSlice(args, nil)
}
