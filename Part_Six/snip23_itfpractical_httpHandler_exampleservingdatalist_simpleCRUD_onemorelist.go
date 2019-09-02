package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type MyDB struct {
	sync.Mutex
	db map[string]int
}

var listHTML = template.Must(template.New("list").Parse(`
<html>
<body>
<table style="border: 1px solid black;">
	<tr>
		<th>item</th>
		<th>price</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<th>{{$k}}</th>
		<th>{{$v}}</th>
	</tr>
{{end}}
</table>
</html>
</body>
`))

// http://localhost:8080/list
// http://localhost:8080/create?item=ITEM&price=PRICE
// http://localhost:8080/read?item=ITEM
// http://localhost:8080/update?item=ITEM&price=PRICE
// http://localhost:8080/delete?item=ITEM
func main() {
	db := &MyDB{}
	db.db = make(map[string]int, 0)

	db.db["shoe"] = 100 // ?item=name&price=100

	http.HandleFunc("/list", db.List)
	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/read", db.Read)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)

	fmt.Println("Now serving at http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (m *MyDB) List(wr http.ResponseWriter, req *http.Request) {
	m.Lock()
	listHTML.Execute(wr, m.db)
	m.Unlock()
}

func (m *MyDB) Create(wr http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(wr, "No item given", http.StatusBadRequest) // 400
		return
	}

	priceStr := req.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(wr, "No integer price given", http.StatusBadRequest) // 400
		return
	}

	if _, ok := m.db[item]; ok {
		http.Error(wr, fmt.Sprintf("%s already exists", item), http.StatusBadRequest)
		return
	}

	m.Lock()
	if m.db == nil {
		m.db = make(map[string]int, 0)
	}
	m.db[item] = price
	m.Unlock()
}

func (m *MyDB) Read(wr http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(wr, "No item given", http.StatusBadRequest) // 400
		return
	}

	if _, ok := m.db[item]; !ok {
		http.Error(wr, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
		return
	}

	m.Lock()
	fmt.Fprintf(wr, "%s: %d\n", item, m.db[item])
	m.Unlock()
}

func (m *MyDB) Update(wr http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(wr, "No item given", http.StatusBadRequest) // 400
		return
	}

	priceStr := req.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(wr, "No integer price given", http.StatusBadRequest) // 400
		return
	}

	if _, ok := m.db[item]; !ok {
		http.Error(wr, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
		return
	}

	m.Lock()
	m.db[item] = price
	m.Unlock()
}

func (m *MyDB) Delete(wr http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(wr, "No item given", http.StatusBadRequest) // 400
		return
	}

	if _, ok := m.db[item]; !ok {
		http.Error(wr, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
		return
	}

	m.Lock()
	delete(m.db, item)
	m.Unlock()
}
