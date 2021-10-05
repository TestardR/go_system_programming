package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}
	filename := os.Args[1]
	b := []byte("Romain Testard\n")
	ioutil.WriteFile(filename, b, 0644)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	bs := make([]byte, 100)

	n, err := f.Read(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read %d bytes: %s", n, bs)
}
