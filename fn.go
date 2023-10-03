package hello

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "someone"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func GetProblems(w http.ResponseWriter, r *http.Request) {
	allProblemList, _ := leetcodeapi.GetAllProblems(0, 50)
	jsonData, err := json.Marshal(allProblemList.Problems)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetProblemContent(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	titleSlug := queryParam.Get("title-slug")
	problemContent, err := leetcodeapi.GetProblemContentByTitleSlug(titleSlug)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Fprintln(w, problemContent)
}
