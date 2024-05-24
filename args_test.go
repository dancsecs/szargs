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

func TestSzArgs_NeedNextArg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	var (
		arg  string
		args []string
		err  error
	)

	arg, args, err = szargs.NeedNextArg("TestArg", args)

	chk.Str(arg, "")
	chk.StrSlice(args, nil)
	chk.Err(
		err,
		"missing argument: TestArg",
	)

	args = []string{"arg1", "arg2"}

	arg, args, err = szargs.NeedNextArg("TestArg", args)

	chk.Str(arg, "arg1")
	chk.StrSlice(args, []string{"arg2"})
	chk.NoErr(err)

	arg, args, err = szargs.NeedNextArg("TestArg", args)

	chk.Str(arg, "arg2")
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	arg, args, err = szargs.NeedNextArg("TestArg", args)

	chk.Str(arg, "")
	chk.StrSlice(args, nil)
	chk.Err(
		err,
		szargs.ErrMissing.Error()+": TestArg",
	)
}

func TestSzArgs_ArgCount(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	var (
		arg   = szargs.Arg("-v")
		count int
		args  []string
	)

	count, args = arg.Count(args)
	chk.Int(count, 0)
	chk.StrSlice(args, nil)

	count, args = arg.Count([]string{"arg1", "arg2"})
	chk.Int(count, 0)
	chk.StrSlice(args, []string{"arg1", "arg2"})

	count, args = arg.Count([]string{"-v", "arg1", "arg2"})
	chk.Int(count, 1)
	chk.StrSlice(args, []string{"arg1", "arg2"})

	count, args = arg.Count([]string{"arg1", "-v", "arg2"})
	chk.Int(count, 1)
	chk.StrSlice(args, []string{"arg1", "arg2"})

	count, args = arg.Count([]string{"arg1", "arg2", "-v"})
	chk.Int(count, 1)
	chk.StrSlice(args, []string{"arg1", "arg2"})

	count, args = arg.Count([]string{"-v", "arg1", "-v", "arg2", "-v"})
	chk.Int(count, 3)
	chk.StrSlice(args, []string{"arg1", "arg2"})
}

func TestSzArgs_Is(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	var (
		err   error
		arg   = szargs.Arg("-v")
		found bool
		args  []string
	)

	found, args, err = arg.Is(nil)
	chk.False(found)
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	found, args, err = arg.Is([]string{"arg1"})
	chk.False(found)
	chk.StrSlice(args, []string{"arg1"})
	chk.NoErr(err)

	found, args, err = arg.Is([]string{"arg1", "arg2"})
	chk.False(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)

	found, args, err = arg.Is([]string{"-v", "arg1", "arg2"})
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)

	found, args, err = arg.Is([]string{"arg1", "-v", "arg2"})
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)

	found, args, err = arg.Is([]string{"arg1", "arg2", "-v"})
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)

	found, args, err = arg.Is([]string{"-v", "arg1", "arg2", "-v"})
	chk.False(found)
	chk.StrSlice(args, []string{"-v", "arg1", "arg2", "-v"})
	chk.Err(
		err,
		szargs.ErrAmbiguous.Error()+": -v found: 2 times",
	)
}

func TestSzArgs_ArgValueNonePresent(t *testing.T) {
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

func TestSzArgs_ArgValueAtBeginning(t *testing.T) {
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

func TestSzArgs_ArgValueMiddle(t *testing.T) {
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

func TestSzArgs_GetFlagValueEnd(t *testing.T) {
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

func TestSzArgs_ArgValueDuplicate(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := szargs.Arg("-n").Value(
		[]string{"-n", "firstName", "arg1", "arg2", "-n", "secondName"},
	)

	chk.Str(value, "")
	chk.False(found)
	chk.StrSlice(
		args,
		[]string{"-n", "firstName", "arg1", "arg2", "-n", "secondName"},
	)
	chk.Err(
		err,
		szargs.ErrAmbiguous.Error()+
			": '-n secondName' already set to: 'firstName'",
	)
}

func TestSzArgs_ArgValueMissing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := szargs.Arg("-n").Value(
		[]string{"arg1", "arg2", "-n"},
	)

	chk.Str(value, "")
	chk.False(found)
	chk.StrSlice(
		args,
		[]string{"arg1", "arg2", "-n"},
	)
	chk.Err(
		err,
		szargs.ErrMissing.Error()+
			": '-n value'",
	)
}

func TestSzArgs_ArgValuesNonePresent(t *testing.T) {
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

func TestSzArgs_ArgValuesAtBeginning(t *testing.T) {
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

func TestSzArgs_ArgValuesMiddle(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Arg("-n").Values(
		[]string{"arg1", "-n", "theName", "arg2"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzArgs_ArgValuesEnd(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Arg("-n").Values(
		[]string{"arg1", "arg2", "-n", "theName"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzArgs_ArgValuesDuplicate(t *testing.T) {
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

func TestSzArgs_ArgValuesMissing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := szargs.Arg("-n").Values(
		[]string{"arg1", "arg2", "-n"},
	)

	chk.StrSlice(value, nil)
	chk.StrSlice(
		args,
		[]string{"arg1", "arg2", "-n"},
	)
	chk.Err(
		err,
		szargs.ErrMissing.Error()+
			": '-n value'",
	)
}

const (
	tstEnv = "SZARGS_TESTING_ENVIRONMENT_VARIABLE"
	tstArg = "-t"
)

func TestSzArgs_ValueNoArgNoEnv(t *testing.T) {
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

func TestSzArgs_ValueArgAmbiguous(t *testing.T) {
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

func TestSzArgs_ValueArgMissing(t *testing.T) {
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

func TestSzArgs_ValueEnv(t *testing.T) {
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

func TestSzArgs_ValueEnvAndArg(t *testing.T) {
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
