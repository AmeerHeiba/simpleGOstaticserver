package main

import (
	"log"
	"net/http"

	"simpleserver.com/m/routes"
)

func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request Recived %s :: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := routes.RegisterRoutes()
	loggedRouter := LoggingMiddleWare(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedRouter,
	}

	log.Println("Server started on port 8080...")
	log.Fatal(server.ListenAndServe())
}
