package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var otherIssues []*github.Issue
	var monthIssues []*github.Issue
	var yearIssues []*github.Issue

	now := time.Now()
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreatedAt.After(month) {
			monthIssues = append(monthIssues, item)
			continue
		}
		if item.CreatedAt.After(year) {
			yearIssues = append(yearIssues, item)
			continue
		}

		otherIssues = append(otherIssues, item)
		//fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	printIssues(monthIssues)
	printIssues(yearIssues)
	printIssues(otherIssues)
}
func printIssues(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
