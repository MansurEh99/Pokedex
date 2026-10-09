package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands = map[string]cliCommand{
		"help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
		"exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
	}
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			command, ok := commands[input]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			command.callback()
		} else {
			break
		}
	}
}
