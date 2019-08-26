// Get info of repo's issues (HOW: go run THIS repo:golang/go is:open json decoder)
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/code_go_book_go_programming_language/Part_Three/snip20_gojson_examplegetgithubissue/libgetissue"
)

func main() {
	result, err := libgetissue.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	format := "$%-7d %10.9s %.55s\n"
	now := time.Now()

	pastDay := make([]*libgetissue.Issue, 0)
	pastMonth := make([]*libgetissue.Issue, 0)
	pastYear := make([]*libgetissue.Issue, 0)

	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(day):
			pastDay = append(pastDay, item)
		case item.CreatedAt.After(month) && item.CreatedAt.Before(day):
			pastMonth = append(pastMonth, item)
		case item.CreatedAt.After(year) && item.CreatedAt.Before(month):
			pastYear = append(pastYear, item)
		}
	}

	if len(pastDay) > 0 {
		fmt.Printf("\nPast day:\n")
		for _, item := range pastDay {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	} else {
		fmt.Printf("\nPast day:\nNONE\n")
	}

	if len(pastMonth) > 0 {
		fmt.Printf("\nPast month:\n")
		for _, item := range pastMonth {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}

	if len(pastYear) > 0 {
		fmt.Printf("\nPast year:\n")
		for _, item := range pastYear {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
}
