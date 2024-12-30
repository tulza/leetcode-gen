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

	var topics []string
	for _, topic := range problem.Topics {
		topics = append(topics, topic.Slug)
	}

	return []byte(
		fmt.Sprintf(`
#problem: %s
#topics: %s

`, problem.Link, strings.Join(topics, ", ")))
}
