package leetcode

import (
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

type Problem struct {
	Complexity  string   `json:"complexity"`
	QuestionId  string   `json:"id"`
	Title       string   `json:"title"`
	TitleSlug   string   `json:"title-slug"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
}

const LEETCODE_API_FAILURE_MESSAGE = "Failed to Connect to Leetcode API"
const MARKDOWN_PARSE_FAIL = "Failed to Parse Into Markdown"

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
		markdownContent, err := convertHTMLtoMarkdown(content)
		if err != nil {
			return nil, MARKDOWN_PARSE_FAIL, http.StatusInternalServerError
		}

		topicTags := value.TopicTags
		categories := getCategories(topicTags)

		problem := Problem{
			Title:      value.Title,
			TitleSlug:  value.TitleSlug,
			Difficulty: value.Difficulty,
			QuestionId: value.QuestionId,
			Category:   categories,
			Content:    markdownContent.Content,
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

func convertHTMLtoMarkdown(content leetcodeapi.ProblemContent) (*leetcodeapi.ProblemContent, error) {
	converter := md.NewConverter("", true, nil)
	htmlContent := content.Content
	markdownContent, err := converter.ConvertString(htmlContent)
	if err != nil {
		return nil, err
	}
	newContent := new(leetcodeapi.ProblemContent)
	newContent.Content = markdownContent

	return newContent, nil
}
