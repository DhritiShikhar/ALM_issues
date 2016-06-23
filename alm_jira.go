package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"flag"
)

func main() {
	var uname, pwd, jclient, jql string
	
	flag.StringVar(&uname, "uname", "", "Username")
	flag.StringVar(&pwd, "pwd", "", "Password")
	flag.StringVar(&jclient, "jclient", "https://issues.jboss.org/", "Jira Client")
	flag.StringVar(&jql, "jql", "project = ARQ AND text ~ 'javadoc'", "Jira Query Language") 
	flag.Parse()

	jiraClient, err := jira.NewClient(nil, jclient)
	if err != nil {
		panic(err)
	}
	
	res, err := jiraClient.Authentication.AcquireSessionCookie(uname, pwd)
	if err != nil || res == false {
		panic(err)
	}
	
	issue, _, err := jiraClient.Issue.Search(jql, nil)
	if err != nil {
		panic(err)
	}
	
	for i:=0; i<len(issue); i++ {
		fmt.Println("Issue ID:\t", issue[i].ID)
		fmt.Println("Issue Key:\t", issue[i].Key)
	
		issue, _, _ := jiraClient.Issue.Get(issue[i].Key)
		fmt.Println("Issue Summary:\t", issue.Fields.Summary)
		fmt.Println("Issue Description: \n", issue.Fields.Description)

		fmt.Println("\n\n")
	}
}
