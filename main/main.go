package main

import (
	"fmt"
	"snippets/cmd"
)

func main() {

	path, err := cmd.DefaultFilePath()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
}
