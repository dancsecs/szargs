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

/*
Package szargs provides simple methods of setting variables based on command
line arguments ([]string) and environment variables. Arguments come in three
types:  Flagged, positional and settings.

Flagged arguments are generally prefixed with a single dash "-" for a single
letter flag or a double dash "--" for an extended spelled out flag.  A flag
can be stand alone in the case of a boolean such as "-v" indicating a verbose
setting or can be followed by and argument such as "--dir theDirectory".  Once
all flags have been processed then the remaining arguments can be identified
by their ordering.

Positional Arguments are defined by their relative position in the argument
list once all of the flagged arguments have been removed.  There are two forms
the Next and the Last variations with the last insuring that there are no more
arguments in the list.

Settings combine a flagged argument, an environment variable and a default
permitting a setting to be defined as a default that can be overridden by an
environment variable setting that can also be overridden by a command line
flagged argument.

Each area provides for built in parsing to standard go data types.
*/
package szargs
