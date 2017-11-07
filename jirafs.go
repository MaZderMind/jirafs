package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/andygrunwald/go-jira"
	"github.com/mazdermind/jirafs/fetcher"
	"github.com/mazdermind/jirafs/jirastate"
	"github.com/mazdermind/jirafs/fs"
)

const EXIT_CONFIG_ERROR = 1
const EXIT_RUNTIME_ERROR = 2

func main() {
	url := flag.String("url", "", "URL of the Jira-Installation")
	username := flag.String("username", "", "Authentication Username")
	password := flag.String("password", "", "Authentication Password")
	passwordFile := flag.String("passwordFile", "", "File containing Password")
	projectKey := flag.String("project", "", "Project-Key")
	mountpoint := flag.String("mountpoint", "", "Project-Key to mount")
	flag.Parse()

	if *passwordFile != "" {
		content, err := ioutil.ReadFile(*passwordFile)
		if err != nil {
			fmt.Printf("Unable to read password from password-file %s: %s", passwordFile, err)
			os.Exit(EXIT_CONFIG_ERROR)
		}

		passwordFromFile := strings.TrimSpace(string(content))
		password = &passwordFromFile
	}

	if *url == "" || *username == "" || *password == "" || *projectKey == "" || *mountpoint == "" {
		fmt.Printf("Required Fields: url, username, password/passwordFile, project and mountpoint\n")
		flag.Usage()
		os.Exit(EXIT_CONFIG_ERROR)
	}

	jiraClient, err := jira.NewClient(nil, *url)
	if err != nil {
		fmt.Printf("Unable to connect to %q: %s\n", *url, err)
		os.Exit(EXIT_RUNTIME_ERROR)
	}

	jiraClient.Authentication.SetBasicAuth(*username, *password)

	fmt.Printf("Starting Background-Fetcher\n")
	state := jirastate.NewState()
	dataFetcher := fetcher.NewFetcher(jiraClient, state, *projectKey)
	dataFetcher.StartFetcher()

	jirafs, err := fs.NewJiraFs(state, *mountpoint)
	if err != nil {
		fmt.Printf("Unable to mount jirafs to %q: %s\n", *mountpoint, err)
		os.Exit(EXIT_RUNTIME_ERROR)
	}
	jirafs.Serve()
}
