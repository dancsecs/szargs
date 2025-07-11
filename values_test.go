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

func TestSzargsArgument_ValuesNonePresent(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Arg("-n").Values(nil)

	chk.StrSlice(value, nil)
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	value, args, err = szargs.Arg("-n").Values(
		[]string{"arg1", "arg2"},
	)

	chk.StrSlice(value, nil)
	chk.StrSlice(args,
		[]string{"arg1", "arg2"},
	)
	chk.NoErr(err)
}

func TestSzargsArgument_ValuesAtBeginning(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Arg("-n").Values(
		[]string{"-n", "theName"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{})
	chk.NoErr(err)

	value, args, err = szargs.Arg("-n").Values(
		[]string{"-n", "theName", "arg1", "arg2"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargsArgument_ValuesMiddle(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Arg("-n").Values(
		[]string{"arg1", "-n", "theName", "arg2"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargsArgument_ValuesEnd(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Arg("-n").Values(
		[]string{"arg1", "arg2", "-n", "theName"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargsArgument_ValuesDuplicate(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Arg("-n").Values(
		[]string{"-n", "firstName", "arg1", "arg2", "-n", "secondName"},
	)

	chk.StrSlice(value, []string{"firstName", "secondName"})
	chk.StrSlice(
		args,
		[]string{"arg1", "arg2"},
	)
	chk.NoErr(err)
}

func TestSzargsArgument_ValuesMissing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Arg("-n").Values(
		[]string{"arg1", "arg2", "-n"},
	)

	chk.StrSlice(value, nil)
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

func TestSzargsArgument_ValuesString_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesString(nil)

	chk.NoErr(err)
	chk.StrSlice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesString_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"309",
		"-t",
		"testValue1",
		"anotherArg1",
		"-t",
		"testValue2",
		"anotherArg2",
	}

	result, args, err := szargs.Arg("-t").ValuesString(args)

	chk.NoErr(err)
	chk.StrSlice(
		result,
		[]string{
			"testValue1",
			"testValue2",
		},
	)
	chk.StrSlice(
		args,
		[]string{
			"309",
			"anotherArg1",
			"anotherArg2",
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

func TestSzargsArgument_ValuesFloat64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesFloat64(nil)

	chk.NoErr(err)
	chk.Float64Slice(result, nil, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesFloat64_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber1",
		"-t",
		"1.7e+309", // MaxFloat64 * 10 is out of range.
		"-t",
		"-1.7e+309", // MinFloat64 * 10 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesFloat64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			"-t",
			"'notANumber1'",
			szargs.ErrInvalidFloat64,
			szargs.ErrRange,
			"-t",
			"'1.7e+309'",
			szargs.ErrInvalidFloat64,
			szargs.ErrRange,
			"-t",
			"'-1.7e+309'",
		),
	)
	chk.Float64Slice(result, nil, 0) // No tolerance.
	chk.StrSlice(args, nil)          // Argument extracted.
}

func TestSzargsArgument_ValuesFloat64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309.2",
		"-t",
		"127.49",
		"anotherArg",
	}

	result, args, err := szargs.Arg("-t").ValuesFloat64(args)

	chk.NoErr(err)
	chk.Float64Slice(result, []float64{309.2, 127.49}, 0) // No tolerance.
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

func TestSzargsArgument_ValuesFloat32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesFloat32(nil)

	chk.NoErr(err)
	chk.Float32Slice(result, nil, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesFloat32_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"1.7e+309",
		"-t",
		"-1.7e+309",
	}

	result, args, err := szargs.Arg("-t").ValuesFloat32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidFloat32,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidFloat32,
			szargs.ErrRange,
			"-t",
			"'1.7e+309'",
			szargs.ErrInvalidFloat32,
			szargs.ErrRange,
			"-t",
			"'-1.7e+309'",
		),
	)
	chk.Float32Slice(result, nil, 0) // No tolerance.
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesFloat32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"222.5",
		"-t",
		"309.2",
		"anotherArg",
	}

	result, args, err := szargs.Arg("-t").ValuesFloat32(args)

	chk.NoErr(err)
	chk.Float32Slice(result, []float32{222.5, 309.2}, 0) // No tolerance.
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

func TestSzargsArgument_ValuesInt64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesInt64(nil)

	chk.NoErr(err)
	chk.Int64Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesInt64_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"9223372036854775808", // MaxInt64 + 1 is out of range.
		"-t",
		"-9223372036854775809", // MinInt64 - 1 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesInt64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt64,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidInt64,
			szargs.ErrRange,
			"-t",
			"'9223372036854775808'",
			szargs.ErrInvalidInt64,
			szargs.ErrRange,
			"-t",
			"'-9223372036854775809'",
		),
	)
	chk.Int64Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesInt64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"908",
		"-t",
		"309",
		"anotherArg",
	}

	result, args, err := szargs.Arg("-t").ValuesInt64(args)

	chk.NoErr(err)
	chk.Int64Slice(result, []int64{908, 309})
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

func TestSzargsArgument_ValuesInt32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesInt32(nil)

	chk.NoErr(err)
	chk.Int32Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValueInt32_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"2147483648", // MaxInt32 + 1 is out of range.
		"-t",
		"-2147483649", // MinInt32 - 1 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesInt32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt32,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidInt32,
			szargs.ErrRange,
			"-t",
			"'2147483648'",
			szargs.ErrInvalidInt32,
			szargs.ErrRange,
			"-t",
			"'-2147483649'",
		),
	)
	chk.Int32Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesInt32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"4493",
		"-t",
		"309",
		"anotherArg",
	}

	result, args, err := szargs.Arg("-t").ValuesInt32(args)

	chk.NoErr(err)
	chk.Int32Slice(result, []int32{4493, 309})
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

