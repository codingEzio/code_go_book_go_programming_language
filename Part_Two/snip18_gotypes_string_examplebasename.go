package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	/*
		The `basename` func was inspired by the Unix shell utility,
		it removes any prefix/suffixes of the string (=> filename without ext(full))
	*/

	// => archive
	fmt.Println(basenameWithoutLibrary("backup/may/archive.tar.gz"))
	fmt.Println(basenameWithoutLibrary("archive.tar.gz"))
	fmt.Println(basenameWithoutLibrary("archive.zip"))
	fmt.Println(basenameWithoutLibrary("archive") + "\n")

	// => archive.tar AND archive
	fmt.Println(basenameWithStringsLibrary("backup/may/archive.tar.gz"))
	fmt.Println(basenameWithStringsLibrary("archive.tar.gz"))
	fmt.Println(basenameWithStringsLibrary("archive.zip"))
	fmt.Println(basenameWithStringsLibrary("archive") + "\n")

	// => archive.X.Y
	fmt.Println(filepath.Base("backup/may/archive.tar.gz"))
	fmt.Println(filepath.Base("archive.tar.gz"))
	fmt.Println(filepath.Base("archive.zip"))
	fmt.Println(filepath.Base("archive"))
}

func basenameWithoutLibrary(str string) string {
	// Discard last '/' and everything before
	for idx := len(str) - 1; idx >= 0; idx-- {
		if str[idx] == '/' {
			str = str[idx+1:]
			break
		}
	}

	// Preserve everything before last '.'
	for idx := len(str) - 1; idx >= 0; idx-- {
		if str[idx] == '.' {
			str = str[:idx]
		}
	}

	return str
}

func basenameWithStringsLibrary(str string) string {
	slash := strings.LastIndex(str, "/")
	str = str[slash+1:]

	if dot := strings.LastIndex(str, "."); dot >= 0 {
		str = str[:dot]
	}

	return str
}
