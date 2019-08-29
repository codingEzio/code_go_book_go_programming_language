// Expands shell-style (e.g. name=alex age=40) variable references on stdin 😁
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// http://rick.measham.id.au/paste/explain.pl?
var regexPattern = regexp.MustCompile(`\$\w+|\${\w+}`)

func main() {
	/*
		>> go run THIS_FILE name=ALEX greeting="morning\!"
		(INPUT)
			Dear ${name},
				    It's really nice to see you.
				    Good $greeting
		(OUTPUT)
			Dear ALEX,
			    It's really nice to see you.
			    Good morning!
	*/

	log.SetFlags(0) // no date/time info
	log.SetPrefix("[EXPAND] ")

	substi := make(map[string]string, 0)
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			kvpieces := strings.Split(arg, "=")
			if len(kvpieces) != 2 {
				fmt.Fprintln(os.Stderr, "usage: THISCODE KEY=VAL[,KEY=VAL[..]")
				os.Exit(1)
			}

			key, val := kvpieces[0], kvpieces[1]
			substi[key] = val
		}
	}

	missing := make([]string, 0)
	used := make(map[string]bool, 0)

	fn := func(str string) string {
		v, ok := substi[str]
		if !ok {
			missing = append(missing, str)
		}

		used[str] = true
		return v
	}

	byt := &bytes.Buffer{}
	byt.ReadFrom(os.Stdin)
	fmt.Print(expand(byt.String(), fn))

	unused := make([]string, 0)
	for key, _ := range substi {
		if !used[key] {
			unused = append(unused, key)
		}
	}

	if len(unused) > 0 {
		log.Printf("unused  bindings: %s", strings.Join(unused, " "))
	}
	if len(missing) > 0 {
		log.Printf("missing bindings: %s", strings.Join(missing, " "))
	}
}

func expand(str string, fn func(string) string) string {
	wrapper := func(str string) string {
		if strings.HasPrefix(str, "${") {
			str = str[2 : len(str)-1]
		} else {
			str = str[1:]
		}
		return fn(str)
	}
	return regexPattern.ReplaceAllStringFunc(str, wrapper)
}
