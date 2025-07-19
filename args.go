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
	"strings"
)

// Args provides a single point to access and extract program arguments.
type Args struct {
	usageDefined map[string]bool
	usageHeader  string
	usageBody    string
	programName  string
	programDesc  string
	args         []string
	err          error
}

// New creates a new Args object based in the arguments passed.  The first
// element of the arguments must be the program name.
func New(programDesc string, args []string) *Args {
	if len(args) < 1 {
		return &Args{
			usageDefined: make(map[string]bool),
			usageHeader:  "",
			usageBody:    "",
			programName:  "NotDefined",
			programDesc:  programDesc,
			args:         nil,
			err:          ErrNoArgs,
		}
	}

	var myArgs []string

	if len(args) > 1 {
		myArgs = make([]string, len(args)-1)
		copy(myArgs, args[1:])
	}

	return &Args{
		usageDefined: make(map[string]bool),
		usageHeader:  "",
		usageBody:    "",
		programName:  filepath.Base(args[0]),
		programDesc:  programDesc,
		args:         myArgs,
		err:          nil,
	}
}

func (args *Args) addUsage(item, desc string) {
	if !args.usageDefined[item] {
		args.usageHeader += " " + item
		args.usageBody += "\n\n" + item + "\n" + desc
		args.usageDefined[item] = true
	}
}

// PushErr adds the provided error if not nil to the Args error stack.
func (args *Args) PushErr(err error) {
	if err != nil {
		if args.err == nil {
			args.err = err
		} else {
			args.err = fmt.Errorf("%w: %w", args.err, err)
		}
	}
}

// Err returns any errors encountered while parsing the arguments.
func (args *Args) Err() error {
	return args.err
}

// HasErr returns any errors encountered while parsing the arguments.
func (args *Args) HasErr() bool {
	return args.err != nil
}

// HasNext returns true if any arguments remain unabsorbed.
func (args *Args) HasNext() bool {
	return len(args.args) > 0
}

// PushArg places the supplied argument to the end of the internal ags list.
func (args *Args) PushArg(arg string) {
	args.args = append(args.args, arg)
}

// Args returns a copy of the current argument list.
func (args *Args) Args() []string {
	cpy := make([]string, len(args.args))
	copy(cpy, args.args)

	return cpy
}

// Usage returns a usage message based on the parsed arguments.
func (args *Args) Usage() string {
	return args.programName + "\n" +
		args.programDesc + "\n" +
		"\n" +
		"Usage: " + args.programName + args.usageHeader +
		args.usageBody
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

// Done returns an error if there are any remaining arguments.
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
