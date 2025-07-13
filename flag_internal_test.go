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
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztestlog"
)

func TestSzArgs_ArgCount(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	var (
		arg   = argFlag("-v | --verbose")
		count int
		args  []string
	)

	count, args = arg.count(args)
	chk.Int(count, 0)
	chk.StrSlice(args, nil)

	count, args = arg.count([]string{"arg1", "arg2"})
	chk.Int(count, 0)
	chk.StrSlice(args, []string{"arg1", "arg2"})

	count, args = arg.count([]string{"-v", "arg1", "arg2"})
	chk.Int(count, 1)
	chk.StrSlice(args, []string{"arg1", "arg2"})

	count, args = arg.count([]string{"arg1", "-v", "arg2"})
	chk.Int(count, 1)
	chk.StrSlice(args, []string{"arg1", "arg2"})

	count, args = arg.count([]string{"arg1", "arg2", "--verbose"})
	chk.Int(count, 1)
	chk.StrSlice(args, []string{"arg1", "arg2"})

	count, args = arg.count([]string{"-v", "arg1", "-v", "arg2", "-v"})
	chk.Int(count, 3)
	chk.StrSlice(args, []string{"arg1", "arg2"})

	count, args = arg.count(
		[]string{
			"--verbose",
			"-v",
			"--verbose",
			"arg1",
			"--verbose",
			"-v",
			"--verbose",
			"arg2",
			"--verbose",
			"-v",
			"--verbose",
		},
	)
	chk.Int(count, 9)
	chk.StrSlice(args, []string{"arg1", "arg2"})
}

func TestSzArgs_Is(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	var (
		err   error
		arg   = argFlag("-v")
		found bool
		args  []string
	)

	found, args, err = arg.is(nil)
	chk.False(found)
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	found, args, err = arg.is([]string{"arg1"})
	chk.False(found)
	chk.StrSlice(args, []string{"arg1"})
	chk.NoErr(err)

	found, args, err = arg.is([]string{"arg1", "arg2"})
	chk.False(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)

	found, args, err = arg.is([]string{"-v", "arg1", "arg2"})
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)

	found, args, err = arg.is([]string{"arg1", "-v", "arg2"})
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)

	found, args, err = arg.is([]string{"arg1", "arg2", "-v"})
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)

	found, args, err = arg.is([]string{"-v", "arg1", "arg2", "-v"})
	chk.False(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.Err(
		err,
		ErrAmbiguous.Error()+": '-v' found 2 times",
	)
}

func TestSzargs_ValueNonePresent(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := argFlag("[-n theName]").value(nil)

	chk.Str(value, "")
	chk.False(found)
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	value, found, args, err = argFlag("[-n theName]").value(
		[]string{"arg1", "arg2"},
	)

	chk.Str(value, "")
	chk.False(found)
	chk.StrSlice(args,
		[]string{"arg1", "arg2"},
	)
	chk.NoErr(err)
}

func TestSzargs_ValueBeginning(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := argFlag("[-n theName]").value(
		[]string{"-n", "theName"},
	)

	chk.Str(value, "theName")
	chk.True(found)
	chk.StrSlice(args, []string{})
	chk.NoErr(err)

	value, found, args, err = argFlag("-n").value(
		[]string{"-n", "theName", "arg1", "arg2"},
	)

	chk.Str(value, "theName")
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargs_ValueMiddle(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := argFlag("-n").value(
		[]string{"arg1", "-n", "theName", "arg2"},
	)

	chk.Str(value, "theName")
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargs_ValueEnd(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := argFlag("-n").value(
		[]string{"arg1", "arg2", "-n", "theName"},
	)

	chk.Str(value, "theName")
	chk.True(found)
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargs_ValueDuplicate(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := argFlag("-n").value(
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
		ErrAmbiguous.Error()+
			": '-n secondName' already set to: 'firstName'",
	)
}

func TestSzargs_ValueTriplicate(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := argFlag("-n").value(
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
		ErrAmbiguous.Error()+
			": '-n secondName' already set to: 'firstName'"+
			": "+
			ErrAmbiguous.Error()+
			": '-n thirdName' already set to: 'firstName'"+
			"",
	)
}

func TestSzargs_ValueMissing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := argFlag("-n").value(
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
		ErrMissing.Error()+
			": '-n value'",
	)
}

func TestSzargs_ValuesNonePresent(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := argFlag("-n").values(nil)

	chk.StrSlice(value, nil)
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	value, args, err = argFlag("-n").values(
		[]string{"arg1", "arg2"},
	)

	chk.StrSlice(value, nil)
	chk.StrSlice(args,
		[]string{"arg1", "arg2"},
	)
	chk.NoErr(err)
}

func TestSzargs_ValuesAtBeginning(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := argFlag("-n").values(
		[]string{"-n", "theName"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{})
	chk.NoErr(err)

	value, args, err = argFlag("-n").values(
		[]string{"-n", "theName", "arg1", "arg2"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargs_ValuesMiddle(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := argFlag("-n").values(
		[]string{"arg1", "-n", "theName", "arg2"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargs_ValuesEnd(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := argFlag("-n").values(
		[]string{"arg1", "arg2", "-n", "theName"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargs_ValuesDuplicate(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := argFlag("-n").values(
		[]string{"-n", "firstName", "arg1", "arg2", "-n", "secondName"},
	)

	chk.StrSlice(value, []string{"firstName", "secondName"})
	chk.StrSlice(
		args,
		[]string{"arg1", "arg2"},
	)
	chk.NoErr(err)
}

func TestSzargs_ValuesMissing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := argFlag("-n").values(
		[]string{"arg1", "arg2", "-n"},
	)

	chk.StrSlice(value, nil)
	chk.StrSlice(
		args,
		[]string{"arg1", "arg2"},
	)
	chk.Err(
		err,
		ErrMissing.Error()+
			": '-n value'",
	)
}
