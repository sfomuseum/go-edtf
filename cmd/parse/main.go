package main

import (
	"flag"
	"github.com/whosonfirst/go-edtf/parser"
	"log"
)

func main() {

	flag.Parse()

	for _, raw := range flag.Args() {

		_, err := parser.ParseString(raw)

		if err != nil {
			log.Fatal(err)
		}

	}
}
