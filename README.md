ascii-art
Description

ascii-art is a command-line program written in Go that converts a given string into a graphical representation using ASCII characters.
Each character of the input string is rendered as a block of ASCII symbols, producing a large, readable banner-style output directly in the terminal.

The project focuses on string processing, file handling, and algorithmic formatting, using only the Go standard library.

An optional extension adds color support, allowing users to highlight parts of the ASCII-art output.

Features

Converts text into ASCII-art banners

Handles:

Letters

Numbers

Spaces

Special characters

Newlines (\n)

Supports multiple banner styles:

standard

shadow

thinkertoy

Optional color highlighting using command-line flags

Clean and predictable terminal output

How It Works

The program receives a string as a command-line argument

Each character is mapped to its ASCII-art representation using banner files

Each character is rendered with a fixed height of 8 lines

Characters are combined line by line to form the final output

The result is printed to standard output

Banner Format

Each character consists of 8 lines

Characters are separated by a newline in the banner files

Banner files are provided and must not be modified

Example representation (dots represent spaces):

......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....

Usage
Basic usage
go run . "Hello"

With newlines
go run . "Hello\nThere"

With empty input
go run . ""

Color Extension (Optional)

The program supports coloring ASCII-art output using the following flag format:

--color=<color> <substring> "string"

Rules

<color> is chosen by the user (ANSI, RGB, etc. — implementation-specific)

<substring> can be:

A single character

A word

Multiple characters

If <substring> is omitted, the entire output is colored

If the flag format is invalid, the program returns a usage message

Example
go run . --color=red kit "a king kitten have kit"


In this example:

The substring kit inside kitten

And the word kit
are both colored in the output.

Invalid usage
Usage: go run . [OPTION] [STRING]

Supported Arguments

[STRING] — text to convert into ASCII-art

[BANNER] — optional banner selection

[OPTION] — optional flags such as --color

The program must still work correctly when only a string is provided.

Project Structure
.
├── main.go
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── go.mod
└── README.md

Constraints

Written entirely in Go

Uses only standard Go packages

Banner files must not be modified

Output must correctly handle all valid ASCII characters

Learning Outcomes

This project helps develop understanding of:

ASCII encoding and representation

String and data manipulation

File system (fs) API in Go

Algorithmic text rendering

Terminal output formatting

Color handling in CLI applications
