package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/oleiade/reflections"
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
	fmt.Println(res)
	
	issue, _, err := jiraClient.Issue.Search(jql, nil)
	if err != nil {
		panic(err)
	}
	
	my_fields := []string{"ID", "Self", "Key"}

	for i:=0; i<len(issue); i++ {
		for j:=0; j<len(my_fields); j++ {
			value, _ := reflections.GetField(issue[i], my_fields[j])
			fmt.Printf("%s: %s\n", my_fields[j], value)
		}
		fmt.Printf("\n\n")
	}
}
