package main

import (
	"log"
	"time"
)

func main() {
	/*
		From what I see, the process behaves like this
		[1] Passing `msg` into `trace` (proceed but ends before `return`
		[2] The `bigSlow..` regains control again, sleep for two more seconds
		[3] The `trace`     regains control (or the other way), execute the `return` stmt
	*/

	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("Hello... Neeeew York!")()
	time.Sleep(2 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("Enter %s", msg)

	return func() {
		log.Printf("Exit  %s (%s)", msg, time.Since(start))
	}
}
