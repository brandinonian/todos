package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var Files []string

func main() {
	path := "."

	if len(os.Args) > 2 {
		panic("Usage: todos <path>")
	}

	if len(os.Args) == 2 {
		path = os.Args[1]
	}

	err := os.Chdir(path)
	if err != nil {
		panic(err)
	}

	if len(Files) > 0 {
		fmt.Println("TODO items:")

		for i := 0; i < len(Files); i++ {
			fmt.Printf("%s\n", Files[i])
		}
		return
	}

	fmt.Println("No TODO items found")
}

func SearchDir() error {
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// no need to waste resources
		if info.Name() == "node_modules" {
			return filepath.SkipDir
		}

		if info.Name() == ".git" {
			return filepath.SkipDir
		}

		if info.IsDir() == false {
			err = CheckFile(path, info)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func CheckFile(path string, info fs.FileInfo) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		if strings.Contains(lineText, "TODO") {
			Files = append(Files, info.Name())
			return nil
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
