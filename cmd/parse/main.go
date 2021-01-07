package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-edtf/parser"
	"log"
)

func main() {

	flag.Parse()

	for _, raw := range flag.Args() {

		d, err := parser.ParseString(raw)

		if err != nil {
			log.Fatal(err)
		}

		enc, err := json.Marshal(d)

		if err != nil {
			log.Fatalf("Failed to marshal date, %v", err)
		}

		fmt.Println(string(enc))
	}
}
