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
	"strings"

	"golang.org/x/term"
)

const (
	usagePrefix      = "usage: "
	defaultLineWidth = 78
)

// Prepare desc appends paragraphs (separated by blank lines) to single lines
// in preparation for reflowing to different widths.
func prepareDesc(prefix, desc string) string {
	var res strings.Builder

	lines := strings.Split(strings.Trim(desc, "\n"), "\n")
	first := 0

	for i, l := range lines {
		if l == "" {
			res.WriteString(
				prefix +
					strings.Join(lines[first:i], " ") +
					"\n\n",
			)

			first = i + 1
		}
	}

	if first < len(lines) {
		res.WriteString(prefix + strings.Join(lines[first:], " "))
	}

	return strings.TrimRight(res.String(), "\n")
}

const spacePlaceholder = "\u001f"

// RegisterUsage registers a new flag and its description if and only if the
// flag has not been already  registered.
func (args *Args) RegisterUsage(item, desc string) {
	if !args.usageDefined[item] {
		args.usageHeader += " " +
			strings.ReplaceAll(item, " ", spacePlaceholder)
		args.usageBody += "\n" + item + "\n" +
			prepareDesc("    ", desc) + "\n"
		args.usageDefined[item] = true
	}
}

func terminalWidth(out *os.File) int {
	fd := int(out.Fd())

	if term.IsTerminal(fd) {
		termWidth, _, err := term.GetSize(fd)
		if err != nil || termWidth == 0 {
			termWidth = defaultLineWidth
		}

		return termWidth
	}

	return defaultLineWidth
}

func (args *Args) buildSynopsisHeader(synopsisVersions []string) string {
	var header string

	for i, synopsis := range synopsisVersions {
		if i == 0 {
			header += "\n" + usagePrefix
		} else {
			header += "\n" + strings.Repeat(" ", len(usagePrefix))
		}

		header += args.programName + " " + synopsis
	}

	return header
}

// Usage returns a usage messages representing the Args object.  It is
// formatted to the lineWidth provided.  A zero uses the defaultLineWidth
// while a negative value caused an effort to determine if writing to a
// terminal and if so using its width otherwise defaulting.
func (args *Args) Usage(lineWidth int) string {
	var header string

	if lineWidth < 0 {
		lineWidth = terminalWidth(os.Stdout)
	}

	if lineWidth < 1 {
		lineWidth = defaultLineWidth
	}

	args.lineWidth = lineWidth

	if len(args.usageSynopsis) > 0 {
		header = args.buildSynopsisHeader(args.usageSynopsis)
	} else {
		if args.usageHeader == "" {
			header = "\nusage: " + args.programName
		} else {
			header = "\n" + reflowLine("usage: "+args.programName+" ",
				args.usageHeader,
				args.lineWidth,
			)
			header = strings.ReplaceAll(header, spacePlaceholder, " ")
		}
	}

	return strings.TrimRight(
		header[1:]+
			"\n\n"+
			reflowLine("", args.programDesc, args.lineWidth)+"\n\n"+
			reflowLines("    ", args.usageBody, args.lineWidth)+"\n",
		"\n",
	)
}

// AddSynopsis includes another static synopsis message.
func (args *Args) AddSynopsis(s string) {
	args.usageSynopsis = append(args.usageSynopsis, s)
}
