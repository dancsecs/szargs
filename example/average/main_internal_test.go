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
	"programName\n" +
	"A simple utility to add or average a number list.\n" +
	"\n" +
	"Usage: programName" +
	" [-v | --verbose ...]" +
	" [-n | --number float64 ...]" +
	" [operation]" +
	"\n\n" +
	"[-v | --verbose ...]\n" +
	"The verbose level.\n" +
	"\n" +
	"[-n | --number float64 ...]\n" +
	"The numbers to act on.\n" +
	"\n" +
	"[operation]\n" +
	"The operation (add or average) defaulting to add.\n" +
	""

// Passing test.
func Test_PASS_NothingToDoAdd(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	chk.SetArgs("programName", "add")

	chk.NoPanic(main)

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Sum: 0.000000",
	)
}

// Passing test.
func Test_PASS_NothingToDoAverage(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	chk.SetArgs("programName", "average")

	chk.NoPanic(main)

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Avg: 0.000000",
	)
}

func Test_PASS_OneNumberAverage(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	chk.SetArgs("programName", "-n", "123.456789", "average")

	chk.NoPanic(main)

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Avg: 123.456789",
	)
}

func Test_PASS_ThreeNumberAverage(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	chk.SetArgs(
		"programName",
		"-n", "10",
		"--number", "20",
		"-n", "30",
		"average",
	)

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Avg: 20.000000",
	)
}

func Test_PASS_NoArgs(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	chk.SetArgs("programName")

	main()

	chk.Log()
	chk.Stderr()
	chk.Stdout(
		"Sum: 0.000000",
	)
}

// Failing test.
func Test_FAIL_MissingNumber(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	chk.SetArgs("programName", "-n")

	main()

	chk.Log()
	chk.Stderr(
		chk.ErrChain(
			"Error",
			szargs.ErrMissing,
			"'[-n | --number float64 ...]'",
		),
		"",
		usageText,
	)
	chk.Stdout()
}

func Test_FAIL_InvalidNumberDefaultOperation(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	chk.SetArgs("programName", "-n", "invalidNumber")

	main()

	chk.Log()
	chk.Stderr(
		chk.ErrChain(
			"Error",
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			"[-n | --number float64 ...]",
			"'invalidNumber'",
		),
		"",
		usageText,
	)
	chk.Stdout()
}

func Test_FAIL_InvalidNumber(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	chk.SetArgs("programName", "-n", "invalidNumber", "add")

	main()

	chk.Log()
	chk.Stderr(
		chk.ErrChain(
			"Error",
			szargs.ErrInvalidFloat64,
			szargs.ErrSyntax,
			"[-n | --number float64 ...]",
			"'invalidNumber'",
		),
		"",
		usageText,
	)
	chk.Stdout()
}

func Test_FAIL_UnexpectedArguments(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	chk.SetArgs("programName", "--count", "unexpectedArgument")

	main()

	chk.Log()
	chk.Stderr(
		chk.ErrChain(
			"Error",
			szargs.ErrInvalidOption,
			"'--count' ([operation] must be one of [add average])",
			szargs.ErrUnexpected,
			"[unexpectedArgument]",
		),
		"",
		usageText,
	)
	chk.Stdout()
}
