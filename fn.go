package leetcode

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"peerprep.assignment6/leetcode"
)

type Response struct {
	Total    int                `json:"total"`
	Problems []leetcode.Problem `json:"problems"`
}

func init() {
	functions.HTTP("GetProblems", getProblems)
}

func getProblems(w http.ResponseWriter, r *http.Request) {
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

	problemList, totalNumberOfQuestion, errorMessage, httpStatusCode := leetcode.GetAllProblems(offset, pageSize)
	if httpStatusCode != http.StatusOK {
		http.Error(w, errorMessage, httpStatusCode)
	}

	problems, errorMessage, httpStatusCode := leetcode.GetAllProblemsWithContent(problemList)

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	response := Response{
		Total:    totalNumberOfQuestion,
		Problems: problems,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
