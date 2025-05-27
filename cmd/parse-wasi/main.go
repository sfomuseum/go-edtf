//go:build wasip1
package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/sfomuseum/go-edtf/parser"
)

func main() {

	flag.Parse()

	for _, raw := range flag.Args() {
		fmt.Println(parse(raw))
	}
}

//export parse
func parse(raw string) string {

	d, err := parser.ParseString(raw)

	if err != nil {
		return err.Error()
	} else {

		v, err := json.Marshal(d)

		if err != nil {
			return err.Error()
		}

		return string(v)
	}

}
