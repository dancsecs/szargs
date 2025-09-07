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

	"github.com/dancsecs/sztestlog"
)

func TestSzargs_Next(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	var (
		arg  string
		args []string
		err  error
	)

	arg, args, err = next("TestArg", args)

	chk.Str(arg, "")
	chk.StrSlice(args, nil)
	chk.Err(
		err,
		"missing argument: TestArg",
	)

	args = []string{"arg1", "arg2"}

	arg, args, err = next("TestArg", args)

	chk.Str(arg, "arg1")
	chk.StrSlice(args, []string{"arg2"})
	chk.NoErr(err)

	arg, args, err = next("TestArg", args)

	chk.Str(arg, "arg2")
	chk.StrSlice(args, nil)
	chk.NoErr(err)

	arg, args, err = next("TestArg", args)

	chk.Str(arg, "")
	chk.StrSlice(args, nil)
	chk.Err(
		err,
		ErrMissing.Error()+": TestArg",
	)
}
