package main

import (
	"flag"
	"github.com/whosonfirst/go-edtf"
	"log"
)

func main() {

	flag.Parse()

	for _, str := range flag.Args() {

		d, err := edtf.Parse(str)
		log.Println(str, d, err)
	}
}
