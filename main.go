package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	args := os.Args[1:]

	numOfArgs := len(args)

	if numOfArgs != 2 {
		fmt.Println("Error: unexpected number of arguments supplied")
		fmt.Printf("Got %d, expected 2\n", numOfArgs)

		os.Exit(1)
	}

	search_string := args[0]
	search_dir := args[1]

	files, err := os.ReadDir(search_dir)

	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)

		go readAndPrint(&wg, file, search_dir, search_string)
	}

	wg.Wait()
}

func readAndPrint(wg *sync.WaitGroup, file fs.DirEntry, dir, str string) {
	defer wg.Done()

	if file.IsDir() {
		path := filepath.Join(dir, file.Name())

		files, err := os.ReadDir(path)

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			wg.Add(1)

			go readAndPrint(wg, file, path, str)
		}
	} else {
		path := filepath.Join(dir, file.Name())

		opened, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer opened.Close()

		scanner := bufio.NewScanner(opened)

		lineNumber := 0

		for scanner.Scan() {
			lineNumber += 1

			line := scanner.Text()

			if strings.Contains(strings.ToLower(line), strings.ToLower(str)) {
				fmt.Printf("%s:%d: %s\n", path, lineNumber, strings.TrimSpace(line))
			}
		}
	}
}
