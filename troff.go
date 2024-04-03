/*
Copyright (C) 2024 Mark Wilkerson

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
package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Markdown regex
var (
	LINK      = regexp.MustCompile(`\[([^\[]+)\]\(([^)]+)\)`)
	BOLD      = regexp.MustCompile(`\*\*([^\*]+)\*\*`)
	ITALIC    = regexp.MustCompile(`\*([^\*]+)\*`)
	CODEBLOCK = regexp.MustCompile("^```")
)

func escapeTroff(text string) string {
	return strings.ReplaceAll(text, "\\", "\\e")
}

func MDtoTroff(r io.Reader, fm *FrontMatter) string {
	var (
		troff       strings.Builder
		inCodeBlock bool
	)

	scanner := bufio.NewScanner(r)

	// Write header
	fm.Parse(scanner)
	troff.WriteString(fmt.Sprintf("%s\n", fm))

	for scanner.Scan() {
		line := scanner.Text()

		// Code blocks
		if inCodeBlock {
			if CODEBLOCK.MatchString(line) {
				troff.WriteString(".EE\n")
				inCodeBlock = false
				continue
			}
			troff.WriteString(escapeTroff(line) + "\n")
			continue
		}

		if CODEBLOCK.MatchString(line) {
			troff.WriteString(".EX\n")
			inCodeBlock = true
			continue
		}

		// Formatting
		line = LINK.ReplaceAllString(line, "$1 (URL: $2)")
		line = BOLD.ReplaceAllString(line, "\\fB$1\\fP")
		line = ITALIC.ReplaceAllString(line, "\\fI$1\\fP")

		// Sections
		switch {
		case strings.HasPrefix(line, "# "):
			troff.WriteString(fmt.Sprintf(".SH %s\n", strings.TrimPrefix(line, "# ")))
		case strings.HasPrefix(line, "## "):
			troff.WriteString(fmt.Sprintf(".SS %s\n", strings.TrimPrefix(line, "## ")))
		case strings.HasPrefix(line, "* "):
			troff.WriteString(fmt.Sprintf(".IP \\(bu 4\n%s\n", strings.TrimPrefix(line, "* ")))
		case line == "":
			troff.WriteString(".PP\n")
		default:
			troff.WriteString(line + "\n")
		}
	}

	return troff.String()
}
