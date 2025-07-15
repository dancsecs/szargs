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

/*
 ***************************************************************************
 *
 *  Test string positional value.
 *
 ***************************************************************************
 */

func TestSzargs_NextString_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextString("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(szargs.ErrMissing, "TestArg"),
	)
	chk.Str(result, "")
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextString_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.NextString("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Str(result, "309")
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextFloat64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextFloat64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Float64(result, 0, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextFloat64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextFloat64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Float64(result, 0, 0)      // No tolerance.
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_NextFloat64_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"1.7e+309", // MaxFloat64 * 10 is out of range.
	})

	result := args.NextFloat64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrRange,
			"TestArg",
			"'1.7e+309'",
		),
	)
	chk.Float64(result, math.Inf(1), 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)      // Argument extracted.
}

func TestSzargs_NextFloat64_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-1.7e+309", // MinFloat64 * 10 is out of range.
	})

	result := args.NextFloat64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrRange,
			"TestArg",
			"'-1.7e+309'",
		),
	)
	chk.Float64(result, math.Inf(-1), 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)       // Argument extracted.
}

func TestSzargs_NextFloat64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309.2",
		"anotherArg",
	})

	result := args.NextFloat64("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Float64(result, 309.2, 0) // No tolerance.
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextFloat32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextFloat32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Float32(result, 0, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextFloat32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextFloat32("TestArg", "the arg being tested")

	// Parsing syntax error.
	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Float32(result, 0, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextFloat32_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"1.7e+309",
	})

	result := args.NextFloat32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrRange,
			"TestArg",
			"'1.7e+309'",
		),
	)
	chk.Float32(result, float32(math.Inf(1)), 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextFloat32_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-1.7e+309",
	})

	result := args.NextFloat32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrRange,
			"TestArg",
			"'-1.7e+309'",
		),
	)
	chk.Float32(result, float32(math.Inf(-1)), 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextFloat32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309.2",
		"anotherArg",
	})

	result := args.NextFloat32("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Float32(result, 309.2, 0) // No tolerance.
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextInt64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextInt64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int64(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextInt64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)

	chk.Int64(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt64_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"9223372036854775808", // MaxInt64 + 1 is out of range.
	})

	result := args.NextInt64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrRange,
			"TestArg",
			"'9223372036854775808'",
		),
	)
	chk.Int64(result, math.MaxInt64)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt64_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-9223372036854775809", // MinInt64 - 1 is out of range.
	})

	result := args.NextInt64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrRange,
			"TestArg",
			"'-9223372036854775809'",
		),
	)
	chk.Int64(result, math.MinInt64)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.NextInt64("TestArg", "the arg being tested")

	// No Error
	chk.NoErr(args.Err())
	chk.Int64(result, 309)
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextInt32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextInt32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int32(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextInt32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Int32(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt32_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"2147483648", // MaxInt32 + 1 is out of range.
	})

	result := args.NextInt32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrRange,
			"TestArg",
			"'2147483648'",
		),
	)
	chk.Int32(result, math.MaxInt32)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt32_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-2147483649", // MinInt32 - 1 is out of range.
	})

	result := args.NextInt32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrRange,
			"TestArg",
			"'-2147483649'",
		),
	)
	chk.Int32(result, math.MinInt32)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.NextInt32("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Int32(result, 309)
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextInt16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextInt16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int16(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt16_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextInt16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Int16(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt16_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"32768", // MaxInt16 + 1 is out of range.
	})

	result := args.NextInt16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrRange,
			"TestArg",
			"'32768'",
		),
	)
	chk.Int16(result, math.MaxInt16)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt16_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-32769", // MainInt16 - 1 is out of range.
	})

	result := args.NextInt16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrRange,
			"TestArg",
			"'-32769'",
		),
	)
	chk.Int16(result, math.MinInt16)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.NextInt16("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Int16(result, 309)
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextInt8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextInt8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int8(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt8_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextInt8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Int8(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt8_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"128", // MaxInt8 + 1 is out of range.
	})

	result := args.NextInt8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrRange,
			"TestArg",
			"'128'",
		),
	)
	chk.Int8(result, math.MaxInt8)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt8_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-129", // MinInt8 - 1 is out of range.
	})

	result := args.NextInt8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrRange,
			"TestArg",
			"'-129'",
		),
	)
	chk.Int8(result, math.MinInt8)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"109",
		"anotherArg",
	})

	result := args.NextInt8("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Int8(result, 109)
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextInt_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextInt("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextInt("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Int(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"9223372036854775808", // MaxInt + 1 is out of range.
	})

	result := args.NextInt("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrRange,
			"TestArg",
			"'9223372036854775808'",
		),
	)
	chk.Int(result, math.MaxInt)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-9223372036854775809", // MinInt - 1 is out of range.
	})

	result := args.NextInt("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrRange,
			"TestArg",
			"'-9223372036854775809'",
		),
	)
	chk.Int(result, math.MinInt)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextInt_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.NextInt("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Int(result, 309)
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextUint64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextUint64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint64(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextUint64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Uint64(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint64_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"18446744073709551616", // MaxUint64 + 1 is out of range.
	})

	result := args.NextUint64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrRange,
			"TestArg",
			"'18446744073709551616'",
		),
	)
	chk.Uint64(result, math.MaxUint64)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.NextUint64("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Uint64(result, 309)
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextUint32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextUint32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint32(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextUint32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Uint32(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint32_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"4294967296", // MaxUint32 + 1 is out of range.
	})

	result := args.NextUint32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrRange,
			"TestArg",
			"'4294967296'",
		),
	)
	chk.Uint32(result, math.MaxUint32)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.NextUint32("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Uint32(result, 309)
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextUint16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextUint16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint16(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint16_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextUint16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Uint16(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint16_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"65536", // MaxUint16 + 1 is out of range.
	})

	result := args.NextUint16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrRange,
			"TestArg",
			"'65536'",
		),
	)
	chk.Uint16(result, math.MaxUint16)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.NextUint16("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Uint16(result, 309)
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextUint8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextUint8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint8(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint8_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextUint8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Uint8(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint8_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"256", // MaxUint8 + 1 is out of range.
	})

	result := args.NextUint8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrRange,
			"TestArg",
			"'256'",
		),
	)
	chk.Uint8(result, math.MaxUint8)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"109",
		"anotherArg",
	})

	result := args.NextUint8("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Uint8(result, 109)
	chk.StrSlice(
		args.Args(),
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

func TestSzargs_NextUint_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextUint("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextUint("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Uint(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"18446744073709551616", // MaxUint + 1 is out of range.
	})

	result := args.NextUint("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrRange,
			"TestArg",
			"'18446744073709551616'",
		),
	)
	chk.Uint(result, math.MaxUint)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextUint_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.NextUint("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Uint(result, 309)
	chk.StrSlice(
		args.Args(),
		[]string{
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test option positional value.
 *
 ***************************************************************************
 */

func TestSzargs_NextOption_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.NextOption(
		"TestArg", []string{"a", "b"}, "the arg being tested",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Str(result, "")
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextOption_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.NextOption(
		"TestArg", []string{"a", "b"}, "the arg being tested",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidOption,
			"'notANumber' "+
				"(TestArg must be one of [a b])",
		),
	)
	chk.Str(result, "")
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_NextOption_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"b",
		"anotherArg",
	})

	result := args.NextOption(
		"TestArg", []string{"a", "b"}, "the arg being tested",
	)

	chk.NoErr(args.Err())
	chk.Str(result, "b")
	chk.StrSlice(
		args.Args(),
		[]string{
			"anotherArg",
		},
	)
}
