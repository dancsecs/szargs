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
	"os"
	"runtime"
	"testing"

	"github.com/dancsecs/sztestlog"
)

func TestSzargs_UsageWidth_TerminalDefault(t *testing.T) {
	chk := sztestlog.CaptureNothing(t)
	defer chk.Release()

	if runtime.GOOS != "linux" {
		t.Skipf("unknown terminal path for GOOS %v", runtime.GOOS)
	}

	file, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	chk.NoErr(err)

	defer func() {
		_ = file.Close()
	}()

	args := New(""+
		"1234567890"+ // 10.
		"1234567890"+ // 20.
		"1234567890"+ // 30.
		"1234567890"+ // 40.
		"1234567890"+ // 50.
		"1234567890"+ // 60.
		"1234567890"+ // 70.
		" 2 4 6 8 0 "+ // Split at 77.
		"123456789"+ // 90.
		"",
		nil,
	)

	chk.Err(
		args.Err(),
		ErrNoArgs.Error(),
	)

	chk.Int(terminalWidth(file), defaultLineWidth)
}
