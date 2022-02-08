package functions

import (
	"strings"
)

func Transform(word string, chars map[int]string) string {
	var result string

	var splitNewLine []string

	if len(word) == 0 {
		return ""
	}

	if word == "\\n" {
		return "\n"
	}

	for x, y := 0, 0; x < len(word); x++ {
		if x+1 <= len(word)-1 {
			if word[x] == '\\' && word[x+1] == 'n' {
				splitNewLine = append(splitNewLine, word[y:x])
				y = x+2
			}
		} else if x == len(word)-1 {
			splitNewLine = append(splitNewLine, word[y:])
		}
	}

	for _, z := range splitNewLine {
		var enlarged []string

		for _, i := range z {
			enlarged = append(enlarged, chars[int(i)])
		}

		for j := 0; j < 8; j++ {
			for _, k := range enlarged {
				splitLine := strings.Split(k, "\n")
				result += splitLine[j]
			}
			result += "\n"
		}
	}

	return result
}