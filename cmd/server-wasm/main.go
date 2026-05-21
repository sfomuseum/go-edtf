package main

import (
	_ "context"
	"flag"
	"log"
	"net/http"

	"github.com/sfomuseum/go-edtf/app/server/www"
	"github.com/sfomuseum/go-edtf/wasm"
)

func main() {

	// var server_uri string
	// flag.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	flag.Parse()

	/*
		ctx := context.Background()

		s, err := server.NewServer(ctx, server_uri)

		if err != nil {
			log.Fatalf("Failed to create new server, %v", err)
		}
	*/

	mux := http.NewServeMux()

	wasm_fs := http.FS(wasm.FS)
	wasm_handler := http.FileServer(wasm_fs)

	www_fs := http.FS(www.FS)
	www_handler := http.FileServer(www_fs)

	mux.Handle("/wasm/", http.StripPrefix("/wasm/", wasm_handler))
	mux.Handle("/", www_handler)

	addr := "localhost:8080"

	// log.Printf("Listening on %s", s.Address())
	err := http.ListenAndServe(addr, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
