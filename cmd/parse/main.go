// parse is a command-line tool to read one or more EDTF strings and return a list of JSON-encode edtf.EDTFDate objects.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-edtf/parser"
	"io"
	"log"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Parse one or more EDTF strings and return a list of JSON-encode edtf.EDTFDate objects.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s edtf_string(N) edtf_string(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	writers := []io.Writer{
		os.Stdout,
	}

	wr := io.MultiWriter(writers...)

	wr.Write([]byte(`[`))

	for i, raw := range flag.Args() {

		d, err := parser.ParseString(raw)

		if err != nil {
			log.Fatalf("Failed to parse EDTF string '%s', %v", raw, err)
		}

		enc, err := json.Marshal(d)

		if err != nil {
			log.Fatalf("Failed to encode EDTFDate, %v", err)
		}

		if i > 0 {
			wr.Write([]byte(`,`))
		}

		wr.Write(enc)
	}

	wr.Write([]byte(`]`))
}
