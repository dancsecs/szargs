/*
   Szerszam argument library: szargs.
   Copyright (C) 2024-2025  Leslie Dancsecs

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
	"strings"
)

const (
	bitsDefault = 0
	bits8       = 8
	bits16      = 16
	bits32      = 32
	bits64      = 64

	base2  = 2
	base8  = 8
	base10 = 10
	base16 = 16
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

func intBase(str string) (string, int) {
	if len(str) > 2 && str[0:1] == "0" {
		switch strings.ToLower(str[0:2]) {
		case "0b":
			return str[2:], base2
		case "0o":
			return str[2:], base8
		case "0x":
			return str[2:], base16
		}
	}

	if len(str) > 1 && str[0:1] == "0" {
		return str[1:], base8
	}

	return str, base10
}

func parseOption(
	name, str string, validOptions []string,
) (string, error) {
	for _, validOption := range validOptions {
		if str == validOption {
			return str, nil
		}
	}

	return "", fmt.Errorf(
		"%w: '%s' (%s must be one of %v)",
		ErrInvalidOption,
		str,
		name,
		validOptions,
	)
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
	str, base := intBase(str)

	result, err := strconv.ParseUint(str, base, bits64)
	if err != nil {
		err = makeParseErr(ErrInvalidUint64, err, name, str)
	}

	return result, err
}

func parseUint32(name, str string) (uint32, error) {
	str, base := intBase(str)

	result, err := strconv.ParseUint(str, base, bits32)
	if err != nil {
		err = makeParseErr(ErrInvalidUint32, err, name, str)
	}

	return uint32(result), err
}

func parseUint16(name, str string) (uint16, error) {
	str, base := intBase(str)

	result, err := strconv.ParseUint(str, base, bits16)
	if err != nil {
		err = makeParseErr(ErrInvalidUint16, err, name, str)
	}

	return uint16(result), err
}

func parseUint8(name, str string) (uint8, error) {
	str, base := intBase(str)

	result, err := strconv.ParseUint(str, base, bits8)
	if err != nil {
		err = makeParseErr(ErrInvalidUint8, err, name, str)
	}

	return uint8(result), err
}

func parseUint(name, str string) (uint, error) {
	str, base := intBase(str)

	result, err := strconv.ParseUint(str, base, bitsDefault)
	if err != nil {
		err = makeParseErr(ErrInvalidUint, err, name, str)
	}

	return uint(result), err
}

func parseInt64(name, str string) (int64, error) {
	str, base := intBase(str)

	result, err := strconv.ParseInt(str, base, bits64)
	if err != nil {
		err = makeParseErr(ErrInvalidInt64, err, name, str)
	}

	return result, err
}

func parseInt32(name, str string) (int32, error) {
	str, base := intBase(str)

	result, err := strconv.ParseInt(str, base, bits32)
	if err != nil {
		err = makeParseErr(ErrInvalidInt32, err, name, str)
	}

	return int32(result), err
}

func parseInt16(name, str string) (int16, error) {
	str, base := intBase(str)

	result, err := strconv.ParseInt(str, base, bits16)
	if err != nil {
		err = makeParseErr(ErrInvalidInt16, err, name, str)
	}

	return int16(result), err
}

func parseInt8(name, str string) (int8, error) {
	str, base := intBase(str)

	result, err := strconv.ParseInt(str, base, bits8)
	if err != nil {
		err = makeParseErr(ErrInvalidInt8, err, name, str)
	}

	return int8(result), err
}

func parseInt(name, str string) (int, error) {
	str, base := intBase(str)

	result, err := strconv.ParseInt(str, base, bitsDefault)
	if err != nil {
		err = makeParseErr(ErrInvalidInt, err, name, str)
	}

	return int(result), err
}
