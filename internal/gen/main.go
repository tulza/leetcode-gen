package gencmd

import (
	gencmd_flag "lcgen/internal/gen/flag"
	"lcgen/internal/utils"
	"lcgen/model"

	"github.com/fatih/color"
)

func Generate(args []string) {
	// check for correct number of args
	if len(args) < 2 {
		color.Red("Invalid number of arguments. Please provide at least 2 arguments.")
		return
	}

	flag := utils.IsFlag(args[1])
	var isDaily bool = false
	var slugName string

	// debug
	// fmt.Println("flag:", flag)
	// fmt.Println(args, len(args))

	switch flag {
	// no flag
	case "":
		if len(args) != 2 {
			color.Red("Invalid number of arguments. Please exactly 2 arguments.")
			color.Red("lcgen gen [URL]")
			return
		}

		// check for valid LeetCode URL
		if !utils.HasLeetCodeURL(args[1]) {
			color.Red("Invalid LeetCode URL. Please provide a valid LeetCode problem URL.")
			return
		}

		color.Blue("Fetching with URL...")
		name, err := gencmd_flag.GetSlugFromURL(args[1])

		if err != nil {
			color.Red(err.Error())
			return
		}

		slugName = name

	case "-n", "--name":
		if len(args) != 3 {
			color.Red("Invalid number of arguments. Please exactly 3 arguments.")
			color.Red("lcgen gen -n [NAME]")
			return
		}
		color.Blue("Fetching with name...")

		slugName = args[2]

	case "-d", "--daily":
		if len(args) != 2 {
			color.Red("Invalid number of arguments. Please exactly 2 arguments.")
			color.Red("lcgen gen -d")
			return
		}
		color.Blue("Fetching daily...")
		isDaily = true

	default:
		color.Red("Unknown flag.")
		return
	}

	var problem model.Problem

	if isDaily {
		p, err := fetchDaily()
		if err != nil {
			color.Red(err.Error())
			return
		}
		problem = p
	} else {
		p, err := fetchProblem(slugName)
		if err != nil {
			color.Red(err.Error())
			return
		}
		problem = p
	}

	fileName, err := generateFile(problem)
	if err != nil {
		color.Red(err.Error())
		return
	}

	color.Yellow("Generated file: %s", fileName)
}
