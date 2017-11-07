package fs

import (
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/mazdermind/jirafs/jirastate"
)

type JiraFs struct {
	server *fuse.Server
}

func NewJiraFs(state *jirastate.State, mountpoint string) (*JiraFs, error) {
	impl := &JiraFsImpl{FileSystem: pathfs.NewDefaultFileSystem()}
	impl.state = state

	pathNodeFs := pathfs.NewPathNodeFs(impl, nil)
	server, _, err := nodefs.MountRoot(mountpoint, pathNodeFs.Root(), nil)

	if err != nil {
		return nil, err
	}

	jfs := &JiraFs{}
	jfs.server = server
	return jfs, nil
}

func (jfs *JiraFs) Serve() {
	jfs.server.Serve()
}
