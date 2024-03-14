package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/acorn-io/z"
	"github.com/google/go-github/v60/github"
	"github.com/sirupsen/logrus"
)

type args struct {
	Repo      string `json:"repo"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Assignees string `json:"assignees"`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if len(os.Args) != 2 {
		logrus.Errorf("Usage: %s <JSON parameters>", os.Args[0])
		os.Exit(1)
	}

	var a args
	if err := json.Unmarshal([]byte(os.Args[1]), &a); err != nil {
		logrus.Errorf("Failed to unmarshal JSON parameters: %v", err)
		os.Exit(1)
	}

	gh := github.NewClient(nil)
	if os.Getenv("GPTSCRIPT_GITHUB_TOKEN") != "" {
		gh = gh.WithAuthToken(os.Getenv("GPTSCRIPT_GITHUB_TOKEN"))
	}

	req := &github.IssueRequest{
		Title: &a.Title,
		Body:  &a.Body,
	}

	if a.Assignees != "" {
		assignees := strings.Split(a.Assignees, ",")
		if len(assignees) == 1 {
			req.Assignee = z.Pointer(assignees[0])
		} else {
			req.Assignees = &assignees
		}
	}

	if len(strings.Split(a.Repo, "/")) != 2 {
		logrus.Errorf("invalid repo format (should be 'owner/repo'): %s", a.Repo)
		os.Exit(1)
	}
	owner, repo := strings.Split(a.Repo, "/")[0], strings.Split(a.Repo, "/")[1]

	issue, _, err := gh.Issues.Create(ctx, owner, repo, req)
	if err != nil {
		logrus.Errorf("Failed to create issue: %v", err)
		os.Exit(1)
	}

	fmt.Println("Created issue: ", *issue.HTMLURL)
}
