package main

import (
	"fmt"

	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

type Problem struct {
	Difficulty string
	QuestionId string
	Title      string
	TitleSlug  string
	Content    string
	TopicTag   []string
}

func GetProblems() {
	noOfQuestions := 3
	var problems []Problem
	allProblemList, _ := leetcodeapi.GetAllProblems(0, noOfQuestions)
	for _, value := range allProblemList.Problems {
		var topicTags []string
		titleSlug := value.TitleSlug
		content, err := leetcodeapi.GetProblemContentByTitleSlug(titleSlug)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, value := range value.TopicTags {
			topicTags = append(topicTags, value.Name)
		}
		problem := Problem{
			Title:      value.Title,
			TitleSlug:  value.TitleSlug,
			Difficulty: value.Difficulty,
			QuestionId: value.QuestionId,
			TopicTag:   topicTags,
			Content:    content.Content,
		}
		problems = append(problems, problem)
	}
	fmt.Printf("%+v\n", problems)
}

func main() {
	GetProblems()
}
