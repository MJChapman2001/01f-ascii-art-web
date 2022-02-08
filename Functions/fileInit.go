package functions

import (
	"io/ioutil"
	"log"
	"strings"
)

func FileInit(banner string) map[int]string {
	file, err := ioutil.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
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

	return art
}