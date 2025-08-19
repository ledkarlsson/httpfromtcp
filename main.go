package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}

	b1 := make([]byte, 8)
	lineToPrint := ""
	for {
		n1, err := f.Read(b1)
		if err != nil {
			if err == io.EOF {
				fmt.Print("read: " + lineToPrint)
				break
			} else {
				panic(err)
			}
		}

		splittedString := strings.Split(string(b1[:n1]), "\n")
		if len(splittedString) == 1 {
			lineToPrint = lineToPrint + splittedString[0]
		} else {
			lineToPrint = lineToPrint + splittedString[0]
			fmt.Printf("read: %s", lineToPrint)
			fmt.Print("\n")
			lineToPrint = splittedString[1]
		}
	}
}
