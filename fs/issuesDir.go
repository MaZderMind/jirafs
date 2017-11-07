package fs

import "github.com/hanwen/go-fuse/fuse"

func issuesDirList(fs *JiraFsImpl, _ []string) ([]fuse.DirEntry, fuse.Status) {
	entries := make([]fuse.DirEntry, len(fs.state.Issues))
	for i, issue := range fs.state.Issues {
		entries[i] = fuse.DirEntry{
			Name: issue.Key,
			Mode: fuse.S_IFDIR,
		}
	}
	return entries, fuse.OK
}

func issuesDirAttr(_ *JiraFsImpl, _ []string) (*fuse.Attr, fuse.Status) {
	return &fuse.Attr{
		Mode: fuse.S_IFDIR | 0755,
	}, fuse.OK
}
