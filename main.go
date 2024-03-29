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

	guide, err := pokeapi.Versus(attacker, defender)

	if err != nil {
		return errors.New("something wrong")
	}

	fmt.Println("You can use these 2xDmg attacks")
	for _, move := range guide.MovesAgainstDefender {
		fmt.Printf("  -%s\n", move.Name)
	}

	fmt.Println("You are vulnerable to these 2xDmg attacks")
	for _, move := range guide.MovesAgainstAttacker {
		fmt.Printf("  -%s\n", move.Name)
	}

	return nil
}
func infoCommand(args []string) error {
	if len(args) != 1 {
		return errors.New("invalid number of arguments")
	}

	pokemon, err := pokeapi.GetPokemon(args[0])

	if err != nil {
		return err
	}

	padding := strings.Repeat(" ", 4)

	fmt.Println(pokemon.Name)
	fmt.Printf("%sMoves\n", padding)
	for index, move := range pokemon.Moves {
		fmt.Printf("%s%s+%s%s", padding, padding, move.Name, padding)
		if (index+1)%3 == 0 {
			fmt.Println()
		}

		if (index+1)%3 != 0 && index == len(pokemon.Moves)-1 {
			fmt.Println()
		}
	}

	fmt.Printf("%sTypes\n", padding)
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("%s%s+%s%s\n", padding, padding, pokemonType.Name, padding)
	}

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
