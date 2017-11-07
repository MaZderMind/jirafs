package jirastate

import "github.com/andygrunwald/go-jira"

type State struct {
	Issues             []jira.Issue
	IssueKeys          []string
	IssueKeyToIssueMap map[string]*jira.Issue
}

func NewState() *State {
	return &State{}
}

func (state *State) SetIssues(issues []jira.Issue) {
	state.Issues = issues
	state.IssueKeys = make([]string, len(issues))
	state.IssueKeyToIssueMap = make(map[string]*jira.Issue, len(issues))

	for i, issue := range issues {
		state.IssueKeys[i] = issue.Key
		state.IssueKeyToIssueMap[issue.Key] = &issues[i]
	}
}
