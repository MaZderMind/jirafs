mkdir -p $GOPATH/src/github.com/mazdermind
cd $GOPATH/src/github.com/mazdermind/
git clone ssh://git@bitbucket.apps.seibert-media.net:7999/~pkoerner/jirafs.git
cd jirafs
go get
go build

./jirafs -url https://jira.apps.seibert-media.net -username pkoerner -project=JUWEB -passwordFile=/tmp/jirafspw -mountpoint /mnt
