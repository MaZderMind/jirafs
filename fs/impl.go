package fs

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/mazdermind/jirafs/jirastate"
	"regexp"
)

type path struct {
	regex  *regexp.Regexp
	attrFn func(fs *JiraFsImpl, matches []string) (*fuse.Attr, fuse.Status)
	listFn func(fs *JiraFsImpl, matches []string) ([]fuse.DirEntry, fuse.Status)
	readFn func(fs *JiraFsImpl, matches []string) (file nodefs.File, code fuse.Status)
}

func compileOrNil(expr string) *regexp.Regexp {
	r, err := regexp.Compile(expr)
	if err != nil {
		return nil
	}

	return r
}

var paths = [...]path{
	{
		regex:  compileOrNil("^$"),
		attrFn: rootDirAttr,
		listFn: rootDirList,
		readFn: nil,
	},
	{
		regex:  compileOrNil("^issues$"),
		attrFn: issuesDirAttr,
		listFn: issuesDirList,
		readFn: nil,
	},
	{
		regex:  compileOrNil("^issues/([^/]+)$"),
		attrFn: issueDirAttr,
		listFn: issueDirList,
		readFn: nil,
	},
	{
		regex:  compileOrNil("^issues/([^/]+)/(summary|description)$"),
		attrFn: issuePropertyAttr,
		listFn: nil,
		readFn: issuePropertyRead,
	},
}

type JiraFsImpl struct {
	pathfs.FileSystem

	state *jirastate.State
}

func (fs *JiraFsImpl) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	for _, path := range paths {
		matches := path.regex.FindStringSubmatch(name)
		if matches == nil {
			continue
		}

		if path.attrFn == nil {
			return nil, fuse.ENOENT
		}

		return path.attrFn(fs, matches)
	}

	return nil, fuse.ENOENT
}

func (fs *JiraFsImpl) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	for _, path := range paths {
		matches := path.regex.FindStringSubmatch(name)
		if matches == nil {
			continue
		}

		if path.listFn == nil {
			return nil, fuse.ENOENT
		}

		return path.listFn(fs, matches)
	}

	return nil, fuse.ENOENT
}

func (fs *JiraFsImpl) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	for _, path := range paths {
		matches := path.regex.FindStringSubmatch(name)
		if matches == nil {
			continue
		}

		if flags&fuse.O_ANYWRITE != 0 {
			return nil, fuse.EPERM
		}

		if path.readFn == nil {
			return nil, fuse.ENOENT
		}

		return path.readFn(fs, matches)
	}

	return nil, fuse.ENOENT
}
