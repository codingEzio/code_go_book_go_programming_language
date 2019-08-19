package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Signum(2) =", Signum(2))

	switch CoinFlip() {
	case "heads":
		fmt.Println("Toss =>", "it's head!")
	case "tails":
		fmt.Println("Toss =>", "it's tail!")
	default:
		fmt.Println("Toss =>", "WHAT, IN THE MIDDLE?!")
	}
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func CoinFlip() string {
	rand.Seed(time.Now().UnixNano())
	coin := []string{
		"heads",
		"tails",
		"inthemiddle",
	}

	return coin[rand.Intn(len(coin))]
}
