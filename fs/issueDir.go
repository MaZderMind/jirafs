package fs

import (
	"github.com/hanwen/go-fuse/fuse"
)

func issueDirList(_ *JiraFsImpl, _ []string) ([]fuse.DirEntry, fuse.Status) {
	return []fuse.DirEntry{
		{Name: "summary", Mode: fuse.S_IFREG},
		{Name: "description", Mode: fuse.S_IFREG},
	}, fuse.OK
}

func issueDirAttr(fs *JiraFsImpl, matches []string) (*fuse.Attr, fuse.Status) {
	issueKey := matches[1]
	_, exists := fs.state.IssueKeyToIssueMap[issueKey]

	if !exists {
		return nil, fuse.ENOENT
	}

	return &fuse.Attr{
		Mode: fuse.S_IFDIR | 0755,
	}, fuse.OK
}
