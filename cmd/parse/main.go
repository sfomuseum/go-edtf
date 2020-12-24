package main

import (
	"flag"
	"log"
	"github.com/whosonfirst/go-edtf/parser"
)

func main() {

	flag.Parse()

	for _, raw := range flag.Args(){

		_, err := parser.ParseString(raw)

		if err != nil {
			log.Fatal(err)
		}
		
	}
}
