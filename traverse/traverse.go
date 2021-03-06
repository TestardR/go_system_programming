package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	Path := arguments[1]
	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func walkFunction(path string, info os.FileInfo, err error) error {
	_, err = os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(path)
	return nil
}
