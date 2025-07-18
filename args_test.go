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
	"strings"
	"testing"

	"github.com/dancsecs/szargs"
	"github.com/dancsecs/sztestlog"
)

func TestSzargs_New_NoArgs(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	args := szargs.New("description", nil)

	chk.Err(
		args.Err(),
		szargs.ErrNoArgs.Error(),
	)

	chk.StrSlice(
		strings.Split(args.Usage(), "\n"),
		[]string{
			"NotDefined",
			"description",
			"",
			"Usage: NotDefined",
		},
	)

	chk.Log()
	chk.Stderr()
	chk.Stdout()
}

func TestSzargs_New_PushArgs(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	args := szargs.New("description", []string{"programName"})

	if !args.HasNext() {
		args.PushArg("theDefaultOptionalArg")
	}

	arg := args.NextString("theOptionalArg", "An optional string argument.")

	chk.Str(arg, "theDefaultOptionalArg")

	args.Done()

	chk.NoErr(args.Err())

	chk.Log()
	chk.Stderr()
	chk.Stdout()
}

func TestSzargs_New_JustProgramName(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	args := szargs.New("description", []string{"noProgName"})

	chk.Int(
		args.Count("[ -v | --verbose]", "how chatty should I be"),
		0,
	)

	chk.False(args.Is("[-f|--flag]", "a test flag"))

	args.Done()
	chk.NoErr(args.Err())

	chk.StrSlice(
		strings.Split(args.Usage(), "\n"),
		[]string{
			"noProgName",
			"description",
			"",
			"Usage: noProgName [ -v | --verbose] [-f|--flag]",
			"",
			"[ -v | --verbose]",
			"how chatty should I be",
			"",
			"[-f|--flag]",
			"a test flag",
		},
	)

	chk.Log()
	chk.Stderr()
	chk.Stdout()
}

func TestSzargs_New_AmbiguousIsName(t *testing.T) {
	chk := sztestlog.CaptureAll(t)
	defer chk.Release()

	args := szargs.New("description", []string{
		"noProgName",
		"-f", "FirstOccurrence",
		"--flag", "SecondOccurrence",
	})

	chk.Int(
		args.Count("[ -v | --verbose]", "how chatty should I be"),
		0,
	)

	chk.False(args.Is("[-f|--flag]", "a test flag"))

	args.Done()

	chk.StrSlice(
		strings.Split(args.Usage(), "\n"),
		[]string{
			"noProgName",
			"description",
			"",
			"Usage: noProgName [ -v | --verbose] [-f|--flag]",
			"",
			"[ -v | --verbose]",
			"how chatty should I be",
			"",
			"[-f|--flag]",
			"a test flag",
		},
	)

	chk.True(args.HasNext())
	chk.True(args.HasError())
	chk.Err(
		args.Err(),
		chk.ErrChain(
			szargs.ErrAmbiguous,
			"'[-f|--flag]' found 2 times",
			szargs.ErrUnexpected,
			"[FirstOccurrence SecondOccurrence]",
		),
	)

	chk.Log()
	chk.Stderr()
	chk.Stdout()
}
