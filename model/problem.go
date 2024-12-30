package model

type Problem struct {
	Link               string `json:"link"`
	QuestionId         string `json:"questionId"`
	QuestionFrontendId string `json:"questionFrontendId"`
	QuestionTitle      string `json:"questionTitle"`
	TitleSlug          string `json:"titleSlug"`
	Difficulty         string `json:"difficulty"`
	IsPaidOnly         bool   `json:"isPaidOnly"`
	Topics             []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"topicTags"`
	Hints            []string `json:"hints"`
	Likes            int      `json:"likes"`
	Dislikes         int      `json:"dislikes"`
	SimilarQuestions string   `json:"similarQuestions"`
	ExampleTestcases string   `json:"exampleTestcases"`
}