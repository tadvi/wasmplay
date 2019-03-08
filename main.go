package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)

	// TODO: http.Handle("/wasm", set_your_handler_here_)

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Fatal(http.ListenAndServe(*listen, nil))
}
