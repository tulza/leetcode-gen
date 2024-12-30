package gencmd

import (
	"fmt"
	models "lcgen/model"
)

func printProblemObj(problem models.Problem) {
	fmt.Println("Link:", problem.Link)
	fmt.Println("QuestionId:", problem.QuestionId)
	fmt.Println("QuestionFrontendId:", problem.QuestionFrontendId)
	fmt.Println("QuestionTitle:", problem.QuestionTitle)
	fmt.Println("TitleSlug:", problem.TitleSlug)
	fmt.Println("Difficulty:", problem.Difficulty)
	fmt.Println("IsPaidOnly:", problem.IsPaidOnly)
	fmt.Println("Topics:", problem.Topics[0].Name)
	fmt.Println("Hints:", problem.Hints)
}