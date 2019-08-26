// Get info of repo's issues (HOW: go run THIS repo:golang/go is:open json decoder)
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/code_go_book_go_programming_language/Part_Three/snip20_gojson_examplegetgithubissue/libgetissue"
)

func main() {

	result, err := libgetissue.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-7d %10.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
