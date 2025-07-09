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
