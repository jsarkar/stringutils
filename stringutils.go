package stringutils

import (
	"fmt"
	"strings"
)

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func WordCount(s string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		wc := len(strings.Fields(l))
		fmt.Printf("%d. %s (WC: %d)\n", i, l, wc)
	}
	return s
}
