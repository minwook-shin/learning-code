package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/andygrunwald/go-jira"
)

func main() {
	jiraClient, _ := jira.NewClient(nil, "https://issues.apache.org/jira/")
	issue, _, _ := jiraClient.Issue.Get("MESOS-1", nil)

	fmt.Println(issue.Key)
	fmt.Println(issue.Fields.Summary)
	fmt.Println(issue.Fields.Type.Name)
	fmt.Println(issue.Fields.Priority.Name)

	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: transport}

	jiraClient, _ = jira.NewClient(client, "https://issues.apache.org/jira/")
	issue, _, _ = jiraClient.Issue.Get("MESOS-1", nil)

	fmt.Println(issue.Key)
	fmt.Println(issue.Fields.Summary)
	fmt.Println(issue.Fields.Type.Name)
	fmt.Println(issue.Fields.Priority.Name)

	jiraClient, _ = jira.NewClient(nil, "https://jira.atlassian.com/")
	request, _ := jiraClient.NewRequest("GET", "/rest/api/2/project", nil)

	projects := new([]jira.Project)
	_, err := jiraClient.Do(request, projects)
	if err != nil {
		panic(err)
	}

	for _, project := range *projects {
		fmt.Println(project.Key, project.Name)
	}
}
