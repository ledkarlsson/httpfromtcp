package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}

	b1 := make([]byte, 8)

	for {
		n1, err := f.Read(b1)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		fmt.Printf("read: %s\n", b1[:n1])
	}

}
