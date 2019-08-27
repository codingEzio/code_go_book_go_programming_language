package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/code_go_book_go_programming_language/Part_Three/snip20_gojson_examplegetgithubissue/libgetissue"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User  : {{.User.Login}}
Title : {{.Title | printf "%.64s"}}
Age   : {{.CreatedAt | daysAgo }} days
{{end}}`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := libgetissue.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
