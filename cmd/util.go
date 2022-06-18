package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func DirExist(pth string) bool {
	if _, err := os.Stat(pth); os.IsNotExist(err) {
		return false
	}
	return true
}

// create path creates a directory, nested if multiple new paths are defined
func CreatePathAll(pth string) error {
	err := os.MkdirAll(pth, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// writeLines writes the lines to the given file.
func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func unique(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func uniqueInt(s []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
