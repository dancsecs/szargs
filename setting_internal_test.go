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

const (
	tstEnv = "SZARGS_TESTING_ENVIRONMENT_VARIABLE"
	tstArg = "-t"
)

func TestSzargs_ValueNoArgNoEnv(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, srcErr, err := setting(tstArg, tstEnv, "def", nil)

	chk.Str(value, "def")
	chk.StrSlice(args, nil)
	chk.NoErr(srcErr)
	chk.NoErr(err)

	value, args, srcErr, err = setting(tstArg, tstEnv, "def",
		[]string{"arg1", "arg2"},
	)

	chk.Str(value, "def")
	chk.StrSlice(args,
		[]string{"arg1", "arg2"},
	)
	chk.NoErr(srcErr)
	chk.NoErr(err)
}

func TestSzargs_ValueArgAmbiguous(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, srcErr, err := setting(tstArg, tstEnv, "def",
		[]string{tstArg, "first", "arg1", "arg2", tstArg, "second"},
	)

	chk.Str(value, "")
	chk.StrSlice(args,
		[]string{tstArg, "first", "arg1", "arg2", tstArg, "second"},
	)
	chk.Err(srcErr, ErrInvalidFlag.Error())
	chk.Err(
		err,
		ErrAmbiguous.Error()+
			": '-t second' already set to: 'first'",
	)
}

func TestSzargs_ValueArgMissing(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	value, args, srcErr, err := setting(tstArg, tstEnv, "def",
		[]string{"arg1", "arg2", tstArg},
	)

	chk.Str(value, "")
	chk.StrSlice(args,
		[]string{"arg1", "arg2", tstArg},
	)
	chk.Err(srcErr, ErrInvalidFlag.Error())
	chk.Err(
		err,
		ErrMissing.Error()+
			": '-t value'",
	)
}

func TestSzargs_ValueEnv(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	chk.SetEnv(tstEnv, "env")

	value, args, srcErr, err := setting(tstArg, tstEnv, "def",
		[]string{"arg1", "arg2"},
	)

	chk.Str(value, "env")
	chk.StrSlice(args,
		[]string{"arg1", "arg2"},
	)
	chk.Err(srcErr, ErrInvalidEnv.Error())
	chk.NoErr(err)
}

func TestSzargs_ValueEnvAndArg(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	chk.SetEnv(tstEnv, "env")

	value, args, srcErr, err := setting(tstArg, tstEnv, "def",
		[]string{"arg1", tstArg, "arg", "arg2"},
	)

	chk.Str(value, "arg")
	chk.StrSlice(args,
		[]string{"arg1", "arg2"},
	)
	chk.Err(srcErr, ErrInvalidFlag.Error())
	chk.NoErr(err)
}
