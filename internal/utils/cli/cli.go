package cli

import "fmt"

var Colors = map[string]string{
	"RESET":  "\033[0m",
	"RED":    "\033[31m",
	"GREEN":  "\033[32m",
	"YELLOW": "\033[33m",
	"BLUE":   "\033[34m",
}

var Prompt = Colors["GREEN"] + "pokedex-cli> " + Colors["RESET"]

func PrintError(errorMessage string) {
	fmt.Print(Colors["RED"] + errorMessage + Colors["RESET"])
}
