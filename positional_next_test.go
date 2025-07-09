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
	"math"
	"testing"

	"github.com/dancsecs/szargs"
	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztestlog"
)

func TestSzargsPositional_Next(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	var (
		arg  string
		args []string
		err  error
	)

	arg, args, err = szargs.Next("TestArg", args)

	chk.Str(arg, "")
	chk.StrSlice(args, nil)
	chk.Err(
		err,
		"missing argument: TestArg",
	)

	args = []string{"arg1", "arg2"}

	arg, args, err = szargs.Next("TestArg", args)

	chk.Str(arg, "arg1")
	chk.StrSlice(args, []string{"arg2"})
	chk.NoErr(err)

	arg, args, err = szargs.Next("TestArg", args)

	chk.Str(arg, "arg2")
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	arg, args, err = szargs.Next("TestArg", args)

	chk.Str(arg, "")
	chk.StrSlice(args, nil)
	chk.Err(
		err,
		szargs.ErrMissing.Error()+": TestArg",
	)
}

/*
 ***************************************************************************
 *
 *  Test string positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextString_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextString("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(szargs.ErrMissing, "TestArg"),
	)
	chk.Str(result, "")
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextString_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"anotherArg",
	}

	result, args, err := szargs.NextString("TestArg", args)

	chk.NoErr(err)
	chk.Str(result, "309")
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test args.float64 positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextFloat64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextFloat64("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Float64(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextFloat64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextFloat64("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Float64(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)   // Argument extracted.
}

func TestSzargsPositional_NextFloat64_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"1.7e+309", // MaxFloat64 * 10 is out of range.
	}

	result, args, err := szargs.NextFloat64("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrRange,
			"TestArg",
			"'1.7e+309'",
		),
	)
	chk.Float64(result, math.Inf(1), 0) // No tolerance.
	chk.StrSlice(args, nil)             // Argument extracted.
}

func TestSzargsPositional_NextFloat64_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-1.7e+309", // MinFloat64 * 10 is out of range.
	}

	result, args, err := szargs.NextFloat64("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrRange,
			"TestArg",
			"'-1.7e+309'",
		),
	)
	chk.Float64(result, math.Inf(-1), 0) // No tolerance.
	chk.StrSlice(args, nil)              // Argument extracted.
}

func TestSzargsPositional_NextFloat64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309.2",
		"anotherArg",
	}

	result, args, err := szargs.NextFloat64("TestArg", args)

	chk.NoErr(err)
	chk.Float64(result, 309.2, 0) // No tolerance.
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test float32 positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextFloat32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextFloat32("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Float32(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextFloat32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextFloat32("TestArg", args)

	// Parsing syntax error.
	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Float32(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextFloat32_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"1.7e+309",
	}

	result, args, err := szargs.NextFloat32("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrRange,
			"TestArg",
			"'1.7e+309'",
		),
	)
	chk.Float32(result, float32(math.Inf(1)), 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextFloat32_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-1.7e+309",
	}

	result, args, err := szargs.NextFloat32("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrRange,
			"TestArg",
			"'-1.7e+309'",
		),
	)
	chk.Float32(result, float32(math.Inf(-1)), 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextFloat32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309.2",
		"anotherArg",
	}

	result, args, err := szargs.NextFloat32("TestArg", args)

	chk.NoErr(err)
	chk.Float32(result, 309.2, 0) // No tolerance.
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test int64 positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextInt64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextInt64("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int64(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextInt64("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)

	chk.Int64(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt64_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"9223372036854775808", // MaxInt64 + 1 is out of range.
	}

	result, args, err := szargs.NextInt64("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrRange,
			"TestArg",
			"'9223372036854775808'",
		),
	)
	chk.Int64(result, math.MaxInt64)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt64_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-9223372036854775809", // MinInt64 - 1 is out of range.
	}

	result, args, err := szargs.NextInt64("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrRange,
			"TestArg",
			"'-9223372036854775809'",
		),
	)
	chk.Int64(result, math.MinInt64)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"anotherArg",
	}

	result, args, err := szargs.NextInt64("TestArg", args)

	// No Error
	chk.NoErr(err)
	chk.Int64(result, 309)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test int32 positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextInt32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextInt32("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int32(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextInt32("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Int32(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt32_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"2147483648", // MaxInt32 + 1 is out of range.
	}

	result, args, err := szargs.NextInt32("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrRange,
			"TestArg",
			"'2147483648'",
		),
	)
	chk.Int32(result, math.MaxInt32)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt32_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-2147483649", // MinInt32 - 1 is out of range.
	}

	result, args, err := szargs.NextInt32("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrRange,
			"TestArg",
			"'-2147483649'",
		),
	)
	chk.Int32(result, math.MinInt32)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"anotherArg",
	}

	result, args, err := szargs.NextInt32("TestArg", args)

	chk.NoErr(err)
	chk.Int32(result, 309)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test int16 positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextInt16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextInt16("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int16(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt16_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextInt16("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Int16(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt16_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"32768", // MaxInt16 + 1 is out of range.
	}

	result, args, err := szargs.NextInt16("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrRange,
			"TestArg",
			"'32768'",
		),
	)
	chk.Int16(result, math.MaxInt16)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt16_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-32769", // MainInt16 - 1 is out of range.
	}

	result, args, err := szargs.NextInt16("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrRange,
			"TestArg",
			"'-32769'",
		),
	)
	chk.Int16(result, math.MinInt16)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"anotherArg",
	}

	result, args, err := szargs.NextInt16("TestArg", args)

	chk.NoErr(err)
	chk.Int16(result, 309)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test int8 positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextInt8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextInt8("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int8(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt8_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextInt8("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Int8(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt8_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"128", // MaxInt8 + 1 is out of range.
	}

	result, args, err := szargs.NextInt8("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrRange,
			"TestArg",
			"'128'",
		),
	)
	chk.Int8(result, math.MaxInt8)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt8_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-129", // MinInt8 - 1 is out of range.
	}

	result, args, err := szargs.NextInt8("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrRange,
			"TestArg",
			"'-129'",
		),
	)
	chk.Int8(result, math.MinInt8)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"109",
		"anotherArg",
	}

	result, args, err := szargs.NextInt8("TestArg", args)

	chk.NoErr(err)
	chk.Int8(result, 109)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test int positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextInt_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextInt("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextInt("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Int(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"9223372036854775808", // MaxInt + 1 is out of range.
	}

	result, args, err := szargs.NextInt("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrRange,
			"TestArg",
			"'9223372036854775808'",
		),
	)
	chk.Int(result, math.MaxInt)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-9223372036854775809", // MinInt - 1 is out of range.
	}

	result, args, err := szargs.NextInt("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrRange,
			"TestArg",
			"'-9223372036854775809'",
		),
	)
	chk.Int(result, math.MinInt)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextInt_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"anotherArg",
	}

	result, args, err := szargs.NextInt("TestArg", args)

	chk.NoErr(err)
	chk.Int(result, 309)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test uint64 positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextUint64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextUint64("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint64(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextUint64("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Uint64(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint64_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"18446744073709551616", // MaxUint64 + 1 is out of range.
	}

	result, args, err := szargs.NextUint64("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrRange,
			"TestArg",
			"'18446744073709551616'",
		),
	)
	chk.Uint64(result, math.MaxUint64)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"anotherArg",
	}

	result, args, err := szargs.NextUint64("TestArg", args)

	chk.NoErr(err)
	chk.Uint64(result, 309)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test uint32 positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextUint32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextUint32("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint32(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextUint32("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Uint32(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint32_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"4294967296", // MaxUint32 + 1 is out of range.
	}
	result, args, err := szargs.NextUint32("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrRange,
			"TestArg",
			"'4294967296'",
		),
	)
	chk.Uint32(result, math.MaxUint32)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"anotherArg",
	}

	result, args, err := szargs.NextUint32("TestArg", args)

	chk.NoErr(err)
	chk.Uint32(result, 309)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test uint16 positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextUint16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextUint16("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint16(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint16_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextUint16("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Uint16(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint16_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"65536", // MaxUint16 + 1 is out of range.
	}

	result, args, err := szargs.NextUint16("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrRange,
			"TestArg",
			"'65536'",
		),
	)
	chk.Uint16(result, math.MaxUint16)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"anotherArg",
	}

	result, args, err := szargs.NextUint16("TestArg", args)

	chk.NoErr(err)
	chk.Uint16(result, 309)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test uint8 positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextUint8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextUint8("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint8(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint8_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextUint8("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Uint8(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint8_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"256", // MaxUint8 + 1 is out of range.
	}

	result, args, err := szargs.NextUint8("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrRange,
			"TestArg",
			"'256'",
		),
	)
	chk.Uint8(result, math.MaxUint8)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"109",
		"anotherArg",
	}

	result, args, err := szargs.NextUint8("TestArg", args)

	chk.NoErr(err)
	chk.Uint8(result, 109)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test uint positional value.
 *
 ***************************************************************************
 */

func TestSzargsPositional_NextUint_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.NextUint("TestArg", nil)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"notANumber",
	}

	result, args, err := szargs.NextUint("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Uint(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"18446744073709551616", // MaxUint + 1 is out of range.
	}

	result, args, err := szargs.NextUint("TestArg", args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrRange,
			"TestArg",
			"'18446744073709551616'",
		),
	)
	chk.Uint(result, math.MaxUint)
	chk.StrSlice(args, nil)
}

func TestSzargsPositional_NextUint_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"anotherArg",
	}

	result, args, err := szargs.NextUint("TestArg", args)

	chk.NoErr(err)
	chk.Uint(result, 309)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}
