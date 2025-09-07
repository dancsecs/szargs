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
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

const defaultLineWidth = 75

// Args provides a single point to access and extract program arguments.
type Args struct {
	usageDefined  map[string]bool
	usageHeader   string
	usageSynopsis []string
	usageBody     string
	lineWidth     int
	programName   string
	programDesc   string
	args          []string
	err           error
}

var reIsGroup = regexp.MustCompile(`^-[A-Za-z]+$`)

func makeArgList(arg string) []string {
	if reIsGroup.MatchString(arg) {
		list := make([]string, 0, len(arg)-1)
		for _, option := range arg[1:] {
			list = append(list, "-"+string(option))
		}

		return list
	}

	return []string{arg}
}

// New creates a new Args object based in the arguments passed.  The first
// element of the arguments must be the program name.
func New(programDesc string, args []string) *Args {
	if len(args) < 1 {
		return &Args{
			usageDefined:  make(map[string]bool),
			usageHeader:   "    NotDefined",
			usageSynopsis: nil,
			usageBody:     "",
			programName:   "NotDefined",
			programDesc:   programDesc,
			lineWidth:     defaultLineWidth,
			args:          nil,
			err:           ErrNoArgs,
		}
	}

	var myArgs []string

	for _, arg := range args[1:] {
		myArgs = append(myArgs, makeArgList(arg)...)
	}

	return &Args{
		usageDefined:  make(map[string]bool),
		usageHeader:   "    " + filepath.Base(args[0]),
		usageSynopsis: nil,
		usageBody:     "",
		programName:   filepath.Base(args[0]),
		programDesc:   programDesc,
		lineWidth:     defaultLineWidth,
		args:          myArgs,
		err:           nil,
	}
}

func (args *Args) addBodyLine(lineToAdd string) {
	lineAsAdded := "       " // Just 7. First added word will make it 8.
	for _, wrd := range strings.Split(
		strings.TrimSpace(lineToAdd),
		" ",
	) {
		if len(lineAsAdded)+len(wrd)+1 < args.lineWidth {
			lineAsAdded += " " + wrd
		} else {
			args.usageBody += lineAsAdded + "\n"
			lineAsAdded = "        " + wrd
		}
	}

	args.usageBody += lineAsAdded + "\n"
}

func (args *Args) addUsage(item, desc string) {
	if !args.usageDefined[item] {
		lines := strings.Split(args.usageHeader, "\n")
		line := lines[len(lines)-1]

		if len(line)+len(item)+1 < args.lineWidth {
			args.usageHeader += " " + item
		} else {
			args.usageHeader += "" +
				"\n" +
				strings.Repeat(
					" ",
					5+len(args.programName), //nolint:mnd // Ok.
				) +
				item
		}

		args.usageBody += "\n\n    " + item
		for _, line = range strings.Split(desc, "\n") {
			args.usageBody += "\n"
			args.addBodyLine(line)
		}

		args.usageDefined[item] = true
	}
}

// PushErr registers the provided error if not nil to the Args error stack.
func (args *Args) PushErr(err error) {
	if err != nil {
		if args.err == nil {
			args.err = err
		} else {
			args.err = fmt.Errorf("%w: %w", args.err, err)
		}
	}
}

// Err returns any errors encountered or registered while parsing the
// arguments.
func (args *Args) Err() error {
	return args.err
}

// HasErr returns true if any errors have been encountered or registered.
func (args *Args) HasErr() bool {
	return args.err != nil
}

// HasNext returns true if any arguments remain unabsorbed.
func (args *Args) HasNext() bool {
	return len(args.args) > 0
}

// PushArg places the supplied argument to the end of the internal args list.
func (args *Args) PushArg(arg string) {
	args.args = append(args.args, arg)
}

// Args returns a copy of the current argument list.
func (args *Args) Args() []string {
	cpy := make([]string, len(args.args))
	copy(cpy, args.args)

	return cpy
}

// UsageWidth sets the width the usage message.  Must be called before any
// options or arguments are removed otherwise the width will only apply to
// subsequent additions.
func (args *Args) UsageWidth(newLineWidth int) {
	args.lineWidth = newLineWidth
}

//		Usage returns a usage message based on the parsed arguments.
//	   /*
//	   # Usage
//
//	   Golang to 'github' markdown.
//
//	       gotomd [options] [path ...]
//
//	   	[-v | --verbose ...]
//	   		Provide more information when processing.
//
//	   */
func (args *Args) Usage() string {
	var header string

	if len(args.usageSynopsis) > 0 {
		for _, s := range args.usageSynopsis {
			header += "\n    " + args.programName + " " + s
		}
	} else {
		header = "\n" + args.usageHeader
	}

	return "# Usage: " + args.programName +
		"\n\n" +
		args.programDesc +
		"\n" +
		header +
		args.usageBody
}

// AddSynopsis includes another static synopsis message.
func (args *Args) AddSynopsis(s string) {
	args.usageSynopsis = append(args.usageSynopsis, s)
}

// Count returns the number of times the flag appears.
func (args *Args) Count(flag, desc string) int {
	var count int

	args.addUsage(flag, desc)

	count, args.args = argFlag(flag).count(args.args)

	return count
}

// Is returns true if the flag is present one and only one time.
func (args *Args) Is(flag, desc string) bool {
	var (
		found bool
		err   error
	)

	args.addUsage(flag, desc)

	found, args.args, err = argFlag(flag).is(args.args)
	if err != nil {
		args.PushErr(err)
	}

	return found
}

// Done registers an error if there are any remaining arguments.
func (args *Args) Done() {
	if len(args.args) > 0 {
		args.PushErr(
			fmt.Errorf("%w: [%v]",
				ErrUnexpected,
				strings.Join(args.args, " "),
			),
		)
	}
}

// ProgramName returns the configured program name.
func (args *Args) ProgramName() string {
	return args.programName
}
