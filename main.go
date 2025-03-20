package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()                 // Router
	fs := http.FileServer(http.Dir("public")) // file server for serving static files
	mux.Handle("/", fs)                       /// linking route with handler
	http.ListenAndServe(":8080", mux)         // Http server
}
