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
	"os"
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

// Count scans argument array (args) removing and counting the number of
// times the argument is encountered.
func (a Arg) Count(args []string) (int, []string) {
	count := 0
	cleanedArgs := make([]string, 0, len(args))

	for _, arg := range args {
		if arg == string(a) {
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

// Value scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from th argument list.  If there is no next arg or the flag appears more
// than once an error is returned.  If an error occurs then the original
// arg array is returned.
func (a Arg) Value(args []string) (string, bool, []string, error) {
	found := false
	value := ""
	cleanedArgs := make([]string, 0, len(args))

	for i, mi := 0, len(args); i < mi; i++ {
		if args[i] == string(a) {
			if (i + 1) >= mi {
				return "", false, args,
					fmt.Errorf(
						"%w: '%s value'",
						ErrMissing,
						a,
					)
			}

			if found {
				return "", false, args,
					fmt.Errorf(
						"%w: '%s %s' already set to: '%s'",
						ErrAmbiguous,
						a,
						args[i+1],
						value,
					)
			}

			i++
			value = args[i]
			found = true
		} else {
			cleanedArgs = append(cleanedArgs, args[i])
		}
	}

	return value, found, cleanedArgs, nil
}

// Values scans the args looking for the specified flag.  If it finds
// it then the next arg as the value absorbing both the flag the value
// from th argument list.  If there is no next arg an error is returned and
// the original arg array is returned.
func (a Arg) Values(args []string) ([]string, []string, error) {
	values := []string(nil)
	cleanedArgs := make([]string, 0, len(args))

	for i, mi := 0, len(args); i < mi; i++ {
		if args[i] == string(a) {
			if (i + 1) >= mi {
				return nil, args,
					fmt.Errorf(
						"%w: '%s value'",
						ErrMissing,
						a,
					)
			}

			i++
			values = append(values, args[i])
		} else {
			cleanedArgs = append(cleanedArgs, args[i])
		}
	}

	return values, cleanedArgs, nil
}

// Next is a helper method to consume the next arg returning an error
// if there are no more args.  Otherwise it extracts the first argument
// returning the modified list.
func Next(name string, args []string) (string, []string, error) {
	if len(args) < 1 {
		return "", args, fmt.Errorf(
			"%w: %s", ErrMissing, name,
		)
	}

	return args[0], args[1:], nil
}

// Last is a helper method to consume the next arg returning an error
// if there are no more args or any extra args.  Otherwise it extracts the
// argument.
func Last(name string, args []string) (string, error) {
	value, cleanedArgs, err := Next(name, args)

	if err == nil {
		err = Done(cleanedArgs)
	}

	if err != nil {
		return "", err
	}

	return value, nil
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

// Value is a convenience function that selects a values from a default that
// can be overridden by an optional environment variable or an optional
// argument.  If the argument is found its value is chosen and removed from
// the returned argument list.  An error is returned if the argument is missing
// or ambiguous.
func Value(
	defaultValue, envOverride, argOverride string,
	args []string,
) (string, []string, error) {
	value, found, cleanArgs, err := Arg(argOverride).Value(args)
	if err != nil {
		return "", args, err
	}

	if !found {
		value = defaultValue

		if envOverride != "" {
			envValue, ok := os.LookupEnv(envOverride)
			if ok {
				value = envValue
			}
		}
	}

	return value, cleanArgs, nil
}
