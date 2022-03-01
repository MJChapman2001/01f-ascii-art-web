package functions

import (
	"io/ioutil"
	"strings"
)

func FileInit(banner string) (map[int]string, error) {
	file, err := ioutil.ReadFile(banner)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(file), "\n")

	art := make(map[int]string)

	for i, j := 32, 0; i < 127; i++ {
		k := j + 9
		temp := ""

		for j < k {
			temp += lines[j+1] + "\n"
			j++
		}

		art[i] = temp
	}

	return art, nil
}