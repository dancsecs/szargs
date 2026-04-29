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

func TestSzargs_New_NoArgsWithUsage(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	args := szargs.New("description", nil)

	chk.Err(
		args.Err(),
		szargs.ErrNoArgs.Error(),
	)

	chk.StrSlice(
		strings.Split(args.Usage(0), "\n"),
		[]string{
			"usage: NotDefined",
			"",
			"description",
		},
	)

	chk.Log()
	chk.Stderr()
	chk.Stdout()
}

//nolint:funlen // Ok.
func TestSzargs_New_JustProgramNameWithUsage(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	description := "" +
		"This description will demo\n" +
		"patching together shorter lines and then reformat\n" +
		"to fit in the usage width.\n\n" +
		"While this line starts anew."

	args := szargs.New(description, []string{"noProgName"})

	//  args.UsageWidth(78) // Otherwise -f description changes.

	chk.Int(
		args.Count("[ -v | --verbose]", "how chatty should I be"),
		0,
	)

	chk.False(
		args.Is(
			"[-f|--flag]",
			"a test flag\n\n"+
				"And a needlessly long sentence to over the eighty byte "+
				"threshold so we can see a result of wrapping a line at "+
				"eighty characters."+
				"",
		),
	)
	chk.False(args.Is("[-g|--group]", "a group flag"))
	chk.False(args.Is("[-h|--human]", "a human flag"))
	chk.False(args.Is("[-q|--quick_mode]", "magic mystery method"))
	chk.False(args.Is("[-o|--over]", "an over flag"))

	args.Done()
	chk.NoErr(args.Err())

	chk.StrSlice(
		strings.Split(args.Usage(50), "\n"),
		[]string{
			"usage: noProgName [ -v | --verbose] [-f|--flag]",
			"                  [-g|--group] [-h|--human]",
			"                  [-q|--quick_mode] [-o|--over]",
			"",
			"This description will demo patching together",
			"shorter lines and then reformat to fit in the",
			"usage width.",
			"",
			"While this line starts anew.",
			"",
			"    [ -v | --verbose]",
			"        how chatty should I be",
			"",
			"    [-f|--flag]",
			"        a test flag",
			"",
			"        And a needlessly long sentence to over",
			"        the eighty byte threshold so we can see",
			"        a result of wrapping a line at eighty",
			"        characters.",
			"",
			"    [-g|--group]",
			"        a group flag",
			"",
			"    [-h|--human]",
			"        a human flag",
			"",
			"    [-q|--quick_mode]",
			"        magic mystery method",
			"",
			"    [-o|--over]",
			"        an over flag",
		},
	)

	chk.Log()
	chk.Stderr()
	chk.Stdout()
}

func TestSzargs_New_AmbiguousIsNameWithUsage(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
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
		strings.Split(args.Usage(0), "\n"),
		[]string{
			"usage: noProgName [ -v | --verbose] [-f|--flag]",
			"",
			"description",
			"",
			"    [ -v | --verbose]",
			"        how chatty should I be",
			"",
			"    [-f|--flag]",
			"        a test flag",
		},
	)

	chk.True(args.HasNext())
	chk.True(args.HasErr())
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

func TestSzargs_Synopsis(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	args := szargs.New("description", []string{
		"noProgName",
		"-v",
	})

	args.AddSynopsis("[OPTION ...] [PATH ...]")

	chk.Int(
		args.Count("[ -v | --verbose]", "how chatty should I be"),
		1,
	)

	chk.False(args.Is("[-f|--flag]", "a test flag"))

	args.Done()

	chk.StrSlice(
		strings.Split(args.Usage(0), "\n"),
		[]string{
			"usage: noProgName [OPTION ...] [PATH ...]",
			"",
			"description",
			"",
			"    [ -v | --verbose]",
			"        how chatty should I be",
			"",
			"    [-f|--flag]",
			"        a test flag",
		},
	)

	chk.False(args.HasNext())
	chk.False(args.HasErr())

	chk.Log()
	chk.Stderr()
	chk.Stdout()
}

func TestSzargs_SynopsisTwo(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	args := szargs.New("description", []string{
		"noProgName",
		"-v",
	})

	args.AddSynopsis("[OPTION ...] [PATH ...]")
	args.AddSynopsis("help [OPTION]")

	chk.Int(
		args.Count("[ -v | --verbose]", "how chatty should I be"),
		1,
	)

	chk.False(args.Is("[-f|--flag]", "a test flag"))

	args.Done()

	chk.StrSlice(
		strings.Split(args.Usage(0), "\n"),
		[]string{
			"usage: noProgName [OPTION ...] [PATH ...]",
			"       noProgName help [OPTION]",
			"",
			"description",
			"",
			"    [ -v | --verbose]",
			"        how chatty should I be",
			"",
			"    [-f|--flag]",
			"        a test flag",
		},
	)

	chk.False(args.HasNext())
	chk.False(args.HasErr())

	chk.Log()
	chk.Stderr()
	chk.Stdout()
}

func TestSzargs_UsageWidth_NoTerminalDefault(t *testing.T) {
	chk := sztestlog.CaptureLogAndStderrAndStdout(t)
	defer chk.Release()

	args := szargs.New(""+
		"1234567890"+ // 10.
		"1234567890"+ // 20.
		"1234567890"+ // 30.
		"1234567890"+ // 40.
		"1234567890"+ // 50.
		"1234567890"+ // 60.
		"1234567890"+ // 70.
		" 2 4 6 8 0 "+ // Split at 77.
		"123456789"+ // 90.
		"",
		nil,
	)

	chk.Err(
		args.Err(),
		szargs.ErrNoArgs.Error(),
	)

	chk.StrSlice(
		strings.Split(args.Usage(-1), "\n"),
		[]string{
			"usage: NotDefined",
			"",
			"123456789012345678901234567890123456789012345678901234567890" +
				"1234567890 2 4 6",
			"8 0 123456789",
		},
	)

	chk.Log()
	chk.Stderr()
	chk.Stdout()
}
