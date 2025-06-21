package main

import (
	AsciiPrint "ascii/Ascii-print"
	checker "ascii/check"
	"flag"
	"fmt"
	"strings"
)

func main() {

	colorFlag := flag.String("color", "reset", "Color to print text with")
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)

	input := args[0]

	checker.ASCII_Check(input)

	path_to_banner := "banners/standard.txt"
	banner, err := AsciiPrint.BannerLoader(path_to_banner)
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	if len(args) > 1 {
		input = args[1]
		sub := args[0]

		lines := strings.Split(input, "\n")
		for _, line := range lines {
			AsciiPrint.PrintAsciiColor(line, *colorFlag, sub, banner)
		}
	} else {
		lines := strings.Split(input, "\n")
		for _, line := range lines {
			AsciiPrint.PrintAscii(line, *colorFlag, banner)
		}
	}
}
