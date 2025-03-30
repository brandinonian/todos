package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func SearchDir(path string) (error, []file) {

	err := os.Chdir(path)
	if err != nil {
		panic(err)
	}

	var files []file
	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		if info.Name() == "node_modules" {
			return filepath.SkipDir
		}

		if info.Name() == ".git" {
			return filepath.SkipDir
		}

		if info.IsDir() == false {
			err, has_todo, lines := CheckFile(path, info)
			if err != nil {
				return err
			}

			if has_todo {
				files = append(files, file{info.Name(), lines})
			}
		}

		return nil
	})

	if err != nil {
		return err, nil
	}

	return nil, files
}

func CheckFile(path string, info fs.FileInfo) (error, bool, []string) {
	has_todo := false
	lines := []string{}
	file, err := os.Open(path)
	if err != nil {
		return err, has_todo, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	next := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		if strings.Contains(lineText, "TODO:") {
			has_todo = true
			next = 10
		}

		if next > 0 {
			lines = append(lines, lineText)
			next--
		}
	}

	if err := scanner.Err(); err != nil {
		return err, has_todo, nil
	}

	return nil, has_todo, lines
}
