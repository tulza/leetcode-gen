package main

import (
	baseCmd "lcgen/internal/base"
	"lcgen/internal/errorMsg"
	gencmd "lcgen/internal/gen"
	"lcgen/internal/utils"
	"os"
)

func main() {
    args := os.Args[1:]
	key := utils.GetKey(args)

	// debug
	// fmt.Println("key:", key)
	// fmt.Println("args:", args)

	switch key {
	case "":
		baseCmd.About()
	case "init":
		errorMsg.NotImplemented()
		// function here
	case "search":
		errorMsg.NotImplemented()
		// function here
	case "gen":
		gencmd.Generate(args)
	default:
		errorMsg.NotFoundCmd()
	}

}
