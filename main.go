package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
    args := os.Args[1:]
	key := getKey(args)

	// debug
	// fmt.Println("key:", key)
	// fmt.Println("args:", args)

	switch key {
	case "":
		color.Yellow("leetcode exercise generator\n")
		fmt.Println("usage: lcgen [-v | --version]\n")
		fmt.Println("These are the list of available commands:\n")
		fmt.Println("\tinit\t\tinitialise a leetcode folder for tracking")
		fmt.Println("\tsearch\t\tsearch for your leetcode problem")
		fmt.Println("\tgen\t\tgenerate a leetcode problem")
		fmt.Println("")
	case "init":
		color.Yellow("This does nothing at the moment.")
		// function here
	case "search":
		color.Yellow("This does nothing at the moment.")
		// function here
	case "gen":
		color.Yellow("generating...")
		// function here
	default:
		color.Red("command not found.")
	}

}

func getKey(args []string) string {
	// check if args is empty
	if len(args) == 0 {
		return ""
	}

	// check if the first arg is a flag
	char := string([]rune(args[0])[0])	
	if char == "-" {
		return ""
	}

	return args[0]
}