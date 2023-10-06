package hello

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func GetProblems(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	offsetString := queryParams.Get("offset")
	pageSizeString := queryParams.Get("page-size")
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		http.Error(w, "Invalid Query Parameters: offset", http.StatusBadRequest)
	}
	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil {
		http.Error(w, "Invalid Query Parameters: page-size", http.StatusBadRequest)
	}
	var problems []Problem
	allProblemList, err := leetcodeapi.GetAllProblems(offset, pageSize)
	if err != nil {
		http.Error(w, "Failed to Connect to Leetcode API", http.StatusInternalServerError)
		return
	}
	for _, value := range allProblemList.Problems {
		var categories []string
		titleSlug := value.TitleSlug
		content, err := leetcodeapi.GetProblemContentByTitleSlug(titleSlug)
		if err != nil {
			http.Error(w, "Failed to Connect to Leetcode API", http.StatusInternalServerError)
			return
		}
		for _, value := range value.TopicTags {
			categories = append(categories, value.Name)
		}
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

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	w.Header().Set("Content-Type", "application/json")
	if err := encoder.Encode(problems); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
