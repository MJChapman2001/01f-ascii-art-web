package functions

import (
	"errors"
)

func CheckBanner(s string) (map[int]string, error) {
	bannerList := []string{"standard", "shadow", "thinkertoy"}

	for _, i := range bannerList {
		if i == s {
			path := "Banners/" + s + ".txt"

			return FileInit(path), nil
		}
	}

	return nil, errors.New("Banner not found")
}

func CheckOption(s string) (string, error) {
	for i, j := range s {
		if j == '=' {
			return s[2:i], nil
		}
	}

	return "", errors.New("Option not found")
}