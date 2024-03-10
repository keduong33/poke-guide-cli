package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"pokedex-cli/internal/pokeapi"
	"pokedex-cli/internal/utils/cli"
)

type cliCommand struct {
	name        string
	description string
	callback    func(args []string) error
}

func helpCommand(args []string) error {
	fmt.Println("How to use my pokedex-cli")
	return nil
}

func exitCommand(args []string) error {
	fmt.Println("Bye bye")
	os.Exit(0)
	return nil
}

func clearCommand(args []string) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

func versusCommand(args []string) error {
	if len(args) != 2 {
		return errors.New("invalid number of arguments")
	}

	attacker, defender := args[0], args[1]

	err := pokeapi.Versus(attacker, defender)

	if err != nil {
		return errors.New("something wrong")
	}
	return nil
}
func infoCommand(args []string) error {
	if len(args) != 1 {
		return errors.New("invalid number of arguments")
	}

	pokeapi.GetPokemonMoves(args[0])
	return nil
}

var commands = map[string]cliCommand{
	"help": {
		name:        "Help",
		description: "How to use",
		callback:    helpCommand,
	},
	"exit": {
		name:        "Exit",
		description: "Exit the program",
		callback:    exitCommand,
	},
	"clear": {
		name:        "Clear",
		description: "Clear the screen",
		callback:    clearCommand,
	},
	"versus": {
		name:        "Versus",
		description: "Comparison between 2 pokemons",
		callback:    versusCommand,
	},
	"info": {
		name:     "Info",
		callback: infoCommand,
	},
}

func introduce() {
	fmt.Printf("%sWelcome to Pokedex-CLI\n", cli.Prompt)
	fmt.Printf("%sWritten by David Duong - 2024\n", cli.Prompt)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	introduce()

	for {
		fmt.Print(cli.Prompt)
		reader.Scan()
		userInputs := strings.Split(reader.Text(), " ")

		if len(userInputs) != 0 && userInputs[0] != "" {
			userCommand := userInputs[0]
			command, ok := commands[userCommand]

			if !ok {
				cli.PrintError(fmt.Sprintf("Command %s found\n", userCommand))
				continue
			}

			err := command.callback(userInputs[1:])

			if err != nil {
				cli.PrintError(fmt.Sprintf("Uh oh: %s\n", err.Error()))
			}
		}
	}

}
