./jirafs
You are not Connected to any JIRA yet. Run `./jirafs connect <url>` connect via OAuth to the JIRA located at <url>.

./jirafs connect https://jira.apps.seibert-media.net
…
https://jira.apps.seibert-media.net is not connected as the default-Instance. You can run all Commands with a `--instance`-Parameter to work with multiple JIRA instance, ie. `./jirafs --instance another-instance connect <url>`.
Run `./jirafs /mnt &` to mount the Default-View of the connected JIRA to /mnt.

./jirafs -url https://jira.apps.seibert-media.net -username pkoerner -project=JUWEB -passwordFile=/tmp/jirafspw -mountpoint /mnt



----

prefetch issue-keys
prefetch status

prefetch assignees
prefetch labels
prefetch epics
prefetch sprints
prefetch types

update issue-keys every 5 minutes

/projects/JUREC/by-assignee/label/JUREC-1234 -> …
/projects/JUREC/by-assignee/label/JUREC-1234 -> …


/projects/JUREC/by-label/label/JUREC-1234 -> …
/projects/JUREC/by-label/label/JUREC-1234 -> …

/projects/JUREC/by-epic/Epic Title/JUREC-1234 -> …

/projects/JUREC/issues/JUREC-1234




preloading issues?
544 issues, status,labels,assignee, ~1s w/o load
544 issues, summary, 0.76s w/o load

300ms https://jira.apps.seibert-media.net/rest/api/2/issue/createmeta
-> projects

update flag
