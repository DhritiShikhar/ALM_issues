package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	//"reflect"
	"github.com/oleiade/reflections"
)

func main() {
	jiraClient, _ := jira.NewClient(nil, "https://issues.jboss.org/")
	issue, _, err := jiraClient.Issue.Search("project = ARQ AND text ~ 'javadoc'", nil)

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
