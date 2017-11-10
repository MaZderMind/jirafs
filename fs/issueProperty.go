package fs

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"strings"
)

type issueProperty struct {
	name       string
	getContent func(issue *jira.Issue) string
}

var issueProperties = [...]issueProperty{
	{
		name: "summary",
		getContent: func(issue *jira.Issue) string {
			return issue.Fields.Summary
		},
	}, {
		name: "description",
		getContent: func(issue *jira.Issue) string {
			return issue.Fields.Description
		},
	}, {
		name: "assignee",
		getContent: func(issue *jira.Issue) string {
			if issue.Fields.Assignee == nil {
				return ""
			}
			return issue.Fields.Assignee.Name
		},
	}, {
		name: "creator",
		getContent: func(issue *jira.Issue) string {
			if issue.Fields.Creator == nil {
				return ""
			}
			return issue.Fields.Creator.Name
		},
	}, {
		name: "reporter",
		getContent: func(issue *jira.Issue) string {
			if issue.Fields.Reporter == nil {
				return ""
			}
			return issue.Fields.Reporter.Name
		},
	}, {
		name: "type",
		getContent: func(issue *jira.Issue) string {
			return issue.Fields.Type.Name
		},
	},
}

func issuePropertiesRegex() string {
	var issuePropertyNames []string

	for _, issueProperty := range issueProperties {
		issuePropertyNames = append(issuePropertyNames, issueProperty.name)
	}

	return strings.Join(issuePropertyNames, "|")
}

func propertyValue(fs *JiraFsImpl, matches []string) (string, error) {
	issueKey := matches[1]
	property := matches[2]

	issue, exists := fs.state.IssueKeyToIssueMap[issueKey]

	if !exists {
		return "", nil
	}

	for _, issueProperty := range issueProperties {
		if issueProperty.name == property {
			return issueProperty.getContent(issue), nil
		}
	}

	return "", fmt.Errorf("unknown Issue-Property: %s", property)
}

func issuePropertyRead(fs *JiraFsImpl, matches []string) (file nodefs.File, code fuse.Status) {
	value, err := propertyValue(fs, matches)
	if err != nil {
		return nil, fuse.ENOENT
	}

	return nodefs.NewDataFile([]byte(value)), fuse.OK
}

func issuePropertyAttr(fs *JiraFsImpl, matches []string) (*fuse.Attr, fuse.Status) {
	value, err := propertyValue(fs, matches)
	if err != nil {
		return nil, fuse.ENOENT
	}

	return &fuse.Attr{
		Mode: fuse.S_IFREG | 0644,
		Size: uint64(len(value)),
	}, fuse.OK
}
