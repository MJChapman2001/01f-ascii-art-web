package main

import (
	"fmt"
	"os"
	//"strings"

	"ascii-art/Functions"
)

func main() {
	args := os.Args[1:]

	 // ascii-art
	if len(args) == 1 {
		word := args[0]
		chars := functions.FileInit("Banners/standard.txt")

		result := functions.Transform(word, chars)

		fmt.Print(result)
	}
	

	/* // fs
	message := "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"

	if len(args) == 2 {
		word := args[0]
		banner := args[1]

		chars, err := CheckBanner(banner)
		if err != nil {
			fmt.Println(message)
			panic(err)
		}

		result := functions.Transform(word, chars)

		fmt.Print(result)
	}
	*/

	/* // options
	if len(args) == 3 {
		word := args[0]
		banner := args[1]
		optionInput := args[2]

		chars, err := functions.CheckBanner(banner)
		if err != nil {
			panic(err)
		}
		
		message := "Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard --output=<fileName.txt>\nEX: go run . something standard --align=right"

		userOption, err := functions.CheckOption(optionInput)
		if err != nil {
			fmt.Println(message)
			panic(err)
		}

		switch userOption {
		case "output":
			fileName := strings.TrimPrefix(optionInput, "--output=")
			
			result := functions.Transform(word, chars)

			os.WriteFile(fileName, []byte(result), 0666)
		case "align":
			alignType := strings.TrimPrefix(optionInput, "--align=")
			result := functions.Transform(word, chars)

			functions.AlignText(result, alignType)
		case "color":
		case "reverse":
		default:
			fmt.Println(message)
		}
	}
	*/
}