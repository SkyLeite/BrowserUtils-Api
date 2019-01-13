package main

import (
	"BrowserUtils/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()

	router.HandleFunc("/about", routes.About).Methods("GET")
	router.HandleFunc("/play/mpv", routes.RunWithMpv).Methods("POST")
	router.HandleFunc("/youtubedl", routes.DownloadWithYoutubeDL).Methods("POST")

	log.Printf("Listening on port 36284")
	log.Fatal(http.ListenAndServe(":36284", router))
}
