package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		io.Copy(os.Stdout, file)
	}
}
