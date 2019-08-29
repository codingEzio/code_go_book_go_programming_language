package main

import (
	"fmt"
)

func main() {
	/*
		The origin of this example code
		* Q	https://stackoverflow.com/q/52980172/6273859
		* A	https://stackoverflow.com/a/52980350/6273859 (here)
	*/

	var rmdirs []func()
	tempDirs := []string{"01", "02", "03", "04"}

	for _, dir := range tempDirs {

		dir := dir

		fmt.Printf("[OUTER] dir=%2v, *dir=%p\n", dir, &dir)

		rmdirs = append(rmdirs, func() {
			fmt.Printf("[INNER] dir=%2v, *dir=%p\n", dir, &dir)
		})
	}

	fmt.Println("---------------------------------------")

	for _, f := range rmdirs {
		f()
	}
}
