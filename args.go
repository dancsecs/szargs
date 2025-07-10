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
	"errors"
	"fmt"
	"strings"
)

// Exported errors.
var (
	ErrAmbiguous  = errors.New("ambiguous argument")
	ErrMissing    = errors.New("missing argument")
	ErrUnexpected = errors.New("unexpected argument")
)

// Arg represents a single argument.
type Arg string

func (a Arg) argIs(arg string) bool {
	for _, v := range strings.Split(string(a), "|") {
		if strings.TrimSpace(v) == arg {
			return true
		}
	}

	return false
}

// Count scans argument array (args) removing and counting the number of
// times the argument is encountered.
func (a Arg) Count(args []string) (int, []string) {
	count := 0
	cleanedArgs := make([]string, 0, len(args))

	for _, arg := range args {
		if a.argIs(arg) {
			count++
		} else {
			cleanedArgs = append(cleanedArgs, arg)
		}
	}

	return count, cleanedArgs
}

// Is scans the args returning true if found after removing the arg from
// the list.  If the argument appears twice then and ambiguous error will be
// returned with the arg array untouched.
func (a Arg) Is(args []string) (bool, []string, error) {
	count, cleanedArgs := a.Count(args)
	if count > 1 {
		return false, args,
			fmt.Errorf(
				"%w: %s found: %d times",
				ErrAmbiguous,
				a,
				count,
			)
	}

	return count == 1, cleanedArgs, nil
}

// Done insures that there are no further args exist returning an error
// if any are found.
func Done(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("%w: [%v]",
			ErrUnexpected,
			strings.Join(args, " "),
		)
	}

	return nil
}
