package functions

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	//"strings"
)

func AlignText(str, alignment string) {
	switch alignment {
	case "center":
		os.WriteFile("text.txt", []byte(str), 0666)

		cmd := exec.Command("/bin/sh", "center.sh")
		output, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(output)
	case "left":
	case "right":
		
	case "justify":
	}
}