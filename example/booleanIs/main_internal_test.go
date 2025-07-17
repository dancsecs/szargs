/*
   Szerszam argument library: szargs.
   Copyright (C) 2024, 2025  Leslie Dancsecs

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

package main

import (
	"testing"

	"github.com/dancsecs/szargs"
	"github.com/dancsecs/sztestlog"
)

const usageText = "" +
	"programName" +
	"\n" +
	"A simple utility to demo identifying a boolean flag." +
	"\n\n" +
	"Usage: programName" +
	" [-t | --true]" +
	"\n\n" +
	"[-t | --true]" +
	"\n" +
	"A single flag indicating true." +
	""

func Test_PASS_NothingToDoAdd(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Is NOT true.",
	)
}

func Test_PASS_IsTrueLong(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "--true")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Is true.",
	)
}

func Test_PASS_IsTrueShort(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "-t")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Is true.",
	)
}

// Failing test.
func Test_FAIL_UnknownArgument(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "unknownArgument")

	main()

	chk.Log()
	chk.Stderr(
		chk.ErrChain(
			"Error",
			szargs.ErrUnexpected,
			"[unknownArgument]",
		),
		"",
		usageText,
	)
	chk.Stdout()
}

func Test_FAIL_InvalidExtraTrue(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "-t", "--true")

	main()

	chk.Log()
	chk.Stderr(
		chk.ErrChain(
			"Error",
			szargs.ErrAmbiguous,
			"'[-t | --true]' found 2 times",
		),
		"",
		usageText,
	)
	chk.Stdout()
}
