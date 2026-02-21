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

		if scanner.Scan() {
			cmd := cleanInput(scanner.Text())[0]
			fmt.Printf("Your command was: %v\n", cmd)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}
	}
}

func cleanInput(text string) []string {
	//Lowercase the input
	lowered := strings.ToLower(text)

	//Split input into words based on whitespace
	cleaned := strings.Fields(lowered)

	//Return first word
	return cleaned
}
