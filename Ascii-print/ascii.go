package AsciiPrint

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const bannerHeight = 8

func BannerLoader(path string) (map[rune][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	asciiMap := make(map[rune][]string)

	scanner.Scan()

	char := rune(32)

	for scanner.Scan() {
		var lines []string
		lines = append(lines, scanner.Text())
		for i := 1; i < bannerHeight; i++ {
			scanner.Scan()
			lines = append(lines, scanner.Text())
		}
		asciiMap[char] = lines
		char++
		scanner.Scan()
	}
	return asciiMap, nil
}

var TerminalColors = map[string]string{
	"reset":   "\033[0m",
	"bold":    "\033[1m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",

	// Яркие версии
	"brightRed":     "\033[91m",
	"brightGreen":   "\033[92m",
	"brightYellow":  "\033[93m",
	"brightBlue":    "\033[94m",
	"brightMagenta": "\033[95m",
	"brightCyan":    "\033[96m",
	"brightWhite":   "\033[97m",
}

func PrintAscii(text, color string, banner map[rune][]string) {
	for i := 0; i < bannerHeight; i++ {
		for _, ch := range text {
			if val, ok := banner[ch]; ok {
				fmt.Print(ColorPrint(color, val[i]))
			} else {
				fmt.Print(strings.Repeat(" ", len(banner[' '][i])))
			}
		}
		fmt.Println()
	}
}

func PrintAsciiColor(str, color, sub string, banner map[rune][]string) {

	// Step 1: Track which positions should be highlighted
	highlight := make([]bool, len(str))
	for i := 0; i <= len(str)-len(sub); i++ {
		if str[i:i+len(sub)] == sub {
			for j := 0; j < len(sub); j++ {
				highlight[i+j] = true
			}
			i += len(sub) - 1
		}
	}

	// Step 2: Print each ASCII art row
	for row := 0; row < bannerHeight; row++ {
		for pos, ch := range str {
			if val, ok := banner[ch]; ok && row < len(val) {
				line := val[row]
				if highlight[pos] {
					fmt.Print(ColorPrint(color, line))
				} else {
					fmt.Print(line)
				}
			} else {
				// Unknown character: print space of correct width
				fmt.Print(strings.Repeat(" ", len(banner[' '][row])))
			}
		}
		fmt.Println()
	}
}

func ColorPrint(color, str string) string {
	code, ok := TerminalColors[color]
	if !ok {
		code = TerminalColors["reset"]
	}
	return code + str + TerminalColors["reset"]
}
