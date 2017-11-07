package fs

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"fmt"
)

func propertyValue(fs *JiraFsImpl, matches []string) (string, error) {
	issueKey := matches[1]
	property := matches[2]

	issue, exists := fs.state.IssueKeyToIssueMap[issueKey]

	if !exists {
		return "", nil
	}

	switch property {
	case "summary":
		return issue.Fields.Summary, nil
		break
	case "description":
		return issue.Fields.Description, nil
		break
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
