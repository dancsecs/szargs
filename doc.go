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
Package szargs provides a minimal and consistent interface for retrieving
settings from command-line arguments ([]string) and environment variables.

It supports three types of arguments:
  - Flagged: Identified by a single dash (e.g., "-v") for short flags, or a
    double dash (e.g., "--dir") for long-form flags. Flags may be standalone
    booleans or followed by a value.
  - Positional: Identified by their order in the argument list after all
    flagged arguments have been processed. Positional arguments can be
    retrieved in two forms: Next (in order) or Last (ensuring no trailing
    arguments remain).
  - Settings: A composite configuration mechanism that combines a default
    value, an environment variable, and a flagged argument—allowing each to
    override the previous in precedence: default < env < flag.

The package includes built-in parsers for standard Go data types.

Usage centers around the Args type, created using:

	szargs.New(programDesc string, programArgs []string)

The `programArgs` slice must include the program name as the first element;
this is ignored during argument parsing.

After retrieving all relevant arguments, the `Args.Done()` method must be
called to report an error if any unprocessed arguments remain.

This utility reflects a preference for simplicity and clarity in tooling. If
it helps your project flow a little more smoothly, it's done its job.

# Dedication

This project is dedicated to Reem.
Your brilliance, courage, and quiet strength continue to inspire me.
Every line is written in gratitude for the light and hope you brought into my
life.

NOTE: Documentation reviewed and polished with the assistance of ChatGPT from
OpenAI.
*/
package szargs
