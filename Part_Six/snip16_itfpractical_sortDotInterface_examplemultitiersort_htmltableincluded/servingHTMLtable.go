// Serves an HTML table with a stable column sort
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/code_go_book_go_programming_language/Part_Six/snip16_itfpractical_sortDotInterface_examplemultitiersort_htmltableincluded/libmultitiersort"
)

var people = []multitiersort.Person{
	{"Alice", 20},
	{"Bob", 12},
	{"Bob", 20},
	{"Alice", 12},
}

var htmlTmpl = template.Must(template.New("people").Parse(`
<html>
<body>

<table>
	<tr>
		<th><a href="?sort=name">name</a></th>
		<th><a href="?sort=age">age</a></th>
	</tr>
{{range .}}
	<tr>
		<td>{{.Name}}</td>
		<td>{{.Age}}</td>
	<tr>
{{end}}
</table>

</body>
</html>
`))

func main() {
	col := multitiersort.NewByColumns(people, 2)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.FormValue("sort") {
		case "age":
			col.Select(col.LessAge)
		case "name":
			col.Select(col.LessName)
		}

		sort.Sort(col)
		err := htmlTmpl.Execute(writer, people)
		if err != nil {
			log.Printf("template error: %s", err)
		}
	})
	fmt.Println("Now serving at http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
