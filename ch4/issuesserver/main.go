// Exercício 4.14
// Run passing github search params, e.g.: repo:golang/go is:open json decoder

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"gopl.io/ch4/github"
)

//!+template

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>Milestone</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td><a href='{{if .Milestone}}{{.Milestone.HTMLURL}}{{end}}'>{{if .Milestone}}{{.Milestone.Title}}{{end}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

//!-template

//!+
func main() {
	http.HandleFunc("/", handler)      // each request calls handler
	http.HandleFunc("/search", search) // each request calls search
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func search(w http.ResponseWriter, r *http.Request) {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}

//!-