func TestSzargsArgument_ValuesInt16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesInt16(nil)

	chk.NoErr(err)
	chk.Int16Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesInt16_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"32768", // MaxInt16 + 1 is out of range.
		"-t",
		"-32769", // MainInt16 - 1 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesInt16(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt16,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidInt16,
			szargs.ErrRange,
			"-t",
			"'32768'",
			szargs.ErrInvalidInt16,
			szargs.ErrRange,
			"-t",
			"'-32769'",
		),
	)
	chk.Int16Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesInt16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"-t",
		"111",
		"anotherArg",
	}

	result, args, err := szargs.Arg("-t").ValuesInt16(args)

	chk.NoErr(err)
	chk.Int16Slice(result, []int16{309, 111})
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

func TestSzargsArgument_ValuesInt8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesInt8(nil)

	chk.NoErr(err)
	chk.Int8Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesInt8_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"128", // MaxInt8 + 1 is out of range.
		"-t",
		"-129", // MinInt8 - 1 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesInt8(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt8,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidInt8,
			szargs.ErrRange,
			"-t",
			"'128'",
			szargs.ErrInvalidInt8,
			szargs.ErrRange,
			"-t",
			"'-129'",
		),
	)
	chk.Int8Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesInt8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"109",
		"-t",
		"-109",
		"anotherArg",
	}

	result, args, err := szargs.Arg("-t").ValuesInt8(args)

	chk.NoErr(err)
	chk.Int8Slice(result, []int8{109, -109})
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

func TestSzargsArgument_ValuesInt_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesInt(nil)

	chk.NoErr(err)
	chk.IntSlice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesInt_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"9223372036854775808", // MaxInt + 1 is out of range.
		"-t",
		"-9223372036854775809", // MinInt - 1 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesInt(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidInt,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidInt,
			szargs.ErrRange,
			"-t",
			"'9223372036854775808'",
			szargs.ErrInvalidInt,
			szargs.ErrRange,
			"-t",
			"'-9223372036854775809'",
		),
	)
	chk.IntSlice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesInt_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"-1308",
		"-t",
		"309",
		"anotherArg",
	}

	result, args, err := szargs.Arg("-t").ValuesInt(args)

	chk.NoErr(err)
	chk.IntSlice(result, []int{-1308, 309})
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

func TestSzargsArgument_ValuesUint64_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesUint64(nil)

	chk.NoErr(err)
	chk.Uint64Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesUint64_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"18446744073709551616", // MaxUint64 + 1 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesUint64(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint64,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidUint64,
			szargs.ErrRange,
			"-t",
			"'18446744073709551616'",
		),
	)
	chk.Uint64Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesUint64_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"-t",
		"15309",
		"anotherArg",
	}

	result, args, err := szargs.Arg("-t").ValuesUint64(args)

	chk.NoErr(err)
	chk.Uint64Slice(result, []uint64{309, 15309})
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

func TestSzargsArgument_ValuesUint32_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesUint32(nil)

	chk.NoErr(err)
	chk.Uint32Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesUint32_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"4294967296", // MaxUint32 + 1 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesUint32(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint32,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidUint32,
			szargs.ErrRange,
			"-t",
			"'4294967296'",
		),
	)
	chk.Uint32Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesUint32_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
		"-t",
		"123",
	}

	result, args, err := szargs.Arg("-t").ValuesUint32(args)

	chk.NoErr(err)
	chk.Uint32Slice(result, []uint32{309, 123})
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

func TestSzargsArgument_ValuesUint16_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesUint16(nil)

	chk.NoErr(err)
	chk.Uint16Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesUint16_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"65536", // MaxUint16 + 1 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesUint16(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint16,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidUint16,
			szargs.ErrRange,
			"-t",
			"'65536'",
		),
	)
	chk.Uint16Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesUint16_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
		"-t",
		"867",
	}

	result, args, err := szargs.Arg("-t").ValuesUint16(args)

	chk.NoErr(err)
	chk.Uint16Slice(result, []uint16{309, 867})
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

func TestSzargsArgument_ValuesUint8_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesUint8(nil)

	chk.NoErr(err)
	chk.Uint8Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesUint8_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"256", // MaxUint8 + 1 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesUint8(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint8,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidUint8,
			szargs.ErrRange,
			"-t",
			"'256'",
		),
	)
	chk.Uint8Slice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesUint8_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"109",
		"-t",
		"111",
		"anotherArg",
	}

	result, args, err := szargs.Arg("-t").ValuesUint8(args)

	chk.NoErr(err)
	chk.Uint8Slice(result, []uint8{109, 111})
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

func TestSzargsArgument_ValuesUint_Missing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	result, args, err := szargs.Arg("-t").ValuesUint(nil)

	chk.NoErr(err)
	chk.UintSlice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesUint_Invalid(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"notANumber",
		"-t",
		"18446744073709551616", // MaxUint + 1 is out of range.
	}

	result, args, err := szargs.Arg("-t").ValuesUint(args)

	chk.Err(
		err,
		chk.ErrChain(
			szargs.ErrInvalidUint,
			szargs.ErrSyntax,
			"-t",
			"'notANumber'",
			szargs.ErrInvalidUint,
			szargs.ErrRange,
			"-t",
			"'18446744073709551616'",
		),
	)
	chk.UintSlice(result, nil)
	chk.StrSlice(args, nil)
}

func TestSzargsArgument_ValuesUint_Success(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := []string{
		"-t",
		"309",
		"anotherArg",
		"-t",
		"765",
	}

	result, args, err := szargs.Arg("-t").ValuesUint(args)

	chk.NoErr(err)
	chk.UintSlice(result, []uint{309, 765})
	chk.StrSlice(
		args,
		[]string{
			"anotherArg",
		},
	)
}
