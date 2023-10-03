package main

import (
	"fmt"

	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

func GetProblems() {
	allProblemList, _ := leetcodeapi.GetAllProblems(0, 50)
	fmt.Println(allProblemList)
}

func main() {
	GetProblems()
}
