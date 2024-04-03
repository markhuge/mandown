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
	"flag"
	"fmt"
	"io"
	"os"
)

func isPipe() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func parseArgs() (io.Reader, *FrontMatter, error) {
	var title, section, date, version, name string

	flag.StringVar(&title, "t", "", "Title of the manual page")
	flag.StringVar(&section, "s", "", "Section number of the manual page")
	flag.StringVar(&date, "d", "", "Date of the manual page")
	flag.StringVar(&version, "v", "", "Version of the manual page")
	flag.StringVar(&name, "n", "", "Name of the command")

	flag.Parse()

	fm := &FrontMatter{
		Title:   title,
		Section: section,
		Date:    date,
		Version: version,
		Name:    name,
	}

	args := flag.Args()
	var input io.Reader

	if len(args) > 0 {
		file, err := os.Open(args[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error opening input file: %w", err)
		}
		input = file
	} else if isPipe() {
		input = os.Stdin
	} else {
		return nil, nil, fmt.Errorf("No input file specified")
	}

	return input, fm, nil
}
