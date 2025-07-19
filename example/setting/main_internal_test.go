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
	"A simple demo of a setting." +
	"\n\n" +
	"Usage: programName " +
	"[-t | --temp {c,f}]" +
	"\n\n" +
	"[-t | --temp {c,f}]" +
	"\n" +
	"Temperature measurement to use (celsius or fahrenheit)." +
	""

func Test_PASS_Default(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Using 'c' for temperatures.",
	)
}

func Test_PASS_Environment(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName")

	mainSetEnv()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Using 'f' for temperatures.",
	)
}

func Test_PASS_Argument(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "-t", "f")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Using 'f' for temperatures.",
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
