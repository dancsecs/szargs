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
	"A simple demo of repeating a string." +
	"\n\n" +
	"Usage: programName" +
	" message [times]" +
	"\n\n" +
	"message" +
	"\n" +
	"What to repeat." +
	"\n\n" +
	"[times]" +
	"\n" +
	"The number of times to repeat.  Defaults to 3." +
	""

func Test_PASS_DefaultCount(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "what to repeat")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"what to repeat",
		"what to repeat",
		"what to repeat",
	)
}

func Test_PASS_ValueByte(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "what to repeat", "2")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"what to repeat",
		"what to repeat",
	)
}

// Failing test.
func Test_FAIL_UnknownArgument(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "what to repeat", "5", "unknownArgument")

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

func Test_FAIL_NoArguments(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName")

	main()

	chk.Log()
	chk.Stderr(
		chk.ErrChain(
			"Error",
			szargs.ErrMissing,
			"message",
		),
		"",
		usageText,
	)
	chk.Stdout()
}
