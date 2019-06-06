package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func printIt(fullPath string) {
	extValue := strings.ToLower(path.Ext(fullPath))
	if extValue == ".mp3" {
		fmt.Printf("%s is a music file\n", fullPath)
	} else {
		fmt.Printf("file -> %s (extension is %s)\n", fullPath, path.Ext(fullPath))
	}

}

func visit(fullPath string, f os.FileInfo, err error) error {
	fi, err := os.Stat(fullPath)

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		printIt(fullPath)
	}

	return nil
}

func main() {
	root := "."

	flag.Parse()
	flagCount := len(flag.Args())

	if flagCount == 1 {
		root = flag.Arg(0)
	}

	filepath.Walk(root, visit)
}
