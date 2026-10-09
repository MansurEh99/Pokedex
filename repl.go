package main

import ("strings")

func cleanInput(text string) []string {
	lowrcaseText := strings.ToLower(text)
	words := strings.Fields(lowrcaseText)
	return words
}
