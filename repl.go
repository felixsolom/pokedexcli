package main

import "strings"

func cleanInput(text string) []string {
	lower_str := strings.ToLower(text)
	sliced_strs := strings.Fields(lower_str)
	return sliced_strs
}
