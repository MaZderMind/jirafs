package fs

import "github.com/hanwen/go-fuse/fuse"

func rootDirList(_ *JiraFsImpl, _ []string) ([]fuse.DirEntry, fuse.Status) {
	return []fuse.DirEntry{
		{Name: "issues", Mode: fuse.S_IFDIR},
	}, fuse.OK
}

func rootDirAttr(_ *JiraFsImpl, _ []string) (*fuse.Attr, fuse.Status) {
	return &fuse.Attr{
		Mode: fuse.S_IFDIR | 0755,
	}, fuse.OK
}
