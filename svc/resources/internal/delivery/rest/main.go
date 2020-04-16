package rest

import (
	"log"
	"net/http"
	"os"
	"path"
)

func Start(port string) {
	e, err := os.Executable()
	if err != nil {
		panic(err)
	}
	imgPath := path.Dir(e) + "/public/img/"
	log.Println(imgPath)

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(imgPath))
	mux.Handle("/img/", http.StripPrefix("/img/", fileServer))
	log.Println("Starting server on " + port)
	err = http.ListenAndServe(port, mux)
	log.Fatal(err)
}
