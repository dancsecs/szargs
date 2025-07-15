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

func Test_PASS_NothingToDoAdd(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Name Not Found.",
		"Byte Not Found.",
	)
}

func Test_PASS_ValueName(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "--name", "theName")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Name Found: theName.",
		"Byte Not Found.",
	)
}

func Test_PASS_ValueByte(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "--byte", "35")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Name Not Found.",
		"Byte Found: 35.",
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
	)
	chk.Stdout(
		"Name Not Found.",
		"Byte Not Found.",
	)
}

func Test_FAIL_InvalidExtraByte(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	chk.SetArgs("programName", "-b", "22", "--byte", "35")

	main()

	chk.Log()
	chk.Stderr(
		chk.ErrChain(
			"Error",
			szargs.ErrAmbiguous,
			"'[-b | --byte]' for '35' already set to: '22'",
		),
	)
	chk.Stdout(
		"Name Not Found.",
		"Byte Not Found.",
	)
}
