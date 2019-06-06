// fileio
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func walkDirectory(dirPath string) {
	var wg sync.WaitGroup

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		wg.Add(1)
		go execSearch(dirPath+"/"+f.Name(), &wg)
	}

	wg.Wait()
}

func execSearch(filePath string, wg *sync.WaitGroup) {
	if strings.HasSuffix(filePath, ".log") {
		f, err := os.Open(filePath)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		r := bufio.NewReaderSize(f, 4*4096)

		line, isPrefix, err := r.ReadLine()

		for err == nil {
			s := string(line)
			if strings.Contains(s, "text_to_search_for") {
				fmt.Println(s)
			}
			line, isPrefix, err = r.ReadLine()
		}

		if isPrefix {
			fmt.Println("buffer size too small")
			return
		}

		if err != io.EOF {
			fmt.Println(err)
			os.Exit(-1)
		}
	}

	wg.Done()
}

func main() {
	walkDirectory("root_path")
}
