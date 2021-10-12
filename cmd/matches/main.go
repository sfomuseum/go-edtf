// matches is a command-line tool to read one or more EDTF strings and emit the EDTF level and feature name they match.
package main

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-edtf/parser"
	"log"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Parse one or more EDTF strings and emit the EDTF level and feature name they match.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s edtf_string(N) edtf_string(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	for _, raw := range flag.Args() {

		level, feature, err := parser.Matches(raw)

		if err != nil {
			log.Fatalf("Failed to parse EDTF string '%s', %v", raw, err)
		}

		fmt.Printf("%s level %d (%s)\n", raw, level, feature)
	}

}
