package functions

import (
	"strings"
)

func Transform(word string, chars map[int]string) string {
	var result string

	if len(word) == 0 {
		return ""
	}

	if word == "\\n" {
		return "\n"
	}

	var enlarged []string

	for _, i := range word {
		enlarged = append(enlarged, chars[int(i)])
	}

	for j := 0; j < 8; j++ {
		for _, k := range enlarged {
			splitLine := strings.Split(k, "\n")
			result += splitLine[j]
		}
		result += "\n"
	}

	return result
}