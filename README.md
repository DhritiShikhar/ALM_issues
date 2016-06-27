alm_jira:

* alm_jira uses Jira library for Go and prints out the result of JQL from Jira Server: https://issues.jboss.org/

* Usage:

./alm_jira -uname=<your username> -pwd=<your password> -jclient=<Jira client> -jql=<Jira Query>  

alm_bugzilla:

* alm_bugzilla uses xmlrpc to fetch issue from bugzilla 

* Usage:

./alm_bugzilla -uname=<your username> -pwd=<your password> -bclient=<Bugzilla client> -search=<Saved search>

* To fetch issues from RedHat Bugzilla, put 
	
	https://bugzilla.redhat.com/xmlrpc.cgi

in bclient.

