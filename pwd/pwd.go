package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(pwd)
}
