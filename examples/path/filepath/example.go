package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files, err := getFiles(".", ".go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(files))
}

// listFiles returns files under specific root
func listFiles(path, ext string) ([]string, error) {

	// place where all filtered files stored.
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error occured when accessing path %q: %v\n", path, err)
			return err
		}

		// filter all files with specific ext.
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// getFiles returns files under specific root
func getFiles(path, ext string) ([]string, error) {
	// store for place all filterd files
	var files []string

	// find all files with specific ext.
	err := filepath.Walk(path, visit(&files, ext))
	if err != nil {
		return nil, err
	}
	return files, nil
}

func visit(files *[]string, ext string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error occured during processing path %q: %v\n", path, err)
		}

		// filter all files with specific ext.
		if filepath.Ext(path) == ext {
			*files = append(*files, path)
		}
		return nil
	}
}
