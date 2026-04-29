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
	"strings"
)

func reflowLines(prefix, lines string, width int) string {
	var (
		lastLineWasBlank = true
		reflowedLine     string
		res              strings.Builder
	)

	for line := range strings.SplitSeq(lines, "\n") {
		if line == "" && lastLineWasBlank {
			continue
		}

		lastLineWasBlank = line == ""

		// Line is also indented to ass this to the prefix.
		lineMinusIndent := strings.TrimLeft(line, " ")
		indent := len(line) - len(lineMinusIndent)

		if indent > 0 {
			reflowedLine = reflowLine(
				prefix+strings.Repeat(" ", indent),
				line,
				width,
			)
		} else {
			reflowedLine = reflowLine(prefix, line, width)
		}

		res.WriteString(reflowedLine + "\n")
	}

	return strings.TrimRight(res.String(), "\n")
}

func reflowLine(prefix, line string, width int) string {
	var (
		res     string
		newLine string
	)

	if line == "" {
		return ""
	}

	isFirstLine := true
	prefixWidth := len(prefix)
	contentWidth := width - prefixWidth - 1

	addLine := func(line string) {
		res += prefix + strings.TrimRight(line, " ") + "\n"

		if isFirstLine {
			isFirstLine = false
			prefix = strings.Repeat(" ", prefixWidth)
		}
	}

	for _, wrd := range strings.Split(strings.TrimSpace(line), " ") {
		if len(wrd) >= contentWidth { // Put single long word on own line.
			if newLine != "" {
				addLine(newLine)
				newLine = ""
			}

			addLine(wrd)

			continue
		}

		if len(prefix)+len(newLine)+len(wrd)+1 < width {
			newLine += wrd + " "
		} else {
			addLine(newLine)
			newLine = wrd + " "
		}
	}

	if newLine != "" {
		addLine(newLine)
	}

	return strings.TrimRight(res, "\n")
}
