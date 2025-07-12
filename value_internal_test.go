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

func TestSzargs_ValueNonePresent(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, found, args, err := Flag("-n").value(nil)

	chk.Str(value, "")
	chk.False(found)
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	value, found, args, err = Flag("-n").value(
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

	value, found, args, err := Flag("-n").value(
		[]string{"-n", "theName"},
	)

	chk.Str(value, "theName")
	chk.True(found)
	chk.StrSlice(args, []string{})
	chk.NoErr(err)

	value, found, args, err = Flag("-n").value(
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

	value, found, args, err := Flag("-n").value(
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

	value, found, args, err := Flag("-n").value(
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

	value, found, args, err := Flag("-n").value(
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

	value, found, args, err := Flag("-n").value(
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

	value, found, args, err := Flag("-n").value(
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
