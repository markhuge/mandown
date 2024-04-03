---
title: "mandown"
section: "1"
date: "2024-04-03"
version: "1.0"
name: "mandown"
---

## NAME

**mandown** - a utility for processing data in various formats.

## SYNOPSIS

**mandown** [**-t** title] [**-s** section] [**-d** date] [**-v** version] [**-n** name] FILE

**mandown** [**-t** title] [**-s** section] [**-d** date] [**-v** version] [**-n** name] < FILE

## DESCRIPTION

**mandown** is a utility for converting markdown files into manpages. 

## FRONTMATTER

**mandown** optionally supports markdown frontmatter for setting header values.

If present, **Frontmatter MUST appear at the top of the markdown file** in order to be parsed.

The frontmatter fields map onto the OPTIONS below. OPTIONS override frontmatter.

*EXAMPLE*

```
---
title: "mandown - markdown to man page"
section: "1"
date: "2024-04-01"
version: "1.0"
name: "mandown"
---
```

## OPTIONS

**NOTE: Options override frontmatter.**

 **-t**, title
  
  Sets the title of the document.

 **-s**, section
  
  Specifies the section number of the manual.

 **-d**, date
  
  Sets the date of the document.

 **-v**, version
  
  Specifies the version of the document.

 **-n**, name
  
  Sets the name of the manual.

## EXAMPLES

Convert a Markdown file to troff format, setting the title and section:

```sh
mandown -t "A New Title" -s 1 input.md
```

## LICENSE

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
