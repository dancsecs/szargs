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

// Flag represents a single argument.
type Flag string

func (a Flag) argIs(arg string) bool {
	flagVersions := strings.Split(strings.Trim(string(a), "[]{}."), "|")
	lastFlagVersion := len(flagVersions) - 1

	for i, flg := range flagVersions {
		flg = strings.TrimSpace(flg)
		if i == lastFlagVersion {
			// Remove optional arg name.  IE: [-n theName]
			flg = strings.Split(flg, " ")[0]
		}

		if flg == arg {
			return true
		}
	}

	return false
}

// count scans argument array (args) removing and counting the number of
// times the argument is encountered.
func (a Flag) count(args []string) (int, []string) {
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

// is scans the args counting and removing the arg from the list.  If the
// argument appears more than once an ErrAmbiguous is returned.
func (a Flag) is(args []string) (bool, []string, error) {
	var count int

	count, args = a.count(args)
	if count > 1 {
		return false, args,
			fmt.Errorf(
				"%w: '%s' found %d times",
				ErrAmbiguous,
				a,
				count,
			)
	}

	return count == 1, args, nil
}

// Value scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from the argument list.  If there is no next arg or the flag appears more
// than once an error is returned.
func (a Flag) value(args []string) (string, bool, []string, error) {
	found := false
	value := ""
	cleanedArgs := make([]string, 0, len(args))
	err := error(nil)

	pushErr := func(newErr error) {
		if err == nil {
			err = newErr
		} else {
			err = fmt.Errorf("%w: %w", err, newErr)
		}
	}

	for i, mi := 0, len(args); i < mi; i++ {
		if a.argIs(args[i]) { //nolint:nestif // Ok.
			if (i + 1) >= mi {
				pushErr(
					fmt.Errorf(
						"%w: '%s value'",
						ErrMissing,
						a,
					),
				)
			} else {
				i++
				if found {
					pushErr(
						fmt.Errorf(
							"%w: '%s %s' already set to: '%s'",
							ErrAmbiguous,
							a,
							args[i],
							value,
						),
					)
				} else {
					value = args[i]
					found = true
				}
			}
		} else {
			cleanedArgs = append(cleanedArgs, args[i])
		}
	}

	if err == nil {
		return value, found, cleanedArgs, nil
	}

	return "", false, cleanedArgs, err
}

// Values scans the args looking for all instances of the specified flag.  If
// it finds it then the next arg as the value absorbing both the flag the
// value from the argument list.
func (a Flag) values(args []string) ([]string, []string, error) {
	values := []string(nil)
	cleanedArgs := make([]string, 0, len(args))
	err := error(nil)

	for i, mi := 0, len(args); i < mi; i++ {
		if a.argIs(args[i]) {
			if (i + 1) >= mi {
				err = fmt.Errorf(
					"%w: '%s value'",
					ErrMissing,
					a,
				)
			} else {
				i++
				values = append(values, args[i])
			}
		} else {
			cleanedArgs = append(cleanedArgs, args[i])
		}
	}

	if err == nil {
		return values, cleanedArgs, nil
	}

	return nil, cleanedArgs, err
}
