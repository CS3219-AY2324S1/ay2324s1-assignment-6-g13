package leetcode

import (
	"net/http"

	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

type Problem struct {
	Difficulty string
	QuestionId string
	Title      string
	TitleSlug  string
	Content    string
	Category   []string
}

const LEETCODE_API_FAILURE_MESSAGE = "Failed to Connect to Leetcode API"

func GetAllProblems(offset int, pageSize int) (problems []leetcodeapi.Problem, total int, errorMessage string, httpStatusCode int) {
	response, err := leetcodeapi.GetAllProblems(offset, pageSize)
	if err != nil {
		return nil, 0, LEETCODE_API_FAILURE_MESSAGE, http.StatusInternalServerError
	}
	return response.Problems, response.Total, "", http.StatusOK
}

func GetAllProblemsWithContent(problemList []leetcodeapi.Problem) ([]Problem, string, int) {
	var problems []Problem
	for _, value := range problemList {
		titleSlug := value.TitleSlug
		content, err := leetcodeapi.GetProblemContentByTitleSlug(titleSlug)
		if err != nil {
			return nil, LEETCODE_API_FAILURE_MESSAGE, http.StatusInternalServerError
		}

		topicTags := value.TopicTags
		categories := getCategories(topicTags)

		problem := Problem{
			Title:      value.Title,
			TitleSlug:  value.TitleSlug,
			Difficulty: value.Difficulty,
			QuestionId: value.QuestionId,
			Category:   categories,
			Content:    content.Content,
		}
		problems = append(problems, problem)
	}
	return problems, "", http.StatusOK
}

func getCategories(topicTags []leetcodeapi.TopicTag) []string {
	var categories []string
	for _, value := range topicTags {
		categories = append(categories, value.Name)
	}
	return categories
}
