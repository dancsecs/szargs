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
)

// Exported errors.
var (
	ErrSyntax         = errors.New("syntax")
	ErrRange          = errors.New("range")
	ErrInvalidFloat64 = errors.New("invalid float64")
	ErrInvalidFloat32 = errors.New("invalid float32")
	ErrInvalidInt64   = errors.New("invalid int64")
	ErrInvalidInt16   = errors.New("invalid int16")
	ErrInvalidInt32   = errors.New("invalid int32")
	ErrInvalidInt8    = errors.New("invalid int8")
	ErrInvalidInt     = errors.New("invalid int")
	ErrInvalidUint64  = errors.New("invalid uint64")
	ErrInvalidUint32  = errors.New("invalid uint32")
	ErrInvalidUint16  = errors.New("invalid uint16")
	ErrInvalidUint8   = errors.New("invalid uint8")
	ErrInvalidUint    = errors.New("invalid uint")
)
