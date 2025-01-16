package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex >> ")
		reader.Scan()
		input := reader.Text()
		if len(input) == 0 {
			continue
		}
		if input == "exit" {
			break
		}
		words := cleanInput(input)
		fmt.Println(words)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
