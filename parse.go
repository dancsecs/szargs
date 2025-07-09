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
	"strconv"
)

const (
	bitsDefault = 0
	bits8       = 8
	bits16      = 16
	bits32      = 32
	bits64      = 64
	base10      = 10
)

func makeParseErr(rootErr, err error, name, str string) error {
	var parseErr error

	if errors.Is(err, strconv.ErrRange) {
		parseErr = ErrRange
	} else {
		parseErr = ErrSyntax
	}

	return fmt.Errorf("%w: %w: %s: '%s'", rootErr, parseErr, name, str)
}

func parseFloat64(name, str string) (float64, error) {
	result, err := strconv.ParseFloat(str, bits64)
	if err != nil {
		err = makeParseErr(ErrInvalidFloat64, err, name, str)
	}

	return result, err
}

func parseFloat32(name, str string) (float32, error) {
	result, err := strconv.ParseFloat(str, bits32)
	if err != nil {
		err = makeParseErr(ErrInvalidFloat32, err, name, str)
	}

	return float32(result), err
}

func parseUint64(name, str string) (uint64, error) {
	result, err := strconv.ParseUint(str, base10, bits64)
	if err != nil {
		err = makeParseErr(ErrInvalidUint64, err, name, str)
	}

	return result, err
}

func parseUint32(name, str string) (uint32, error) {
	result, err := strconv.ParseUint(str, base10, bits32)
	if err != nil {
		err = makeParseErr(ErrInvalidUint32, err, name, str)
	}

	return uint32(result), err
}

func parseUint16(name, str string) (uint16, error) {
	result, err := strconv.ParseUint(str, base10, bits16)
	if err != nil {
		err = makeParseErr(ErrInvalidUint16, err, name, str)
	}

	return uint16(result), err
}

func parseUint8(name, str string) (uint8, error) {
	result, err := strconv.ParseUint(str, base10, bits8)
	if err != nil {
		err = makeParseErr(ErrInvalidUint8, err, name, str)
	}

	return uint8(result), err
}

func parseUint(name, str string) (uint, error) {
	result, err := strconv.ParseUint(str, base10, bitsDefault)
	if err != nil {
		err = makeParseErr(ErrInvalidUint, err, name, str)
	}

	return uint(result), err
}

func parseInt64(name, str string) (int64, error) {
	result, err := strconv.ParseInt(str, base10, bits64)
	if err != nil {
		err = makeParseErr(ErrInvalidInt64, err, name, str)
	}

	return result, err
}

func parseInt32(name, str string) (int32, error) {
	result, err := strconv.ParseInt(str, base10, bits32)
	if err != nil {
		err = makeParseErr(ErrInvalidInt32, err, name, str)
	}

	return int32(result), err
}

func parseInt16(name, str string) (int16, error) {
	result, err := strconv.ParseInt(str, base10, bits16)
	if err != nil {
		err = makeParseErr(ErrInvalidInt16, err, name, str)
	}

	return int16(result), err
}

func parseInt8(name, str string) (int8, error) {
	result, err := strconv.ParseInt(str, base10, bits8)
	if err != nil {
		err = makeParseErr(ErrInvalidInt8, err, name, str)
	}

	return int8(result), err
}

func parseInt(name, str string) (int, error) {
	result, err := strconv.ParseInt(str, base10, bitsDefault)
	if err != nil {
		err = makeParseErr(ErrInvalidInt, err, name, str)
	}

	return int(result), err
}
