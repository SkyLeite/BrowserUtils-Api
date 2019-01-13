package main

import (
	"BrowserUtils/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s", c.App.Version)
	}

	app := cli.NewApp()
	app.Name = "Browser Utils"
	app.Version = "v1.0"

	app.Action = func(c *cli.Context) error {
		var router = mux.NewRouter()

		router.HandleFunc("/about", routes.About).Methods("GET")
		router.HandleFunc("/play/mpv", routes.RunWithMpv).Methods("POST")
		router.HandleFunc("/youtubedl", routes.DownloadWithYoutubeDL).Methods("POST")

		log.Printf("Listening on port 36284")
		return http.ListenAndServe(":36284", router)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
