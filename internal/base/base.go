package baseCmd

import (
	"fmt"

	"github.com/fatih/color"
)

func About() {
	color.Yellow("leetcode exercise generator\n")
	fmt.Println("usage: lcgen [-v | --version]\n")
	fmt.Println("These are the list of available commands:\n")
	fmt.Println("\tinit\t\tinitialise a folder for tracking")
	fmt.Println("\tsearch\t\tsearch for your leetcode problem")
	fmt.Println("\tgen\t\tgenerate a leetcode problem")
	fmt.Println("")
}