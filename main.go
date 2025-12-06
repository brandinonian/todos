package main

import (
	"fmt"
	"os"
)

func main() {
	path := "."

	if len(os.Args) > 2 {
		panic("Usage: todos <path>")
	}

	if len(os.Args) == 2 {
		path = os.Args[1]
	}

	err, results := SearchDir(path)
	if err != nil {
		fmt.Print("Error searching directory...\n")
		panic(err)
	}

	if len(results) == 0 {
		fmt.Printf("No todos found!\n")
		os.Exit(0)
	}

	for _, file := range results {
		fmt.Print("-----------------------------------\n")
		fmt.Printf("%s\n", file.name)
		fmt.Print("-----------------------------------\n")
		for _, line := range file.lines {
			fmt.Printf("%s\n", line)
		}
	}
}
