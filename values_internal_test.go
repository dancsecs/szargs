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

func TestSzargs_ValuesNonePresent(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := Flag("-n").values(nil)

	chk.StrSlice(value, nil)
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	value, args, err = Flag("-n").values(
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

	value, args, err := Flag("-n").values(
		[]string{"-n", "theName"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{})
	chk.NoErr(err)

	value, args, err = Flag("-n").values(
		[]string{"-n", "theName", "arg1", "arg2"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargs_ValuesMiddle(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := Flag("-n").values(
		[]string{"arg1", "-n", "theName", "arg2"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargs_ValuesEnd(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := Flag("-n").values(
		[]string{"arg1", "arg2", "-n", "theName"},
	)

	chk.StrSlice(value, []string{"theName"})
	chk.StrSlice(args, []string{"arg1", "arg2"})
	chk.NoErr(err)
}

func TestSzargs_ValuesDuplicate(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, err := Flag("-n").values(
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

	value, args, err := Flag("-n").values(
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
