package gencmd

import (
	"fmt"
	models "lcgen/model"
	"os"
	"strings"
)

func generateFile(problem models.Problem) (string, error) {
    id := problem.QuestionId
    title := problem.TitleSlug

    fileName := fmt.Sprintf("%s-%s.py", id, title)

    generatedContent := generateContent(problem)

    // Check if file already exists
    if _, err := os.Stat(fileName); err == nil {
        return "", fmt.Errorf("%s already exists", fileName)
    }
    err := os.WriteFile(fileName, generatedContent, 0755)
    if err != nil {
        return "", fmt.Errorf("error generating file: %w", err)
    }


    return fileName, nil
}

func generateContent(problem models.Problem) []byte {

	var title string
	var difficulty string
	var topics []string
	var frontendId string
	var problemLink string

	// get title
	title = problem.QuestionTitle

	// get difficulty
	difficulty = problem.Difficulty

	// get topics
	for _, topic := range problem.Topics {
		topics = append(topics, topic.Slug)
	}

	// get frontend id
	frontendId = problem.QuestionFrontendId

	// get link
	problemLink = problem.Link
	if problemLink == "" {
		problemLink = problem.QuestionLink
	}

	return []byte(
		fmt.Sprintf(`
# %s
# %s | %s | #%s
# %s 

class Solution:
    pass
`, 
title,
difficulty,
strings.Join(topics, ", "),
frontendId,
problemLink, 
))
}
