package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/sfomuseum/go-edtf/app/server/www"
	"github.com/sfomuseum/go-edtf/wasm"
)

func main() {

	var host string
	var port int
	
	flag.StringVar(&host, "host", "localhost", "The host to serve requests from")
	flag.IntVar(&port, "port", 8080, "The port to serve requests from")	

	flag.Parse()

	mux := http.NewServeMux()

	wasm_fs := http.FS(wasm.FS)
	wasm_handler := http.FileServer(wasm_fs)

	www_fs := http.FS(www.FS)
	www_handler := http.FileServer(www_fs)

	mux.Handle("/wasm/", http.StripPrefix("/wasm/", wasm_handler))
	mux.Handle("/", www_handler)

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Listening for requests on %s", addr)
	
	err := http.ListenAndServe(addr, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
