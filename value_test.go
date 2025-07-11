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

func TestSzargsArgument_ValueNonePresent(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := szargs.Arg("-n").Value(nil)

	chk.Str(value, "")
	chk.False(found)
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	value, found, args, err = szargs.Arg("-n").Value(
		[]string{"arg1", "arg2"},
	)

	chk.Str(value, "")
	chk.False(found)
	chk.StrSlice(args,
		[]string{"arg1", "arg2"},
	)
	chk.NoErr(err)
}

func TestSzargsArgument_ValueBeginning(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := szargs.Arg("-n").Value(
		[]string{"-n", "theName"},
	)

	chk.Str(value, "theName")
	chk.True(found)
	chk.StrSlice(args, []string{})
	chk.NoErr(err)

	value, found, args, err = szargs.Arg("-n").Value(
		[]string{"-n", "theName", "arg1", "arg2"},
	)

	chk.Str(value, "theName")
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargsArgument_ValueMiddle(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := szargs.Arg("-n").Value(
		[]string{"arg1", "-n", "theName", "arg2"},
	)

	chk.Str(value, "theName")
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargsArgument_ValueEnd(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := szargs.Arg("-n").Value(
		[]string{"arg1", "arg2", "-n", "theName"},
	)

	chk.Str(value, "theName")
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargsArgument_ValueDuplicate(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := szargs.Arg("-n").Value(
		[]string{"-n", "firstName", "arg1", "arg2", "-n", "secondName"},
	)

	chk.Str(value, "")
	chk.False(found)
	chk.StrSlice(
		args,
		[]string{"arg1", "arg2"},
	)
	chk.Err(
		err,
		szargs.ErrAmbiguous.Error()+
			": '-n secondName' already set to: 'firstName'",
	)
}

func TestSzargsArgument_ValueTriplicate(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := szargs.Arg("-n").Value(
		[]string{
			"-n", "firstName",
			"arg1",
			"-n", "secondName",
			"arg2",
			"-n", "thirdName",
		},
	)

	chk.Str(value, "")
	chk.False(found)
	chk.StrSlice(
		args,
		[]string{"arg1", "arg2"},
	)
	chk.Err(
		err,
		szargs.ErrAmbiguous.Error()+
			": '-n secondName' already set to: 'firstName'"+
			": "+
			szargs.ErrAmbiguous.Error()+
			": '-n thirdName' already set to: 'firstName'"+
			"",
	)
}

func TestSzargsArgument_ValueMissing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := szargs.Arg("-n").Value(
		[]string{"arg1", "arg2", "-n"},
	)

	chk.Str(value, "")
	chk.False(found)
	chk.StrSlice(
		args,
		[]string{"arg1", "arg2"},
	)
	chk.Err(
		err,
		szargs.ErrMissing.Error()+
			": '-n value'",
	)
}

/*
 ***************************************************************************
 *
 *  Test string argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueString_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueString(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Str(result, "")
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueString_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"-t",
		"testValue",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueString(args)

	chk.NoErr(err)
	chk.True(found)
	chk.Str(result, "testValue")
	chk.StrSlice(
		args,
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

func TestSzargsArgument_ValueFloat64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueFloat64(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Float64(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueFloat64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueFloat64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Float64(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)   // Argument extracted.
}

func TestSzargsArgument_ValueFloat64_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"1.7e+309", // MaxFloat64 * 10 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueFloat64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrRange,
			"-t",
			"'1.7e+309'",
		),
	)
	chk.False(found)
	chk.Float64(result, math.Inf(1), 0) // No tolerance.
	chk.StrSlice(args, nil)             // Argument extracted.
}

func TestSzargsArgument_ValueFloat64_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"-1.7e+309", // MinFloat64 * 10 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueFloat64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrRange,
			"-t",
			"'-1.7e+309'",
		),
	)
	chk.False(found)
	chk.Float64(result, math.Inf(-1), 0) // No tolerance.
	chk.StrSlice(args, nil)              // Argument extracted.
}

func TestSzargsArgument_ValueFloat64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309.2",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueFloat64(args)

	chk.NoErr(err)
	chk.Float64(result, 309.2, 0) // No tolerance.
	chk.True(found)
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
 *  Test float32 argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueFloat32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueFloat32(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Float32(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueFloat32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueFloat32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Float32(result, 0, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueFloat32_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"1.7e+309",
	}

	result, found, args, err := szargs.Arg("-t").ValueFloat32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrRange,
			"-t",
			"'1.7e+309'",
		),
	)
	chk.False(found)
	chk.Float32(result, float32(math.Inf(1)), 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueFloat32_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"-1.7e+309",
	}

	result, found, args, err := szargs.Arg("-t").ValueFloat32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrRange,
			"-t",
			"'-1.7e+309'",
		),
	)
	chk.False(found)
	chk.Float32(result, float32(math.Inf(-1)), 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueFloat32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309.2",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueFloat32(args)

	chk.NoErr(err)
	chk.True((found))
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
 *  Test int64 argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueInt64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueInt64(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Int64(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueInt64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Int64(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt64_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"9223372036854775808", // MaxInt64 + 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueInt64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrRange,
			"-t",
			"'9223372036854775808'",
		),
	)
	chk.False(found)
	chk.Int64(result, math.MaxInt64)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt64_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"-9223372036854775809", // MinInt64 - 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueInt64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrRange,
			"-t",
			"'-9223372036854775809'",
		),
	)
	chk.False(found)
	chk.Int64(result, math.MinInt64)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueInt64(args)

	chk.NoErr(err)
	chk.True(found)
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
 *  Test int32 argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueInt32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueInt32(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Int32(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueInt32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Int32(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt32_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"2147483648", // MaxInt32 + 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueInt32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrRange,
			"-t",
			"'2147483648'",
		),
	)
	chk.False(found)
	chk.Int32(result, math.MaxInt32)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt32_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"-2147483649", // MinInt32 - 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueInt32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrRange,
			"-t",
			"'-2147483649'",
		),
	)
	chk.False(found)
	chk.Int32(result, math.MinInt32)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueInt32(args)

	chk.NoErr(err)
	chk.True(found)
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
 *  Test int16 argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueInt16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueInt16(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Int16(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt16_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueInt16(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Int16(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt16_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"32768", // MaxInt16 + 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueInt16(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrRange,
			"-t",
			"'32768'",
		),
	)
	chk.False(found)
	chk.Int16(result, math.MaxInt16)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt16_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"-32769", // MainInt16 - 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueInt16(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrRange,
			"-t",
			"'-32769'",
		),
	)
	chk.False(found)
	chk.Int16(result, math.MinInt16)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueInt16(args)

	chk.NoErr(err)
	chk.True(found)
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
 *  Test int8 argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueInt8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueInt8(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Int8(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt8_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueInt8(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Int8(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt8_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"128", // MaxInt8 + 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueInt8(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrRange,
			"-t",
			"'128'",
		),
	)
	chk.False(found)
	chk.Int8(result, math.MaxInt8)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt8_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"-129", // MinInt8 - 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueInt8(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrRange,
			"-t",
			"'-129'",
		),
	)
	chk.False(found)
	chk.Int8(result, math.MinInt8)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"109",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueInt8(args)

	chk.NoErr(err)
	chk.True(found)
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
 *  Test int argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueInt_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueInt(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Int(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueInt(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Int(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt_InvalidRangeHigh(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"9223372036854775808", // MaxInt + 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueInt(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrRange,
			"-t",
			"'9223372036854775808'",
		),
	)
	chk.False(found)
	chk.Int(result, math.MaxInt)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt_InvalidRangeLow(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"-9223372036854775809", // MinInt - 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueInt(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrRange,
			"-t",
			"'-9223372036854775809'",
		),
	)
	chk.False(found)
	chk.Int(result, math.MinInt)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueInt(args)

	chk.NoErr(err)
	chk.True(found)
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
 *  Test uint64 argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueUint64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueUint64(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Uint64(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint64_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueUint64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Uint64(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint64_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"18446744073709551616", // MaxUint64 + 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueUint64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrRange,
			"-t",
			"'18446744073709551616'",
		),
	)
	chk.False(found)
	chk.Uint64(result, math.MaxUint64)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueUint64(args)

	chk.NoErr(err)
	chk.True(found)
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
 *  Test uint32 argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueUint32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueUint32(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Uint32(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint32_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueUint32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Uint32(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint32_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"4294967296", // MaxUint32 + 1 is out of range.
	}
	result, found, args, err := szargs.Arg("-t").ValueUint32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrRange,
			"-t",
			"'4294967296'",
		),
	)
	chk.False(found)
	chk.Uint32(result, math.MaxUint32)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueUint32(args)

	chk.NoErr(err)
	chk.True(found)
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
 *  Test uint16 argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueUint16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueUint16(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Uint16(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint16_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueUint16(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Uint16(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint16_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"65536", // MaxUint16 + 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueUint16(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrRange,
			"-t",
			"'65536'",
		),
	)
	chk.False(found)
	chk.Uint16(result, math.MaxUint16)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueUint16(args)

	chk.NoErr(err)
	chk.True(found)
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
 *  Test uint8 argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueUint8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueUint8(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Uint8(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint8_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueUint8(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Uint8(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint8_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"256", // MaxUint8 + 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueUint8(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrRange,
			"-t",
			"'256'",
		),
	)
	chk.False(found)
	chk.Uint8(result, math.MaxUint8)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"109",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueUint8(args)

	chk.NoErr(err)
	chk.True(found)
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
 *  Test uint argument value.
 *
 ***************************************************************************
 */

func TestSzargsArgument_ValueUint_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, found, args, err := szargs.Arg("-t").ValueUint(nil)

	chk.NoErr(err)
	chk.False(found)
	chk.Uint(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint_InvalidSyntax(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
	}

	result, found, args, err := szargs.Arg("-t").ValueUint(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
		),
	)
	chk.False(found)
	chk.Uint(result, 0)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint_InvalidRange(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"18446744073709551616", // MaxUint + 1 is out of range.
	}

	result, found, args, err := szargs.Arg("-t").ValueUint(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrRange,
			"-t",
			"'18446744073709551616'",
		),
	)
	chk.False(found)
	chk.Uint(result, math.MaxUint)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueUint_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
	}

	result, found, args, err := szargs.Arg("-t").ValueUint(args)

	chk.NoErr(err)
	chk.True(found)
	chk.Uint(result, 309)
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}
