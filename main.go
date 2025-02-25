package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		// Use the scanner's .Scan and .Text methods to get the user's input as a string
		scanner.Scan()
		text := scanner.Text()
		input := strings.ToLower(text)
		words := strings.Fields(input)
		first := words[0]
		fmt.Printf("Your command was: %s\n", first)
	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
