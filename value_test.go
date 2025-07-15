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
 *  Test string argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueString_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueString("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Str(result, "")
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueString_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"309",
		"-t",
		"testValue",
		"anotherArg",
	})

	result, found := args.ValueString("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
	chk.Str(result, "testValue")
	chk.StrSlice(
		args.Args(),
		[]string{
			"309",
			"anotherArg",
		},
	)
}

/*
 ***************************************************************************
 *
 *  Test args.float64 argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueFloat64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueFloat64("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Float64(result, 0, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueFloat64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueFloat64("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Float64(result, 0, 0)      // No tolerance.
	chk.StrSlice(args.Args(), nil) // Argument extracted.
}

func TestSzargs_ValueFloat64_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"1.7e+309", // MaxFloat64 * 10 is out of range.
	})

	result, found := args.ValueFloat64("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrRange,
			"-t",
			"'1.7e+309'",
		),
	)
	chk.False(found)
	chk.Float64(result, math.Inf(1), 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)      // Argument extracted.
}

func TestSzargs_ValueFloat64_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"-1.7e+309", // MinFloat64 * 10 is out of range.
	})

	result, found := args.ValueFloat64("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrRange,
			"-t",
			"'-1.7e+309'",
		),
	)
	chk.False(found)
	chk.Float64(result, math.Inf(-1), 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)       // Argument extracted.
}

func TestSzargs_ValueFloat64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"309.2",
		"anotherArg",
	})

	result, found := args.ValueFloat64("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.Float64(result, 309.2, 0) // No tolerance.
	chk.True(found)
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
 *  Test float32 argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueFloat32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueFloat32("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Float32(result, 0, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueFloat32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueFloat32("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Float32(result, 0, 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueFloat32_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"1.7e+309",
	})

	result, found := args.ValueFloat32("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrRange,
			"-t",
			"'1.7e+309'",
		),
	)
	chk.False(found)
	chk.Float32(result, float32(math.Inf(1)), 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueFloat32_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"-1.7e+309",
	})

	result, found := args.ValueFloat32("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrRange,
			"-t",
			"'-1.7e+309'",
		),
	)
	chk.False(found)
	chk.Float32(result, float32(math.Inf(-1)), 0) // No tolerance.
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueFloat32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"309.2",
		"anotherArg",
	})

	result, found := args.ValueFloat32("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True((found))
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
 *  Test int64 argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueInt64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueInt64("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Int64(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueInt64("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Int64(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt64_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"9223372036854775808", // MaxInt64 + 1 is out of range.
	})

	result, found := args.ValueInt64("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrRange,
			"-t",
			"'9223372036854775808'",
		),
	)
	chk.False(found)
	chk.Int64(result, math.MaxInt64)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt64_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"-9223372036854775809", // MinInt64 - 1 is out of range.
	})

	result, found := args.ValueInt64("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrRange,
			"-t",
			"'-9223372036854775809'",
		),
	)
	chk.False(found)
	chk.Int64(result, math.MinInt64)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"309",
		"anotherArg",
	})

	result, found := args.ValueInt64("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
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
 *  Test int32 argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueInt32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueInt32("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Int32(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueInt32("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Int32(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt32_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"2147483648", // MaxInt32 + 1 is out of range.
	})

	result, found := args.ValueInt32("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrRange,
			"-t",
			"'2147483648'",
		),
	)
	chk.False(found)
	chk.Int32(result, math.MaxInt32)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt32_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"-2147483649", // MinInt32 - 1 is out of range.
	})

	result, found := args.ValueInt32("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrRange,
			"-t",
			"'-2147483649'",
		),
	)
	chk.False(found)
	chk.Int32(result, math.MinInt32)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"309",
		"anotherArg",
	})

	result, found := args.ValueInt32("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
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
 *  Test int16 argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueInt16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueInt16("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Int16(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt16_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueInt16("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Int16(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt16_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"32768", // MaxInt16 + 1 is out of range.
	})

	result, found := args.ValueInt16("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrRange,
			"-t",
			"'32768'",
		),
	)
	chk.False(found)
	chk.Int16(result, math.MaxInt16)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt16_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"-32769", // MainInt16 - 1 is out of range.
	})

	result, found := args.ValueInt16("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrRange,
			"-t",
			"'-32769'",
		),
	)
	chk.False(found)
	chk.Int16(result, math.MinInt16)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"309",
		"anotherArg",
	})

	result, found := args.ValueInt16("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
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
 *  Test int8 argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueInt8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueInt8("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Int8(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt8_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueInt8("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Int8(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt8_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"128", // MaxInt8 + 1 is out of range.
	})

	result, found := args.ValueInt8("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrRange,
			"-t",
			"'128'",
		),
	)
	chk.False(found)
	chk.Int8(result, math.MaxInt8)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt8_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"-129", // MinInt8 - 1 is out of range.
	})

	result, found := args.ValueInt8("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrRange,
			"-t",
			"'-129'",
		),
	)
	chk.False(found)
	chk.Int8(result, math.MinInt8)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"109",
		"anotherArg",
	})

	result, found := args.ValueInt8("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
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
 *  Test int argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueInt_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueInt("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Int(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueInt("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Int(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"9223372036854775808", // MaxInt + 1 is out of range.
	})

	result, found := args.ValueInt("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrRange,
			"-t",
			"'9223372036854775808'",
		),
	)
	chk.False(found)
	chk.Int(result, math.MaxInt)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"-9223372036854775809", // MinInt - 1 is out of range.
	})

	result, found := args.ValueInt("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrRange,
			"-t",
			"'-9223372036854775809'",
		),
	)
	chk.False(found)
	chk.Int(result, math.MinInt)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueInt_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"309",
		"anotherArg",
	})

	result, found := args.ValueInt("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
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
 *  Test uint64 argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueUint64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueUint64("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Uint64(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueUint64("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Uint64(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint64_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"18446744073709551616", // MaxUint64 + 1 is out of range.
	})

	result, found := args.ValueUint64("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrRange,
			"-t",
			"'18446744073709551616'",
		),
	)
	chk.False(found)
	chk.Uint64(result, math.MaxUint64)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"309",
		"anotherArg",
	})

	result, found := args.ValueUint64("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
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
 *  Test uint32 argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueUint32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueUint32("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Uint32(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueUint32("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Uint32(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint32_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"4294967296", // MaxUint32 + 1 is out of range.
	})
	result, found := args.ValueUint32("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrRange,
			"-t",
			"'4294967296'",
		),
	)
	chk.False(found)
	chk.Uint32(result, math.MaxUint32)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"309",
		"anotherArg",
	})

	result, found := args.ValueUint32("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
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
 *  Test uint16 argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueUint16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueUint16("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Uint16(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint16_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueUint16("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Uint16(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint16_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"65536", // MaxUint16 + 1 is out of range.
	})

	result, found := args.ValueUint16("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrRange,
			"-t",
			"'65536'",
		),
	)
	chk.False(found)
	chk.Uint16(result, math.MaxUint16)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"309",
		"anotherArg",
	})

	result, found := args.ValueUint16("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
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
 *  Test uint8 argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueUint8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueUint8("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Uint8(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint8_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueUint8("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Uint8(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint8_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"256", // MaxUint8 + 1 is out of range.
	})

	result, found := args.ValueUint8("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrRange,
			"-t",
			"'256'",
		),
	)
	chk.False(found)
	chk.Uint8(result, math.MaxUint8)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"109",
		"anotherArg",
	})

	result, found := args.ValueUint8("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
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
 *  Test uint argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueUint_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueUint("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Uint(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueUint("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Uint(result, 0)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"18446744073709551616", // MaxUint + 1 is out of range.
	})

	result, found := args.ValueUint("-t", "the test flag")

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrRange,
			"-t",
			"'18446744073709551616'",
		),
	)
	chk.False(found)
	chk.Uint(result, math.MaxUint)
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueUint_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"309",
		"anotherArg",
	})

	result, found := args.ValueUint("-t", "the test flag")

	chk.NoErr(args.Err())
	chk.True(found)
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
 *  Test uint argument value.
 *
 ***************************************************************************
 */

func TestSzargs_ValueOption_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
	})

	result, found := args.ValueOption(
		"-t", []string{"a", "b"}, "the test flag",
	)

	chk.NoErr(args.Err())
	chk.False(found)
	chk.Str(result, "")
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueOption_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"notANumber",
	})

	result, found := args.ValueOption(
		"-t", []string{"a", "b"}, "the test flag",
	)

	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrInvalidOption,
			"'notANumber' "+
				"(-t must be one of [a b])",
		),
	)
	chk.False(found)
	chk.Str(result, "")
	chk.StrSlice(args.Args(), nil)
}

func TestSzargs_ValueOption_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := szargs.New("program description", []string{
		"programName",
		"-t",
		"b",
		"anotherArg",
	})

	result, found := args.ValueOption(
		"-t", []string{"a", "b"}, "the test flag",
	)

	chk.NoErr(args.Err())
	chk.True(found)
	chk.Str(result, "b")
	chk.StrSlice(
		args.Args(),
		[]string{
			"anotherArg",
		},
	)
}
