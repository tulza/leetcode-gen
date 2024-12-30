package gencmd

import (
	"encoding/json"
	"io"
	"lcgen/internal/config"
	"lcgen/internal/utils"
	models "lcgen/model"
	"net/http"

	"github.com/fatih/color"
)

func Generate(args []string) {
    // check for correct number of args
    if len(args) != 2 {
        color.Red("Invalid number of arguments. Please provide a valid LeetCode problem number.")
        return
    }

    // check for valid LeetCode URL
    if (!utils.HasLeetCodeURL(args[1])) {
        color.Red("Invalid LeetCode URL. Please provide a valid LeetCode problem URL.")
        return
    }

    color.Blue("Fetching exercise...")

    // fetch problem title from route 
    _, routes, _ := utils.ParseURL(args[1])
    problemName := routes[1]

    // fetch leetcode data from api
    API_URL := config.API_BASE_URL + `/select?titleSlug=` + problemName
    res, err := http.Get(API_URL)

    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

    // parse data
    var problem models.Problem 
    err = json.Unmarshal(body, &problem)

    if err != nil {
        panic(err)
    }

    // debug
    // printProblem(problem)

    fileName, err := generateFile(problem)
    if err != nil {
        color.Red(err.Error())
        return
    }

    color.Yellow("Generated file: %s", fileName)
}
