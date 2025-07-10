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
line arguments ([]string) and environment variables. Arguments come in two
types:  Flagged and positional.

Flagged arguments are generally prefixed with a single dash "-" for single
letter flag or a double dash "--" for an extended spelled out flag.  A flag
can be stand alone in the case of a boolean such as "-v" indicating a verbose
setting or can be followed by and argument such as "--dir theDirectory".  Once
all flags have been processed then the remaining arguments can be identified
by their ordering.
*/
package szargs
