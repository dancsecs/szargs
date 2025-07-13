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

func TestSzargs_LastString_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastString("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(szargs.ErrMissing, "TestArg"),
	)
	chk.Str(result, "")
}

func TestSzargs_LastString_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.LastString("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Str(result, "309")
}

func TestSzargs_LastString_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
	})

	result := args.LastString("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Str(result, "309")
}

/*
 ***************************************************************************
 *
 *  Test args.float64 positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastFloat64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastFloat64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Float64(result, 0, 0) // No tolerance.
}

func TestSzargs_LastFloat64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastFloat64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			"TestArg",
			"'notANumber'",
		),
	)
	chk.Float64(result, 0, 0) // No tolerance.
}

func TestSzargs_LastFloat64_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"1.7e+309", // MaxFloat64 * 10 is out of range.
	})

	result := args.LastFloat64("TestArg", "the arg being tested")

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
}

func TestSzargs_LastFloat64_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-1.7e+309", // MinFloat64 * 10 is out of range.
	})

	result := args.LastFloat64("TestArg", "the arg being tested")

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
}

func TestSzargs_LastFloat64_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309.2",
		"anotherArg",
	})

	result := args.LastFloat64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Float64(result, 0, 0) // No tolerance.
}

func TestSzargs_LastFloat64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309.2",
	})

	result := args.LastFloat64("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Float64(result, 309.2, 0) // No tolerance.
}

/*
 ***************************************************************************
 *
 *  Test float32 positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastFloat32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastFloat32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Float32(result, 0, 0) // No tolerance.
}

func TestSzargs_LastFloat32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastFloat32("TestArg", "the arg being tested")

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
}

func TestSzargs_LastFloat32_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"1.7e+309",
	})

	result := args.LastFloat32("TestArg", "the arg being tested")

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
}

func TestSzargs_LastFloat32_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-1.7e+309",
	})

	result := args.LastFloat32("TestArg", "the arg being tested")

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
}

func TestSzargs_LastFloat32_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309.2",
		"anotherArg",
	})

	result := args.LastFloat32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Float32(result, 0, 0) // No tolerance.
}

func TestSzargs_LastFloat32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309.2",
	})

	result := args.LastFloat32("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Float32(result, 309.2, 0) // No tolerance.
}

/*
 ***************************************************************************
 *
 *  Test int64 positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastInt64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastInt64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int64(result, 0)
}

func TestSzargs_LastInt64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastInt64("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt64_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"9223372036854775808", // MaxInt64 + 1 is out of range.
	})

	result := args.LastInt64("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt64_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-9223372036854775809", // MinInt64 - 1 is out of range.
	})

	result := args.LastInt64("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt64_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.LastInt64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Int64(result, 0)
}

func TestSzargs_LastInt64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
	})

	result := args.LastInt64("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Int64(result, 309)
}

/*
 ***************************************************************************
 *
 *  Test int32 positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastInt32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastInt32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int32(result, 0)
}

func TestSzargs_LastInt32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastInt32("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt32_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"2147483648", // MaxInt32 + 1 is out of range.
	})

	result := args.LastInt32("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt32_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-2147483649", // MinInt32 - 1 is out of range.
	})

	result := args.LastInt32("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt32_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.LastInt32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Int32(result, 0)
}

func TestSzargs_LastInt32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
	})

	result := args.LastInt32("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Int32(result, 309)
}

/*
 ***************************************************************************
 *
 *  Test int16 positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastInt16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastInt16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int16(result, 0)
}

func TestSzargs_LastInt16_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastInt16("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt16_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"32768", // MaxInt16 + 1 is out of range.
	})

	result := args.LastInt16("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt16_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-32769", // MainInt16 - 1 is out of range.
	})

	result := args.LastInt16("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt16_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.LastInt16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Int16(result, 0)
}

func TestSzargs_LastInt16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
	})

	result := args.LastInt16("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Int16(result, 309)
}

/*
 ***************************************************************************
 *
 *  Test int8 positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastInt8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastInt8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int8(result, 0)
}

func TestSzargs_LastInt8_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastInt8("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt8_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"128", // MaxInt8 + 1 is out of range.
	})

	result := args.LastInt8("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt8_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-129", // MinInt8 - 1 is out of range.
	})

	result := args.LastInt8("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt8_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"109",
		"anotherArg",
	})

	result := args.LastInt8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Int8(result, 0)
}

func TestSzargs_LastInt8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"109",
	})

	result := args.LastInt8("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Int8(result, 109)
}

/*
 ***************************************************************************
 *
 *  Test int positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastInt_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastInt("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Int(result, 0)
}

func TestSzargs_LastInt_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastInt("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"9223372036854775808", // MaxInt + 1 is out of range.
	})

	result := args.LastInt("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-9223372036854775809", // MinInt - 1 is out of range.
	})

	result := args.LastInt("TestArg", "the arg being tested")

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
}

func TestSzargs_LastInt_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.LastInt("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Int(result, 0)
}

func TestSzargs_LastInt_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
	})

	result := args.LastInt("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Int(result, 309)
}

/*
 ***************************************************************************
 *
 *  Test uint64 positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastUint64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastUint64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint64(result, 0)
}

func TestSzargs_LastUint64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastUint64("TestArg", "the arg being tested")

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
}

func TestSzargs_LastUint64_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"18446744073709551616", // MaxUint64 + 1 is out of range.
	})

	result := args.LastUint64("TestArg", "the arg being tested")

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
}

func TestSzargs_LastUint64_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.LastUint64("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Uint64(result, 0)
}

func TestSzargs_LastUint64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
	})

	result := args.LastUint64("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Uint64(result, 309)
}

/*
 ***************************************************************************
 *
 *  Test uint32 positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastUint32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastUint32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint32(result, 0)
}

func TestSzargs_LastUint32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastUint32("TestArg", "the arg being tested")

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
}

func TestSzargs_LastUint32_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"4294967296", // MaxUint32 + 1 is out of range.
	})

	result := args.LastUint32("TestArg", "the arg being tested")

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
}

func TestSzargs_LastUint32_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.LastUint32("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Uint32(result, 0)
}

func TestSzargs_LastUint32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
	})

	result := args.LastUint32("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Uint32(result, 309)
}

/*
 ***************************************************************************
 *
 *  Test uint16 positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastUint16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastUint16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint16(result, 0)
}

func TestSzargs_LastUint16_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastUint16("TestArg", "the arg being tested")

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
}

func TestSzargs_LastUint16_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"65536", // MaxUint16 + 1 is out of range.
	})

	result := args.LastUint16("TestArg", "the arg being tested")

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
}

func TestSzargs_LastUint16_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.LastUint16("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Uint16(result, 0)
}

func TestSzargs_LastUint16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
	})

	result := args.LastUint16("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Uint16(result, 309)
}

/*
 ***************************************************************************
 *
 *  Test uint8 positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastUint8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastUint8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint8(result, 0)
}

func TestSzargs_LastUint8_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastUint8("TestArg", "the arg being tested")

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
}

func TestSzargs_LastUint8_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"256", // MaxUint8 + 1 is out of range.
	})

	result := args.LastUint8("TestArg", "the arg being tested")

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
}

func TestSzargs_LastUint8_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"109",
		"anotherArg",
	})

	result := args.LastUint8("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Uint8(result, 0)
}

func TestSzargs_LastUint8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"109",
	})

	result := args.LastUint8("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Uint8(result, 109)
}

/*
 ***************************************************************************
 *
 *  Test uint positional value.
 *
 ***************************************************************************
 */

func TestSzargs_LastUint_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result := args.LastUint("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrMissing,
			"TestArg",
		),
	)
	chk.Uint(result, 0)
}

func TestSzargs_LastUint_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"notANumber",
	})

	result := args.LastUint("TestArg", "the arg being tested")

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
}

func TestSzargs_LastUint_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"18446744073709551616", // MaxUint + 1 is out of range.
	})

	result := args.LastUint("TestArg", "the arg being tested")

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
}

func TestSzargs_LastUint_NotLast(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"anotherArg",
	})

	result := args.LastUint("TestArg", "the arg being tested")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrUnexpected,
			"[anotherArg]",
		),
	)
	chk.Uint(result, 0)
}

func TestSzargs_LastUint_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
	})

	result := args.LastUint("TestArg", "the arg being tested")

	chk.NoErr(args.Err())
	chk.Uint(result, 309)
}
