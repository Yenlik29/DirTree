package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Path struct {
	Path  string
	Mem   int
	Files []string
}

func PathParse(path string) (string, error) {
	switch dot := path; dot {
	case ".":
		dir, err := filepath.Abs(filepath.Dir("."))
		if err != nil {
			return path, fmt.Errorf("Error")
		}
		path = "" + dir /* path = os.Getenv("PWD")*/
	case "..":
		path = os.Getenv("OLDPWD")
	default:
		break
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("[dirTree:] Dir/file doesn't exist")
		return path, fmt.Errorf("[dirTree:] Dir/file doesn't exist")
	}
	return path, nil
}

func PermissionParse(path string) error {
	if _, err := os.OpenFile(path, os.O_RDWR, 0600); os.IsPermission(err) {
		fmt.Println("[dirTree]: Permission denied")
		return fmt.Errorf("[dirTree]: Permission denied")
	}
	return nil
}

func expandedTree(path string) {
	path = ""
	return
}

func building(FullPath Path) {
	files, err := ioutil.ReadDir(FullPath.Path)
	if err != nil {
		fmt.Println("[ReadDir]: Error")
		return
	}
	for _, file := range files {
		if !strings.Contains(file.Name(), ".") {
			fmt.Println("└───", file.Name())
		}
	}
}

func simpleTree(path string) {
	FullPath := Path{Path: path}
	dir, err := os.Open(FullPath.Path)
	if err != nil {
		fmt.Println("[Open]: Error")
		return
	}
	di, _ := dir.Stat()
	if di.IsDir() {
		fmt.Println("Cool, it's directory")
		building(FullPath)
	} else {
		fmt.Println("[Warning]: It's a file")
		return
	}
	return
}

func buildTree(path string, printFiles int) {
	if printFiles == 1 {
		expandedTree(path)
	} else {
		simpleTree(path)
	}
	return
}

func dirTree(path string, printFiles int) (string, error) {
	path, _ = PathParse(path)
	if path, ok := PathParse(path); ok != nil {
		return path, fmt.Errorf("Parse error")
	}
	if ok := PermissionParse(path); ok != nil {
		return path, fmt.Errorf("Parse error")
	}
	buildTree(path, printFiles)
	return path, nil
}

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("Usage: add/remove args")
		return
	}
	path := os.Args[1]
	printFiles := 0
	if len(os.Args) == 3 && os.Args[2] == "-f" {
		printFiles = 1
	}
	if _, err := dirTree(path, printFiles); err != nil {
		return
	}
}
