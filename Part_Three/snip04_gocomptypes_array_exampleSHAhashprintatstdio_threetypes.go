// Prints the SHA256 hash of its standard input (SHA384/SHA512 is also supported)
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var width = flag.Int("w", 256, "hash with (256 or 384 or 512)")

func main() {
	flag.Parse()

	var function func(bt []byte) []byte

	switch *width {
	case 256:
		function = func(bt []byte) []byte {
			h := sha256.Sum256(bt)
			return h[:]
		}
	case 384:
		function = func(bt []byte) []byte {
			h := sha512.Sum384(bt)
			return h[:]
		}
	case 512:
		function = func(bt []byte) []byte {
			h := sha512.Sum512(bt)
			return h[:]
		}
	default:
		log.Fatal("Unexpected width specified")
	}

	bt, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", function(bt))
}
