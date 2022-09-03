package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	numOfArgs := len(args)

	if numOfArgs == 0 || numOfArgs > 2 {
		fmt.Println("Error: insufficient arguments supplied")
		fmt.Printf("Got %d, expected 2\n", numOfArgs)

		os.Exit(1)
	}

	// search_string := args[0]
	var search_dir = "."

	if numOfArgs == 2 {
		search_dir = args[1]
	}

	files, err := os.ReadDir(search_dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory - ", file.Name())
		} else {
			opened, err := os.Open(file.Name())

			if err != nil {
				log.Fatal(err)
			}

			scanner := bufio.NewScanner(opened)

			for scanner.Scan() {
				line := scanner.Text()

				fmt.Println(line)
			}
		}
	}

}
