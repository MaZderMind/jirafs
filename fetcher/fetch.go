package fetcher

import (
	"github.com/andygrunwald/go-jira"
	"time"
	"fmt"
	"strconv"
	"github.com/mazdermind/jirafs/jirastate"
)

const ISSUE_FETCHER_INTERVAL = 5 * time.Minute

type Fetcher struct {
	client     *jira.Client
	state      *jirastate.State
	projectKey string

	StatusChangeListener []func()
}

func NewFetcher(client *jira.Client, state *jirastate.State, projectKey string) (*Fetcher) {
	f := &Fetcher{
		client:     client,
		state:      state,
		projectKey: projectKey,
	}

	return f
}

func (fetcher *Fetcher) StartFetcher() {
	fetcher.startIssueFetcher()
}

func (fetcher *Fetcher) startIssueFetcher() {
	go fetcher.fetchIssues()

	ticker := time.NewTicker(ISSUE_FETCHER_INTERVAL)
	go func() {
		for range ticker.C {
			fetcher.fetchIssues()
		}
	}()
}

func (fetcher *Fetcher) fetchIssues() {
	fmt.Printf("Fetching issues…\n")
	options := &jira.SearchOptions{
		Fields: []string{
			"resolution",
			"labels",
			"assignee",
			"issuelinks",
			"components",
			"votes",
			"progress",
			"issuetype",
			"watches",
			"description",
			"summary",
			"priority",
			"status",
			"reporter",
			"updated",
			"creator",
			"created",
		},
		MaxResults: 100,
	}
	query := "project=" + strconv.Quote(fetcher.projectKey)

	var issues []jira.Issue
	err := fetcher.client.Issue.SearchPages(query, options, func(issue jira.Issue) error {
		issues = append(issues, issue)
		return nil
	})
	if err != nil {
		fmt.Printf("Error fetching Issues for project %q: %s\n", fetcher.projectKey, err)
		return
	}

	fmt.Printf("Fetched %d issues for query %s\n", len(issues), query)
	fetcher.state.SetIssues(issues)

	fetcher.notifyListener()
}

func (fetcher *Fetcher) notifyListener() {
	for _, listener := range fetcher.StatusChangeListener {
		listener()
	}
}
