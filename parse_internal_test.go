/*
   Szerszam argument library: szargs.
   Copyright (C) 2025  Leslie Dancsecs

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
	"testing"

	"github.com/dancsecs/sztestlog"
)

func TestArgs_IntBase(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	const (
		base2  = 2
		base8  = 8
		base10 = 10
		base16 = 16
	)

	tst := func(rawStr, clnStr string, base int) {
		t.Helper()

		s, b := intBase(rawStr)
		chk.Str(s, clnStr)
		chk.Int(b, base)
	}

	// Base2
	tst("0B1", "1", base2)
	tst("0b101", "101", base2)

	// Base8
	tst("0O7", "7", base8)
	tst("0o123", "123", base8)

	// Base8 - Single leading 0.
	tst("00", "0", base8)
	tst("01", "1", base8)
	tst("0b", "b", base8)

	// Base10
	tst("0", "0", base10)
	tst("1", "1", base10)
	tst("123456", "123456", base10)

	// Base16
	tst("0Xa", "a", base16)
	tst("0xabc", "abc", base16)
}
