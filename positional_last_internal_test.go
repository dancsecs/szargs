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

func TestSzargs_Last(t *testing.T) {
	chk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer chk.Release()

	args := New("program description", []string{
		"programName",
	})

	arg := args.last("TestArg", "TestDesc")

	chk.Str(arg, "")
	chk.Err(
		args.Err(),
		"missing argument: TestArg",
	)

	args = New("program description", []string{
		"programName",
		"arg1",
		"arg2",
	})

	arg = args.last("TestArg", "TestDesc")

	chk.Str(arg, "")
	chk.Err(
		args.Err(),
		"unexpected argument: [arg2]",
	)

	args = New("program description", []string{
		"programName",
		"arg2",
	})

	arg = args.last("TestArg", "TestDesc")

	chk.Str(arg, "arg2")
	chk.NoErr(args.Err())
}
