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
	"strings"
)

type FrontMatter struct {
	Title   string
	Section string
	Date    string
	Version string
	Name    string
}

func (f *FrontMatter) Parse(scanner *bufio.Scanner) {
	var start bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			if start {
				return // End of frontmatter and successfully parsed
			}
			start = true // Start of frontmatter
			continue
		}
		if start {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.Trim(strings.TrimSpace(parts[1]), "\"")
				switch key {
				case "title":
					if f.Title == "" {
						f.Title = value
					}
				case "section":
					if f.Section == "" {
						f.Section = value
					}
				case "date":
					if f.Date == "" {
						f.Date = value
					}
				case "version":
					if f.Version == "" {
						f.Version = value
					}
				case "name":
					if f.Name == "" {
						f.Name = value
					}

				}
			}
		}
	}
}

func (f FrontMatter) String() string {
	return fmt.Sprintf(".TH \"%s\" %s \"%s\" \"%s\" \"%s\"", f.Name, f.Section, f.Date, f.Version, f.Title)
}
