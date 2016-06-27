package main

import (
	"fmt"
	"github.com/kolo/xmlrpc"
	"time"
	"flag"
)

type Bug struct {
	Id             int       `xmlrpc:"id"`
	Summary        string    `xmlrpc:"summary"`
	Status         string    `xmlrpc:"status"`
	CreationTime   time.Time `xmlrpc:"creation_time"`
	AssignedTo     string    `xmlrpc:"assigned_to"`
	Component      []string  `xmlrpc:"component"`
	Product        string    `xmlrpc:"product"`
	Platform       string    `xmlrpc:"platform"`
	Priority       string    `xmlrpc:"priority"`
	Severity       string    `xmlrpc:"severity"`
	Creator        string    `xmlrpc:"creator"`
	IsOpen         bool      `xmlrpc:"is_open"`
	Keywords       []string  `xmlrpc:"keywords"`
	LastChangeTime time.Time `xmlrpc:"last_change_time"`
	Resolution     string    `xmlrpc:"resolution"`
	DependsOn      []int     `xmlrpc:"depends_on"`
	Blocks         []int     `xmlrpc:"blocks"`
	DupeOf         int       `xmlrpc:"dupe_of"`
}

type User struct {
	Id int `xmlrpc:"id"`
	Token string `xmlrpc:"token"`
}

type BugsResult struct {
	Bugs []Bug `xmlrpc:"bugs"`
}

type RpcHash map[string]string

func main(){
	user := User{}
	all_bugs := BugsResult{}

	var uname, pwd, bclient, search string

	flag.StringVar(&bclient, "bclient", "https://bugzilla.redhat.com/xmlrpc.cgi", "Bugzilla Client")
	flag.StringVar(&uname, "uname", "", "Username")
	flag.StringVar(&pwd, "pwd", "", "Password")
	flag.StringVar(&search, "search", "", "Saved search")
	flag.Parse()

	client, _ := xmlrpc.NewClient(bclient, nil)

	client.Call("User.login", RpcHash{"login": uname, "password": pwd}, &user)

	client.Call("Bug.search", RpcHash{"savedsearch": search}, &all_bugs)

	for i:=0; i<len(all_bugs.Bugs); i++ {
		fmt.Println(all_bugs.Bugs[i].Id)
		fmt.Println(all_bugs.Bugs[i].Summary)
		fmt.Println("\n\n")
	}
}
	
