package fetcher

import "github.com/andygrunwald/go-jira"

type StateType struct {
	Issues []jira.Issue
}

var State StateType
