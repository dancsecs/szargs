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
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/term"
)

const (
	usagePrefix      = "usage: "
	defaultLineWidth = 79
)

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
			usageHeader:   "",
			usageSynopsis: nil,
			usageBody:     "",
			programName:   "NotDefined",
			programDesc:   prepareDesc("", programDesc),
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
		usageHeader:   "",
		usageSynopsis: nil,
		usageBody:     "",
		programName:   filepath.Base(args[0]),
		programDesc:   prepareDesc("", programDesc),
		lineWidth:     defaultLineWidth,
		args:          myArgs,
		err:           nil,
	}
}

// Prepare desc appends paragraphs (separated by blank lines) to single lines
// in preparation for reflowing to different widths.
func prepareDesc(prefix, desc string) string {
	var res strings.Builder

	lines := strings.Split(strings.Trim(desc, "\n"), "\n")
	first := 0

	for i, l := range lines {
		if l == "" {
			res.WriteString(
				prefix +
					strings.Join(lines[first:i], " ") +
					"\n\n",
			)

			first = i + 1
		}
	}

	if first < len(lines) {
		res.WriteString(prefix + strings.Join(lines[first:], " "))
	}

	return strings.TrimRight(res.String(), "\n")
}

// RegisterUsage registers a new flag and its description if and only if the
// flag has not been already  registered.
func (args *Args) RegisterUsage(item, desc string) {
	if !args.usageDefined[item] {
		args.usageHeader += " " + item
		args.usageBody += "\n" + item + "\n" +
			prepareDesc("    ", desc) + "\n"
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

func terminalWidth() int {
	fd := int(os.Stdout.Fd())

	if term.IsTerminal(fd) {
		w, _, err := term.GetSize(fd)
		if err == nil {
			return w
		}
	}

	return defaultLineWidth
}

// Usage returns a usage messages representing the Args object.  It is
// formatted to the lineWidth provided.  A zero uses the defaultLineWidth
// while a negative value caused an effort to determine if writing to a
// terminal and if so using its width otherwise defaulting.
func (args *Args) Usage(lineWidth int) string {
	var header string

	if lineWidth < 0 {
		lineWidth = terminalWidth()
	}

	if lineWidth < 1 {
		lineWidth = defaultLineWidth
	}

	args.lineWidth = lineWidth

	if len(args.usageSynopsis) > 0 {
		for i, synopsis := range args.usageSynopsis {
			if i == 0 {
				header += "\n" + usagePrefix
			} else {
				header += "\n" + strings.Repeat(" ", len(usagePrefix))
			}

			header += args.programName + " " + synopsis
		}
	} else {
		if args.usageHeader == "" {
			header = "\nusage: " + args.programName
		} else {
			header = "\n" + reflowLine("usage: "+args.programName+" ",
				args.usageHeader,
				args.lineWidth,
			)
		}
	}

	return strings.TrimRight(
		header[1:]+
			"\n\n"+
			reflowLine("", args.programDesc, args.lineWidth)+"\n\n"+
			reflowLines("    ", args.usageBody, args.lineWidth)+"\n",
		"\n",
	)
}

// AddSynopsis includes another static synopsis message.
func (args *Args) AddSynopsis(s string) {
	args.usageSynopsis = append(args.usageSynopsis, s)
}

// Count returns the number of times the flag appears.
func (args *Args) Count(flag, desc string) int {
	var count int

	args.RegisterUsage(flag, desc)

	count, args.args = argFlag(flag).count(args.args)

	return count
}

// Is returns true if the flag is present one and only one time.
func (args *Args) Is(flag, desc string) bool {
	var (
		found bool
		err   error
	)

	args.RegisterUsage(flag, desc)

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
