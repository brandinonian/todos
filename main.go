package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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
		panic(err)
	}

	if len(results) > 0 {
		fmt.Println("TODO items:")

		for i := 0; i < len(results); i++ {
			fmt.Printf("%s\n", results[i])
		}
		return
	}

	fmt.Println("No TODO items found")
}

func SearchDir(path string) (error, []string) {

	err := os.Chdir(path)
	if err != nil {
		panic(err)
	}

	var files []string
	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if info.Name() == "node_modules" {
			return filepath.SkipDir
		}

		if info.Name() == ".git" {
			return filepath.SkipDir
		}

		if info.IsDir() == false {
			err, hasTodo := CheckFile(path, info)
			if err != nil {
				return err
			}

			if hasTodo {
				files = append(files, info.Name())
			}
		}

		return nil
	})

	if err != nil {
		return err, nil
	}

	return nil, files
}

func CheckFile(path string, info fs.FileInfo) (error, bool) {
	file, err := os.Open(path)
	if err != nil {
		return err, false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		if strings.Contains(lineText, "TODO") {
			return nil, true
		}
	}

	if err := scanner.Err(); err != nil {
		return err, false
	}

	return nil, false
}
