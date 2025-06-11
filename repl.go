package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := fmt.Sprintf(scanner.Text())
		clean_slice := cleanInput(input)
		first_word := clean_slice[0]
		fmt.Printf("Your command was: %s\n", first_word)
	}
}

func cleanInput(text string) []string {
	lower_str := strings.ToLower(text)
	sliced_strs := strings.Fields(lower_str)
	return sliced_strs
}
