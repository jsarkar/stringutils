package stringutils

import (
	"fmt"
	"strings"
)

// Upper case string
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Lower case string
func Lower(s string) string {
	return strings.ToLower(s)
}

// Count number of Word
// TODO: Move the print statement out from the function
func WordCount(s string) {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		wc := len(strings.Fields(l))
		fmt.Printf("%d. %s (WC: %d)\n", i, l, wc)
	}

}

// character count
func CharCount(s string) map[string]int {
	m := make(map[string]int, 0)

	for _, char := range s {
		m[string(char)] = m[string(char)] + 1

	}
	return m
}
