package checker

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

const HashOfStandard = "c3ec7584fb7ecfbd739e6b3f6f63fd1fe557d2ae3e24f870730d9cf8b2559e94"

func IsValidStandard(str string) bool {
	hasher := sha256.New()
	s, err := os.ReadFile(str)
	if err != nil {
		log.Fatal(err)
	}
	hasher.Write(s)
	hashDeclaredMap := map[string]bool{HashOfStandard: true}
	hashCalculated := hex.EncodeToString(hasher.Sum(nil))
	return hashDeclaredMap[hashCalculated]
}

func ASCII_Check(input string) error {
	for i, r := range input {
		if r == '\n' {
			continue
		}
		if r < 32 || r > 126 {
			fmt.Fprintf(os.Stderr, "Error: Invalid character at position %d: %q\n", i+1, r)
			os.Exit(1)
		}
	}
	return nil
}
