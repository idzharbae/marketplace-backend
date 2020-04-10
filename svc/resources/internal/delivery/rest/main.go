package rest

import (
	"log"
	"net/http"
)

func Start(port string) {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./svc/resources/public/img/"))
	mux.Handle("/img/", http.StripPrefix("/img/", fileServer))
	log.Println("Starting server on " + port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
