package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			fmt.Printf("Testing: %s\n", scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}
	}
}

