package fs

import (
	"github.com/hanwen/go-fuse/fuse"
)

func issueDirList(_ *JiraFsImpl, _ []string) ([]fuse.DirEntry, fuse.Status) {
	var entries []fuse.DirEntry;
	for _, issueProperty := range issueProperties {
		entry := fuse.DirEntry{Name: issueProperty.name, Mode: fuse.S_IFREG}
		entries = append(entries, entry)
	}

	return entries, fuse.OK
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
