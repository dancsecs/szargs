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

//nolint:dupl,funlen // Ok.
package szargs

import (
	"strings"
	"testing"

	"github.com/dancsecs/sztestlog"
)

func TestSzargs_Reflow1(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	width := 1

	chk.StrSlice(
		strings.Split(reflowLine("", "", width), "\n"),
		[]string{
			"",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a ", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a b", width), "\n"),
		[]string{
			"a",
			"b",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a bc d", width), "\n"),
		[]string{
			"a",
			"bc",
			"d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bc d", width), "\n"),
		[]string{
			">a",
			" bc",
			" d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bcd ef g", width), "\n"),
		[]string{
			">a",
			" bcd",
			" ef",
			" g",
		},
	)
}

func TestSzargs_Reflow2(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	width := 2

	chk.StrSlice(
		strings.Split(reflowLine("", "", width), "\n"),
		[]string{
			"",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a ", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a b", width), "\n"),
		[]string{
			"a",
			"b",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a bc d", width), "\n"),
		[]string{
			"a",
			"bc",
			"d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bc d", width), "\n"),
		[]string{
			">a",
			" bc",
			" d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bcd ef g", width), "\n"),
		[]string{
			">a",
			" bcd",
			" ef",
			" g",
		},
	)
}

func TestSzargs_Reflow3(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	width := 3

	chk.StrSlice(
		strings.Split(reflowLine("", "", width), "\n"),
		[]string{
			"",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a ", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a b", width), "\n"),
		[]string{
			"a",
			"b",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a bc d", width), "\n"),
		[]string{
			"a",
			"bc",
			"d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bc d", width), "\n"),
		[]string{
			">a",
			" bc",
			" d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bcd ef g", width), "\n"),
		[]string{
			">a",
			" bcd",
			" ef",
			" g",
		},
	)
}

func TestSzargs_Reflow4(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	width := 4

	chk.StrSlice(
		strings.Split(reflowLine("", "", width), "\n"),
		[]string{
			"",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a ", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a b", width), "\n"),
		[]string{
			"a",
			"b",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a bc d", width), "\n"),
		[]string{
			"a",
			"bc",
			"d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bc d", width), "\n"),
		[]string{
			">a",
			" bc",
			" d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bcd ef g", width), "\n"),
		[]string{
			">a",
			" bcd",
			" ef",
			" g",
		},
	)
}

func TestSzargs_Reflow5(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	width := 5

	chk.StrSlice(
		strings.Split(reflowLine("", "", width), "\n"),
		[]string{
			"",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a ", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a b", width), "\n"),
		[]string{
			"a b",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a bc d", width), "\n"),
		[]string{
			"a",
			"bc",
			"d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bc d", width), "\n"),
		[]string{
			">a",
			" bc",
			" d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bcd ef g", width), "\n"),
		[]string{
			">a",
			" bcd",
			" ef",
			" g",
		},
	)
}

func TestSzargs_Reflow6(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	width := 6

	chk.StrSlice(
		strings.Split(reflowLine("", "", width), "\n"),
		[]string{
			"",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a ", width), "\n"),
		[]string{
			"a",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a b", width), "\n"),
		[]string{
			"a b",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine("", "a bc d", width), "\n"),
		[]string{
			"a bc",
			"d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bc d", width), "\n"),
		[]string{
			">a",
			" bc",
			" d",
		},
	)

	chk.StrSlice(
		strings.Split(reflowLine(">", "a bcd ef g", width), "\n"),
		[]string{
			">a",
			" bcd",
			" ef",
			" g",
		},
	)
}
